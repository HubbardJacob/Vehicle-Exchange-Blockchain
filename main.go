package main

import (
	"fmt"
	"math/big"
)

func main() {
	// chain := blockchain.InitBlockChain()

	// chain.AddBlock("First Block")
	// chain.AddBlock("Second Block")
	// chain.AddBlock("Third Block")

	// for _, block := range chain.blocks {
	// 	fmt.Printf("Previous Hash: %x\n", block.PrevHash)
	// 	fmt.Printf("Data in Block: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// }

	target := big.NewInt(1)
	fmt.Println(target.Lsh(target, uint(256-250)))
}
