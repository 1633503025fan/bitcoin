package main

import (
	"fmt"
)

func main()  {
	data:="11111111"
	//block:=NewBlock(data,[]byte{})
	bc:=NewBlockChain(data)
	bc.AddBlock("老师向我转了1比特币")
	bc.AddBlock("老师向我转了1比特币")
	for i,block:=range bc.Blocks{
		fmt.Printf("===================\n")
		fmt.Printf("当前区块高度:%d\n",i)
		fmt.Printf("版本号:%d\n",block.Version)
		fmt.Printf("前区块哈希值:%x\n",block.prevHash)
		fmt.Printf("梅克尔根:%x\n",block.MerkleRoot)
		fmt.Printf("时间戳:%d\n",block.TimeStamp)
		fmt.Printf("难度值:%d\n",block.Difficulty)
		fmt.Printf("随机数%d\n",block.Nonce)
		fmt.Printf("当前区块哈希值:%x\n",block.Hash)
		fmt.Printf("区块数据:%s\n",block.Data)
	}

}