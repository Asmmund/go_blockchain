package main

import (
	"fmt"
	"github.com/Asmmund/go_blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("A second block")
	chain.AddBlock("A third block")
	chain.AddBlock("A forth block")

	for _, block := range chain.Blocks {
		fmt.Printf("Data in block %s\n", block.Data)
		fmt.Printf("hash in block %x\n", block.Hash)
	}
}
