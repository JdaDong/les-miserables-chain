package chain

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

var knowNodes = []string{"localhost:3300"} //3300主节点地址

var nodeAddress string //节点地址

func StartServer(nodeID string, miner string) {
	nodeAddress = fmt.Sprintf("localhost:%s", nodeID)
	listener, err := net.Listen("tcp", nodeAddress)
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()
	//非主节点，需要通阿伯
	if nodeAddress != knowNodes[0] {
		sendMessage(knowNodes[0], nodeAddress)
	}
	//接受客户端消息
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}
		request, err := io.ReadAll(conn)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("Receive a Message:%s\n", request)
	}

}

func sendMessage(to string, from string) {
	fmt.Println("客户端向服务器发送数据.......")
	conn, err := net.Dial("tcp", to)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	_, err = io.Copy(conn, bytes.NewReader([]byte(from)))
	if err != nil {
		log.Panic(err)
	}
}
