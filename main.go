package main

import (
	"flag"
	"fmt"
	"github.com/Asmmund/go_blockchain/blockchain"
	"os"
	"runtime"
	"strconv"
)

type CommandLine struct {
	blockchain *blockchain.BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" Add -block BLOCK_DATA - add a block to the chain")
	fmt.Println(" print - prints the blocks in the chain")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Added Block!")
}

func PrintoutBlockInfo(block *Block) {
	fmt.Printf("previous Hash in block %s\n", block.PrevHash)
	fmt.Printf("Data in block %s\n", block.Data)
	fmt.Printf("hash in block %x\n", block.Hash)

	pow := blockchain.NewProof(block)
	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	fmt.Println()
	return
}

func (cli *CommandLine) run() {
	cli.validateArgs()

	addBlockCmd := flag.newFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.newFlagSet("print", flag.ExitOnError)
	addBlockData := flag.newFlagSet("block", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

}

func (cli *CommandLine) printChain() {
	iter := cli.blockchain.Iterator()

	for {
		block := iter.Next

		PrintoutBlockInfo(block)
		if len(block.PrevHash) == 0 {
			break
		}
	}

}

func main() {
	defer os.Exit(0)

	chain := blockchain.InitBlockChain()
	defer chain.Database.Close()

	cli := CommandLine{chain}

	cli.run()
}
