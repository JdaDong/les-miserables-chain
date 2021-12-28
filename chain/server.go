package chain

import (
	"fmt"
	"log"
	"net"
)

var knowNodes = []string{"localhost:3000"} //3000主节点地址

var nodeAddress string //节点地址

var transactionArry [][]byte

func StartServer(nodeID string, miner string) {
	nodeAddress = fmt.Sprintf("localhost:%s", nodeID)
	listener, err := net.Listen("tcp", nodeAddress)
	fmt.Println(nodeAddress)
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()
	bc := BlockchainObject()
	//非主节点，需要同步
	if nodeAddress != knowNodes[0] {
		sendVersion(knowNodes[0], bc)
	}
	//接受客户端消息
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleMessage(conn, bc)
	}

}
