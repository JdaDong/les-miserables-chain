package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func CmdClient() {
	CmdAddBlock := flag.NewFlagSet("addblock", flag.ExitOnError)
	CmdPrintChain := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := CmdAddBlock.String("data", "", "BLock data")
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
	default:
		fmt.Println("Nothing to do")
	}
	if CmdAddBlock.Parsed() {
		if *addBlockData == "" {
			CmdAddBlock.Usage()
			os.Exit(1)
		}
		fmt.Println("Data:" + *addBlockData)
	}
	if CmdPrintChain.Parsed() {
		fmt.Println("printchain!")
	}
}
