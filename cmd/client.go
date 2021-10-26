package cmd

import (
	"flag"
	"les-miserables-chain/chain"
	"log"
	"os"
)

type CLI struct {
	Chain *chain.Chain
}

//校验参数
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()
	CmdPrintChain := flag.NewFlagSet("printchain", flag.ExitOnError) //打印区块链
	CmdDelete := flag.NewFlagSet("delete", flag.ExitOnError)         //删除区块链
	CmdInit := flag.NewFlagSet("init", flag.ExitOnError)             //初始化区块链
	CmdGetBalance := flag.NewFlagSet("balance", flag.ExitOnError)    //获取账户余额
	CmdSendToken := flag.NewFlagSet("send", flag.ExitOnError)        //转账
	CmdCreateWallet := flag.NewFlagSet("a", flag.ExitOnError)

	cbAddr := CmdInit.String("address", "", "创世区块奖励人")
	balanceAddr := CmdGetBalance.String("addr", "", "获取指定地址的余额")
	sendFrom := CmdSendToken.String("from", "", "转账源地址")
	sendTo := CmdSendToken.String("to", "", "转账目的地址")
	sendAmount := CmdSendToken.Int("amount", 0, "转账金额")

	switch os.Args[1] {
	case "printchain":
		err := CmdPrintChain.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "init":
		err := CmdInit.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "delete":
		err := CmdDelete.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := CmdGetBalance.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := CmdSendToken.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createwallet":
		err := CmdCreateWallet.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}
	if CmdInit.Parsed() {
		if *cbAddr == "" {
			cli.printUsage()
			os.Exit(1)
		}
		cli.initialize(*cbAddr)
	}
	if CmdPrintChain.Parsed() {
		cli.printChain()
	}
	if CmdDelete.Parsed() {
		cli.deleteChain()
	}
	if CmdGetBalance.Parsed() {
		//fmt.Println(*balanceAddr)
		if *balanceAddr == "" {
			cli.printUsage()
			os.Exit(1)
		}
		cli.getBalance(*balanceAddr)
	}
	if CmdSendToken.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount == 0 {
			cli.printUsage()
			os.Exit(1)
		}
		cli.sendToken(*sendFrom, *sendTo, *sendAmount)
	}
	if CmdCreateWallet.Parsed() {
		cli.createWallet()
	}
}
