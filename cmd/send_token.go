package cmd

import (
	"fmt"
	"les-miserables-chain/chain"
)

//转账
func (cli *CLI) sendToken(data string) {
	fmt.Println(data)
	tx1 := chain.NewTransaction("levy", "page1", 1, cli.Chain)
	tx2 := chain.NewTransaction("levy", "page2", 2, cli.Chain)
	cli.Chain.MineBlock([]*chain.Transaction{tx1, tx2})
}
