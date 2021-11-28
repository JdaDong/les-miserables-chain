package cmd

import (
	"fmt"
	"les-miserables-chain/chain"
	"os"
)

func (cli *CLI) startNode(nodeID string, miner string) {
	if miner == "" || chain.CheckAddress([]byte(miner)) {
		fmt.Printf("节点服务启动:127.0.0.1:%s\n", nodeID)
		chain.StartServer(nodeID, miner)
	} else {
		fmt.Println("请输入合法的矿工地址！")
		os.Exit(1)
	}
}
