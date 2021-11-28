package cmd

import "fmt"

//打印帮助信息
func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tgetbalance -addr ADDRESS -- 获取指定地址的余额")
	fmt.Println("\tinit -address ADDRESS -- 初始化区块链数据")
	fmt.Println("\tdelete -- 删除区块链数据")
	fmt.Println("\tprintchain -- 打印所有区块信息")
	fmt.Println("\tsend -from FROM -to TO -amount AMOUNT -- 转账交易")
	fmt.Println("\tcreatewallet -- 创建钱包")
	fmt.Println("\taddresslists -- 获取所有钱包的地址集合")
	fmt.Println("\tstartnode -miner MINERADDR -- 启动节点网络服务")
}
