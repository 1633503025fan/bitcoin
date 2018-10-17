package main

import (
	"github.com/bolt"
	"log"
	"os"
)

type BlockChainIterator struct {
	db *bolt.DB
	//游标，用于不断索引
	currentHashPoint []byte
}

func (bc *BlockChain) NewBlockChainIterator()*BlockChainIterator  {
	return &BlockChainIterator{
		bc.db,
		bc.tail,//最初指向区块链的最后一个区块，最着Next()调用，不断前移
	}
}

//迭代器是属于区块链的
//Next方法是属于迭代器的
//1、返回当前的区块
//2、指针前移
func (it *BlockChainIterator)Next()*Block  {
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(blockBucket))
		if bucket==nil{
			log.Panic("迭代器遍历时bucket不应该为空，请检查")
			os.Exit(1)
		}
		//func (b *Bucket) Get(key []byte) []byte {
		blockTmp:=bucket.Get(it.currentHashPoint)
		//解码
		block=Deserialize(blockTmp)
		//游标哈希左移
		it.currentHashPoint=block.PrevHash
		return nil
	})
	return &block
}