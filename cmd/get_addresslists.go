package cmd

import (
	"fmt"
	"les-miserables-chain/chain"
	"log"
)

func (cli *CLI) addresslists() {
	fmt.Println("当前存在的钱包地址集合为:")

	wallets, err := chain.NewWallets()
	if err != nil {
		log.Panic(err)
	}
	for address, _ := range wallets.WalletMap {
		fmt.Println(address)
	}
}
