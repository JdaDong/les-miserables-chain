package cmd

import (
	"fmt"
	"les-miserables-chain/chain"
	"les-miserables-chain/database"
	"les-miserables-chain/utils"
	"log"
	"math/big"

	"github.com/boltdb/bolt"
)

//打印链信息
func (cli *CLI) printChain() {
	//判断区块链是否已经初始化
	if !database.DbExist() {
		fmt.Println("您需要先初始化区块链!")
		cli.printUsage()
		return
	}
	cli.Chain = chain.BlockchainObject()
	defer cli.Chain.DB.Close()
	fmt.Println("Hello1")
	var blockchainIterator *chain.ChainIterator
	blockchainIterator = cli.Chain.Iterator()
	fmt.Println("Hello2")
	//fmt.Println(blockchainIterator)
	var hashInt big.Int

	for {
		err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(database.BlockBucket))
			blockBytes := b.Get(blockchainIterator.CurrentHash)
			block := chain.DeserializeBlock(blockBytes)
			fmt.Printf("Coinbase Address：%v \n", block.Transactions[0].TxOutputs[0].ScriptPubKey)
			fmt.Printf("Transactions：%v\n", block.Transactions)
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
