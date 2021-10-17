package cmd

import (
	"fmt"
	"les-miserables-chain/chain"
)

func (cli *CLI) getBalance(address string) {
	fmt.Println("查询地址:" + address)

	blockchain := chain.BlockchainObject()
	defer blockchain.DB.Close()

	amount := blockchain.GetBalance(address)

	fmt.Printf("查询结果:%d\n", amount)
}
