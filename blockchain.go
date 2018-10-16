package main

//引入区块链(区块的数组)
type BlockChain struct {
	Blocks []*Block
}

//创建区块链
func NewBlockChain(data string)*BlockChain  {
	// 把创世块，作为区块链的第一个区块
	genesisBlock:=GenesisBlock(data,[]byte{})
	return &BlockChain{[]*Block{genesisBlock}}

}
//添加区块
func (bc *BlockChain)AddBlock(data string)  {
	length:=len(bc.Blocks)
	//最后一个区块
	lastBlock:=bc.Blocks[length-1]
	//最后一个区块的哈希值是新区块的前哈希
	prevHash:=lastBlock.Hash
	block:=NewBlock(data,prevHash)
	bc.Blocks=append(bc.Blocks,block)
}
