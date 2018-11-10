package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// BlockChain
type BlockChain struct {
	Blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
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
