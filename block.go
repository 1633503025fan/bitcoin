package main

import (
	"time"
	"encoding/binary"
	"bytes"
	"log"
	"encoding/gob"
	"os"

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
	PrevHash []byte
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
		PrevHash:prevHash,
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

//序列化
func (block *Block)Serialize()[]byte  {
	//编码的数据放到buffer
	var buffer bytes.Buffer
	//使用gob进行序列化（编码）得到字节流
		//1、定义一个编码器
		//2、使用编码器编码
	//func NewEncoder(w io.Writer) *Encoder {
	encode:=gob.NewEncoder(&buffer)
	//func (enc *Encoder) Encode(e interface{}) error {
	err:=encode.Encode(&block)
	if err!=nil{
		log.Panic("编码出错")
		os.Exit(1)
	}
	return buffer.Bytes()
}
//反序列化
func Deserialize(data []byte) Block {
	//使用gob进行反序列化（解码）
	//1、定义一个解码器
	//2、使用解码器解码
	//func NewDecoder(r io.Reader) *Decoder {
	decoder:=gob.NewDecoder(bytes.NewReader(data))
	var block Block
	err:=decoder.Decode(&block)
	if err!=nil{
		log.Panic("解码出错")
		os.Exit(1)
	}
	return block
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