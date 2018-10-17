package main

import (
	//"fmt"
)

func main()  {
	//block:=NewBlock(data,[]byte{})
	bc:=NewBlockChain()
	//bc.AddBlock("helloworld")
	//bc.AddBlock("helloitcast")
	/*
	for i,block:=range bc.Blocks {
		fmt.Printf("===================\n")
		fmt.Printf("当前区块高度:%d\n", i)
		fmt.Printf("版本号:%d\n", block.Version)
		fmt.Printf("前区块哈希值:%x\n", block.prevHash)
		fmt.Printf("梅克尔根:%x\n", block.MerkleRoot)
		fmt.Printf("时间戳:%d\n", block.TimeStamp)
		fmt.Printf("难度值:%d\n", block.Difficulty)
		fmt.Printf("随机数%d\n", block.Nonce)
		fmt.Printf("当前区块哈希值:%x\n", block.Hash)
		fmt.Printf("区块数据:%s\n", block.Data)

	}
	*/
	/*
	//创建迭代器
	it:=bc.NewBlockChainIterator()
	//不断调用迭代器Next方法，返回每一个区块
	for{
		block:=it.Next()
		fmt.Printf("============================\n")
		fmt.Printf("版本号:%d\n", block.Version)
		fmt.Printf("前区块哈希值:%x\n", block.prevHash)
		fmt.Printf("梅克尔根:%x\n", block.MerkleRoot)
		fmt.Printf("时间戳:%d\n", block.TimeStamp)
		fmt.Printf("难度值:%d\n", block.Difficulty)
		fmt.Printf("随机数%d\n", block.Nonce)
		fmt.Printf("当前区块哈希值:%x\n", block.Hash)
		fmt.Printf("区块数据:%s\n", block.Data)
		//终止条件
		if len(block.prevHash)==0{
			fmt.Printf("区块链遍历结束")
			break
		}
	}
	*/
	cli:=CLI{bc}
	cli.Run()
}