package main

import (
	"fmt"
	"yanhuangpai/go-u1/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	// Example: Adding a block
	bc.AddBlock([]blockchain.Transaction{
		{Sender: "Alice", Receiver: "Bob", Amount: 50},
	})

	// Print the blockchain
	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %s\n", block.PrevBlockHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}
}
