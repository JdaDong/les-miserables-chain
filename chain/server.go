package chain

import (
	"bytes"
	"fmt"
	"io"
	"les-miserables-chain/utils"
	"log"
	"net"
)

var knowNodes = []string{"localhost:3000"} //3000主节点地址

var nodeAddress string //节点地址

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
		//request, err := io.ReadAll(conn)
		//if err != nil {
		//	log.Panic(err)
		//}
		//fmt.Printf("Receive a Message:%s\n", request)
		go handleMessage(conn, bc)
	}

}

//发送版本信息
func sendVersion(toAddress string, bc *Chain) {
	bestHeight := bc.GetHighestHeight() //硬编码
	payload := utils.GobEncode(Version{
		Version:    1, //节点版本 硬编码为1
		BestHeight: bestHeight,
		AddrFrom:   nodeAddress,
	})
	requestMsg := append(utils.MessageTobytes("version"), payload...)
	sendMessage(toAddress, requestMsg)
}

//客户端向服务器发送消息
func sendMessage(to string, msg []byte) {
	fmt.Println("客户端向服务器发送数据.......")
	conn, err := net.Dial("tcp", to)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	_, err = io.Copy(conn, bytes.NewReader([]byte(msg)))
	if err != nil {
		log.Panic(err)
	}
}

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
}

func handleVersion(request []byte, bc *Chain) {

}

func handleAddr(request []byte, bc *Chain) {

}

func handleGetblocks(request []byte, bc *Chain) {

}

func handleGetData(request []byte, bc *Chain) {

}

func handleBlock(request []byte, bc *Chain) {

}

func handleTx(request []byte, bc *Chain) {

}

func handleInv(request []byte, bc *Chain) {

}
