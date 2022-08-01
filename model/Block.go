package model

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

type Block struct {
	//时间戳
	Timestamp int64
	//哈希
	Hash     []byte
	PrevHash []byte
	Data     []byte
}
type BlockChain struct {
	Blocks []*Block
}

//main.go
func (b *Block) SetHash() {
	information := bytes.Join([][]byte{ToHexInt(b.Timestamp), b.PrevHash, b.Data}, []byte{})
	hash := sha256.Sum256(information)
	b.Hash = hash[:]
}

func ToHexInt(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
//区块
func CreateBlock(prevhash, data []byte) *Block {
	block := Block{time.Now().Unix(), []byte{}, prevhash, data}
	block.SetHash()
	return &block
}
//创世区块
func GenesisBlock() *Block {
	genesisWords := "Hello, blockchain!"
	return CreateBlock([]byte{}, []byte(genesisWords))
}

func (bc *BlockChain) AddBlock(data string) {
	newBlock := CreateBlock(bc.Blocks[len(bc.Blocks)-1].Hash, []byte(data))
	bc.Blocks = append(bc.Blocks, newBlock)
}
//区块链初始化

func CreateBlockChain() *BlockChain {
	blockchain := BlockChain{}
	blockchain.Blocks = append(blockchain.Blocks, GenesisBlock())
	return &blockchain
}