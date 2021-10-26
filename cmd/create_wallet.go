package cmd

import (
	"fmt"
	"les-miserables-chain/chain"
)

func (cli *CLI) createWallet() {
	wallets, _ := chain.NewWallets()
	wallets.CreateNewWallet()
	fmt.Println(wallets.WalletMap)
}
