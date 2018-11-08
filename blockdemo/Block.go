package day01

import (
	"crypto/sha256"
	"time"
	"encoding/binary"
	"bytes"
)

//定义区块
type Block struct {
	//版本号
	Version uint64

	//前hash
	PrevHash []byte

	//Merkle树
	MerkleRoot []byte

	//时间戳
	TimeStamp uint64

	//随机数
	Nonce uint64

	//当前hash
	Hash []byte

	Data []byte
}

//创建区块
func NewBlock(prevHash []byte,data string) *Block {

	block := Block{
		Version:0,
		PrevHash: prevHash,
		MerkleRoot:[]byte{},
		TimeStamp:uint64(time.Now().Unix()),
		Nonce:0,
		Hash:     []byte{},
		Data:     []byte(data),
	}
	block.SetHash()
	return &block
}

func (block*Block)SetHash()  {

		//Version:0,
		//PrevHash: prevHash,
		//MerkleRoot:[]byte{},
		//TimeStamp:uint64(time.Now().Unix()),
		//Nonce:0,
		//Hash:     []byte{},
		//Data:     []byte(data),



	//var info []byte
	//info= append(info, Uint64ToByte(block.Version)...)
	//info= append(info, block.PrevHash...)
	//info= append(info, block.MerkleRoot...)
	//info= append(info, Uint64ToByte(block.TimeStamp)...)
	//info= append(info, Uint64ToByte(block.Nonce)...)
	//info = append(info, block.Data...)

	infos:=[][]byte{Uint64ToByte(block.Version),block.PrevHash, block.MerkleRoot, Uint64ToByte(block.TimeStamp), Uint64ToByte(block.Nonce), block.Data}
	info := bytes.Join(infos, []byte{})

	sum256 := sha256.Sum256(info)
	block.Hash=sum256[:]
}

func Uint64ToByte(data uint64)  []byte{
	var buf bytes.Buffer

	binary.Write(&buf,binary.BigEndian,&data)
	return buf.Bytes()
}
