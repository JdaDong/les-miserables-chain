package cmd

import (
	"fmt"
	"les-miserables-chain/chain"
	"les-miserables-chain/database"
	"os"
)

//转账
func (cli *CLI) sendToken(from []string, to []string, amount []string, mineNow bool) {
	if database.DbExist() == false {
		fmt.Println("请先初始化区块链!")
		os.Exit(1)
	}
	if mineNow {
		fmt.Println("启动当前节点验证...")
		blockchain := chain.BlockchainObject()
		defer blockchain.DB.Close()

		_ = blockchain.MineBlock(from, to, amount)
		utxoRecord := &chain.UTXORecord{blockchain}
		utxoRecord.Update()
	} else {
		fmt.Println("启动矿工节点验证...")
	}

}
