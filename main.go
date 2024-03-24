package main

import (
	"fmt"
	"strconv"

	blockchain "github.com/siri-n-shetty/Blockchain-Golang/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	// we are initializing the blockchain
	// we are adding blocks to the chain

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	// to see the blockhain
	for _, block := range chain.Blocks {
		// we are iterating over the blocks in the chain
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		// we are printing the previous hash
		fmt.Printf("Data in Block: %s\n", block.Data)
		// we are printing the data in the block
		fmt.Printf("Hash: %x\n", block.Hash)
		// we are printing the hash of the block

		pow := blockchain.NewProof(block)
		// we are creating a new proof of work with the block
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		// we are printing the proof of work
		fmt.Println()
		// we are printing a new line
	}
}
