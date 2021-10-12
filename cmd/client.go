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

func (cli *CLI) addBlock(data string) {
	cli.sendToken(data)
}

func (cli *CLI) Run() {
	cli.validateArgs()
	CmdAddBlock := flag.NewFlagSet("addblock", flag.ExitOnError)
	CmdPrintChain := flag.NewFlagSet("printchain", flag.ExitOnError)
	CmdDelete := flag.NewFlagSet("delete", flag.ExitOnError)
	addBlockData := CmdAddBlock.String("data", "转账中", "Block data")
	CmdInit := flag.NewFlagSet("init", flag.ExitOnError)
	cbAddr := CmdInit.String("address", "", "coinbase address")
	switch os.Args[1] {
	case "addblock":
		err := CmdAddBlock.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
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
	default:
		cli.printUsage()
		os.Exit(1)
	}
	if CmdAddBlock.Parsed() {
		if *addBlockData == "" {
			cli.printUsage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
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
}
