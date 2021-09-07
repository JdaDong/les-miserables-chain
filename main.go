package main

import (
	"fmt"
	"les-miserables-chain/chain_basic"
)

func main() {
	blockchain := chain_basic.NewBlockChain()
	blockchain.AddBlock("Send 20 lmc token To Levy From Genesis")
	blockchain.AddBlock("Send 20 lmc token To Levy From Genesis")
	blockchain.AddBlock("Send 20 lmc token To Levy From Genesis")
	blockchain.AddBlock("Send 20 lmc token To Levy From Genesis")
	blockchain.AddBlock("Send 20 lmc token To Levy From Genesis")
	blockchain.AddBlock("Send 20 lmc token To Levy From Genesis")

	for _, block := range blockchain.Blocks {
		fmt.Printf("区块时间：%d\n", block.BlockTimestamp)
		fmt.Printf("父区块：%x\n", block.BlockPreHash)
		fmt.Printf("当前区块：%x\n", block.BlockCurrentHash)
		fmt.Printf("Data：%s\n\n", string(block.BlockData))
		fmt.Printf("Nonce: %d\n", block.BlockNonce)
	}
}
