package main

import (
	"day01"
	"fmt"
	"encoding/hex"
)

//import (
//	"BLC"
//	"fmt"
//	"encoding/hex"
//)
//
////func (b BLC.Block) SetHash(PrevHash []byte,Data []byte)   {
////	var blockByteInfo []byte
////	blockByteInfo = append(blockByteInfo, PrevHash...)
////	blockByteInfo = append(blockByteInfo, Data...)
////	b.Hash= sha256.Sum256(blockByteInfo)
////
////}
//
//
//
//
//
//func main() {
//	chain := BLC.GensisBlock()
//	chain.AddBlock("the second")
//	chain.AddBlock("the third")
//	chain.AddBlock("the fouth")
//	for i,v:=range chain.Blocks{
//		fmt.Println(i,string(v.Data),hex.EncodeToString(v.Hash))
//	}
//
//
//
//
//}

func main() {
	//block := day01.NewBlock([]byte("00000000000011211"),[]byte("3454"))
	//block.SetHash()
	//fmt.Println(hex.EncodeToString(block.PrevHash),hex.EncodeToString(block.Hash),string(block.Data))
	blocks := day01.CreateGensisBlock()

	blocks.AddBlock("001")
	blocks.AddBlock("002")
	blocks.AddBlock("003")
	blocks.AddBlock("004")
	blocks.AddBlock("005")
	for i,block:=range blocks.Blocks {
		fmt.Println(i,hex.EncodeToString(block.PrevHash),hex.EncodeToString(block.Hash),string(block.Data))
	}


}

