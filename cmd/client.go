package cmd

import (
	"flag"
	"fmt"
	"github.com/boltdb/bolt"
	"les-miserables-chain/chain"
	"les-miserables-chain/database"
	"les-miserables-chain/utils"
	"log"
	"math/big"
	"os"
)

type CLI struct {
	Chain *chain.Chain
}

//打印帮助信息
func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tgetbalance -address ADDRESS - Get balance of ADDRESS")                                 //获取余额
	fmt.Println("\tinit -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS") //初始化区块链
	fmt.Println("\tprintchain - Print all the blocks of the blockchain:")                                 //
	fmt.Println("\tsend -from FROM -to TO -amount AMOUNT - Send AMOUNT of coins from FROM address to TO")
}

//校验参数
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

//打印链信息
func (cli *CLI) printChain() {
	var blockchainIterator *chain.ChainIterator
	blockchainIterator = cli.Chain.Iterator()

	var hashInt big.Int

	for {
		err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(database.DbFile))
			blockBytes := b.Get(blockchainIterator.CurrentHash)
			block := chain.DeserializeBlock(blockBytes)
			fmt.Printf("Transaction：%v \n", block.Transactions)
			fmt.Printf("PrevBlockHash：%x \n", block.BlockPreHash)
			fmt.Printf("Timestamp：%s \n", utils.ConvertToTime(block.BlockTimestamp/1e3))
			fmt.Printf("Hash：%x \n", block.BlockCurrentHash)
			fmt.Printf("Nonce：%d \n", block.BlockNonce)
			fmt.Println()
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
		blockchainIterator = blockchainIterator.Next()
		hashInt.SetBytes(blockchainIterator.CurrentHash)
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
}

//转账
func (cli *CLI) sendToken(data string) {
	fmt.Println(data)
	tx1 := chain.NewTransaction("levy", "page1", 1, cli.Chain)
	tx2 := chain.NewTransaction("levy", "page2", 2, cli.Chain)
	cli.Chain.MineBlock([]*chain.Transaction{tx1, tx2})
}

func (cli *CLI) addBlock(data string) {
	cli.sendToken(data)
}

func (cli *CLI) Run() {
	cli.validateArgs()
	CmdAddBlock := flag.NewFlagSet("addblock", flag.ExitOnError)
	CmdPrintChain := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := CmdAddBlock.String("data", "转账中", "Block data")
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
	if CmdPrintChain.Parsed() {
		cli.printChain()
	}
}
