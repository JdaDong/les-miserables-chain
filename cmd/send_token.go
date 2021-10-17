package cmd

import (
	"fmt"
)

//转账
func (cli *CLI) sendToken(from, to string, amount int) {
	fmt.Printf("转账来源：%s\n 转账目标：%s\n 转账金额：%d\n", from, to, amount)
}

// 多签名转账
//func (cli *CLI) send(from []string, to []string, amount []string) {
//
//	if DBExists() == false {
//		fmt.Println("数据不存在.......")
//		os.Exit(1)
//	}
//
//	blockchain := BlockchainObject()
//	defer blockchain.DB.Close()
//
//	blockchain.MineNewBlock(from, to, amount)
//
//}
