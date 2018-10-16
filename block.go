package main

import (
	"time"
	"encoding/binary"
	"bytes"
	"log"
)
//2. 升级版（区块字段完整）
//补充区块字段
//更新计算哈希函数
//优化代码
//定义区块结构
type Block struct {
	//版本号
	Version uint64
	//前区块哈希值
	prevHash []byte
	//梅克尔根Merkle,是一个哈希值
	MerkleRoot []byte
	//时间戳
	TimeStamp uint64
	//难度值
	Difficulty uint64
	//随机数
	Nonce uint64
	//当前区块哈希值
	Hash []byte
	//区块数据
	Data []byte
}
//创建区块
func NewBlock(data string,prevHash []byte)*Block  {
	block:= Block{
		Version:00,
		prevHash:prevHash,
		MerkleRoot:[]byte{},
		TimeStamp:uint64(time.Now().Unix()),
		Difficulty:100,//先随便写
		Nonce:100,
		Hash:[]byte{},//TODO
		Data:[]byte(data),
	}
	//block.Hash=block.SetHash()
	//创建一个pow对象
	pow:=NewProofOfWork(&block)
	//查找随机数，不停进行哈希运算
	hash,nonce:=pow.Run()
	block.Hash=hash
	block.Nonce=nonce
	return &block

}

/*
//生成哈希
func (block *Block)SetHash()[]byte  {
	//拼接当前区块的数据
	var blockByteInfo []byte
	/*
	blockByteInfo=append(blockByteInfo,uint64ToByte(block.Version)...)
	blockByteInfo=append(blockByteInfo,block.prevHash...)
	blockByteInfo=append(blockByteInfo,block.Merkle...)
	blockByteInfo=append(blockByteInfo,uint64ToByte(block.TimeStamp)...)
	blockByteInfo=append(blockByteInfo,uint64ToByte(block.Difficulty)...)
	blockByteInfo=append(blockByteInfo,uint64ToByte(block.Nonce)...)
	blockByteInfo=append(blockByteInfo,block.Data...)
	*/
	/*
	tmp:=[][]byte{
		uint64ToByte(block.Version),
		block.prevHash,
		block.MerkleRoot,
		uint64ToByte(block.TimeStamp),
		uint64ToByte(block.Difficulty),
		uint64ToByte(block.Nonce),
		block.Data,
	}
	blockByteInfo=bytes.Join(tmp,[]byte{})
	hash:=sha256.Sum256(blockByteInfo)
	return hash[:]
}
*/

//辅助函数
func uint64ToByte(data uint64) []byte {
	var buffer bytes.Buffer
	err:=binary.Write(&buffer,binary.BigEndian,data)
	if err!=nil{
		log.Panic(err)
	}
	return buffer.Bytes()
}