package cmd

import (
	"fmt"
)

//转账
func (cli *CLI) sendToken(from, to string, amount int) {
	fmt.Printf("转账来源：%s\n 转账目标：%s\n 转账金额：%d\n", from, to, amount)
}
