package chain

import (
	"fmt"
	"io"
	"les-miserables-chain/utils"
	"log"
	"net"
)

// 节点消息列表
const MESSAGE_VERSION = "version"
const MESSAGE_ADDR = "addr"
const MESSAGE_BLOCK = "block"
const MESSAGE_INV = "inv"
const MESSAGE_GETBLOCKS = "getblocks"
const MESSAGE_GETDATA = "getdata"
const MESSAGE_TX = "tx"

//消息处理
func handleMessage(conn net.Conn, bc *Chain) {
	request, err := io.ReadAll(conn)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Receive a Message:%s\n", request[:12])
	message := utils.BytesToMessage(request[:12])
	switch message {
	case MESSAGE_VERSION:
		handleVersion(request, bc)
	case MESSAGE_ADDR:
		handleAddr(request, bc)
	case MESSAGE_BLOCK:
		handleBlock(request, bc)
	case MESSAGE_GETBLOCKS:
		handleGetblocks(request, bc)
	case MESSAGE_GETDATA:
		handleGetData(request, bc)
	case MESSAGE_INV:
		handleInv(request, bc)
	case MESSAGE_TX:
		handleTx(request, bc)
	default:
		fmt.Println("未知的节点消息!")
	}
	defer conn.Close()
}
