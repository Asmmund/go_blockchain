package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain
type BlockChain struct {
	blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Block
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func Genesis() *Block {
	return CreateBlock("Antonio", []byte{})
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// genral functions
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("A second block")
	chain.AddBlock("A third block")
	chain.AddBlock("A forth block")

	for _, block := range chain.blocks {
		fmt.Printf("Data in block %s\n", block.Data)
		fmt.Printf("hash in block %x\n", block.Hash)
	}
}
