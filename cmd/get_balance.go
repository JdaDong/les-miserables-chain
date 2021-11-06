package cmd

import (
	"fmt"
	"les-miserables-chain/chain"
)

func (cli *CLI) getBalance(address string) {
	fmt.Println("查询地址:" + address)

	blockchain := chain.BlockchainObject()
	defer blockchain.DB.Close()

	utxoRecord := &chain.UTXORecord{Blockchain: blockchain}
	amount := utxoRecord.GetBalance(address)

	fmt.Printf("查询结果:%d\n", amount)
}
