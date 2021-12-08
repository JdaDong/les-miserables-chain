package chain

import (
	"bytes"
	"fmt"
	"io"
	"les-miserables-chain/utils"
	"log"
	"net"
)

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
