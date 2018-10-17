package main

import (
	"github.com/bolt"
	"log"
	"os"
)

//B. v3版本思路
//1. bolt数据库的介绍
//轻量级的、开源的
//
//go语言实现的
//
//key->value进行读取（map）
//2. BlockChain结构重写
//使用数据库代替数组
//
//3. NewBlockChain函数重写
//由对数组操作改写成==对数据库操作==，创建数据库
//4. AddBlock函数重写
//对数据库的读取和写入
//5.打印数据
//对数据库的遍历（迭代器Iterator）
//6.命令行
//命令行介绍及编写
//添加区块命令
//打印区块链命令

//引入区块链(区块的数组)

//2. BlockChain结构重写
//使用数据库代替数组
//
type BlockChain struct {
	//Blocks []*Block
	db *bolt.DB
	//存储最后一个区块的哈希
	tail []byte
}

const blockChainDb = "blockChain.Db"
const blockBucket  = "blockBucket"
//创建创世块
func GenesisBlock() *Block {
	return NewBlock("222222222",[]byte{})
}
//创建区块链
//3. NewBlockChain函数重写
//由对数组操作改写成==对数据库操作==，创建数据库
func NewBlockChain()*BlockChain  {
	/*
	把创世块，作为区块链的第一个区块
	genesisBlock:=GenesisBlock()
	return &BlockChain{[]*Block{genesisBlock}}
	*/

	var lastHash []byte	//最后一个区块的哈希，从数据库中读出来
	db,err:=bolt.Open(blockChainDb,0600,nil)
	if err!=nil{
		log.Panic("bolt.Open failed",err)
		os.Exit(1)
	}
	//defer db.Close()
	//写数据
	db.Update(func(tx *bolt.Tx) error {
		//找到抽屉bucket,如果没有创建
		bucket:=tx.Bucket([]byte(blockBucket))
		if bucket==nil{
			//没有bucket，需要创建
			bucket,err=tx.CreateBucket([]byte(blockBucket))
			if err!=nil{
				log.Panic("创建bucket失败",err)
				os.Exit(1)
			}
			//创建一个创世块，作为第一个区块添加到区块链中
			genesisBlock:=GenesisBlock()
			//hash作为key,block的字节流作为value
			bucket.Put(genesisBlock.Hash,genesisBlock.Serialize())
			bucket.Put([]byte("lastHashKey"),genesisBlock.Hash)
			lastHash=genesisBlock.Hash
		}else {
			lastHash=bucket.Get([]byte("lastHashKey"))
		}

		return nil
	})

	return &BlockChain{db,lastHash}
}
//添加区块
func (bc *BlockChain)AddBlock(data string)  {
	/*
	length:=len(bc.Blocks)
	//最后一个区块
	lastBlock:=bc.Blocks[length-1]
	//最后一个区块的哈希值是新区块的前哈希
	prevHash:=lastBlock.Hash
	block:=NewBlock(data,prevHash)
	bc.Blocks=append(bc.Blocks,block)
	*/
	//如何获取前区块的哈希呢？？
	db:=bc.db			//区块链数据库
	lastHash:=bc.tail	//最后一个区块哈希
	//完成数据的添加
	db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(blockBucket))
		if bucket==nil{
			log.Panic("bucket不应该为空，请检查")
			os.Exit(1)
		}
		//创建新的区块
		block:=NewBlock(data,lastHash)

		//添加到区块链db中
		//hash作为key,block字节流作为value
		bucket.Put(block.Hash,block.Serialize())
		bucket.Put([]byte("lastHashKey"),block.Hash)
		bc.tail=block.Hash
		return nil
	})

}
