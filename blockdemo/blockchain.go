package day01

import "crypto/sha256"

type BlockChain struct {

	Blocks []*Block

}

func CreateGensisBlock() *BlockChain {
	gensisInfo:="this is the gensis block"

	sum256 := sha256.Sum256([]byte("00000000000000000000000000000"))

	block := NewBlock(sum256[:], gensisInfo)
	chain := BlockChain{Blocks: []*Block{block}}
	return &chain
}

func (blockChain* BlockChain) AddBlock(data string)  {

	length := len(blockChain.Blocks)
	lastBlock := blockChain.Blocks[length-1]
	prevHash:=lastBlock.Hash
	block := NewBlock(prevHash, data)
	blockChain.Blocks= append(blockChain.Blocks, block)

}
