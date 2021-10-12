package cmd

import "fmt"

//打印帮助信息
func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tgetbalance -address ADDRESS - Get balance of ADDRESS")                                     //获取余额
	fmt.Println("\tinit -address ADDRESS - Initialize a blockchain and send genesis block reward to ADDRESS") //初始化区块链
	fmt.Println("\tdelete - Delete local block data")                                                         //清空本地区块数据
	fmt.Println("\tprintchain - Print all the blocks of the blockchain:")                                     //打印区块链
	fmt.Println("\tsend -from FROM -to TO -amount AMOUNT - Send AMOUNT of coins from FROM address to TO")     //转账
}
