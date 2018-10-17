package main

import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
)

//1. POW介绍
//定义一个工作量证明的结构ProofOfWork
type ProofOfWork struct {
	//a. block
	block *Block
	//b. 目标值（一个很大的数）
	target *big.Int
}

//2. 提供创建POW的函数
func NewProofOfWork(block *Block)*ProofOfWork  {
	pow:=ProofOfWork{
		block:block,
	}
	//自定义难度值，先写成固定值，是string类型，需要转换
	targetString:="0000100000000000000000000000000000000000000000000000000000000000"
	//将string类型转换成big.Int
	bigIntTmp:=big.Int{}
	//func (z *Int) SetString(s string, base int) (*Int, bool) {
	//指定16进制格式
	bigIntTmp.SetString(targetString,16)
	pow.target=&bigIntTmp

	return &pow
}
//3. 提供计算,不断计算hash的函数
func (pow *ProofOfWork) Run() ([]byte,uint64) {
	//1、拼接数据（区块的数据，还有不断变化的随机数）
	//2、哈希运算
	//3、与pow中的target作比较
		//1、找到了返回
		//2、没有找到，继续找,随机数+1
	var hash [32]byte
	var Nonce uint64
	fmt.Printf("target:%x\n",pow.target)
	for{
		//1、拼接数据（区块的数据，还有不断变化的随机数）
		tmp:=[][]byte{
			uint64ToByte(pow.block.Version),
			pow.block.PrevHash,
			pow.block.MerkleRoot,
			uint64ToByte(pow.block.TimeStamp),
			uint64ToByte(pow.block.Difficulty),
			uint64ToByte(Nonce),
			pow.block.Data,
		}
		//将二维切片链接起来，返回一维切片
		blockInfo:=bytes.Join(tmp,[]byte{})
		//2、哈希运算
		hash=sha256.Sum256(blockInfo)

		bigIntTemp:=big.Int{}
		//将得到的定长的哈希值转换成big.Int
		bigIntTemp.SetBytes(hash[:])
		//3、与pow中的target作比较
			//比较当前的哈希值与目标哈希值，如果当前哈希值小于目标哈希值，找到了，否则继续找
			//   -1 if x <  y
			//    0 if x == y
			//   +1 if x >  y
			//
			//func (x *Int) Cmp(y *Int) (r int) {
		if bigIntTemp.Cmp(pow.target) == -1{
			//1、找到了返回
			fmt.Printf("挖矿成功！哈希值：%x,随机数:%d\n",hash,Nonce)
			break
		}else {
			//2、没有找到，继续找,随机数+1
			Nonce++
		}
	}

	return hash[:],Nonce
}

//4. 提供一个校验函数
//IsValid()
