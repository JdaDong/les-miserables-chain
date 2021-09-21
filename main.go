package main

import (
	"les-miserables-chain/chain"
	"les-miserables-chain/cmd"
)

func main() {
	blockchain := chain.NewBlockChain()
	cli := cmd.CLI{
		Chain: blockchain,
	}
	cli.Run()
}
