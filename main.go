package main

import (
	"fmt"
	"les-miserables-chain/chain"
)

func main() {
	blockchain := chain.NewBlockChain()
	blockchain.AddBlock("test tx")
	fmt.Printf("%x", blockchain.LastHash)
}
