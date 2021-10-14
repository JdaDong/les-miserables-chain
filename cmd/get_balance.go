package cmd

import (
	"fmt"
)

func (cli *CLI) getBalance(address string) {
	fmt.Println("查询地址为：", address)
	fmt.Println(cli.Chain.GetBalance(address))
}
