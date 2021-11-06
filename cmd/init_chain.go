package cmd

import (
	"fmt"
	"les-miserables-chain/chain"
)

//初始化区块链
func (cli *CLI) initialize(address string) {
	fmt.Println("初始化区块链中...")
	blockchain := chain.InitBlockChain(address)
	defer blockchain.DB.Close()
	utxoRecord := &chain.UTXORecord{Blockchain: blockchain}
	utxoRecord.ResetUTXORecord()
}
