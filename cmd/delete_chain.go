package cmd

import (
	"fmt"
	"les-miserables-chain/database"
)

func (cli *CLI) deleteChain() {
	fmt.Println("删除本地区块数据中...")
	_ = database.DeleteDbFile()
}
