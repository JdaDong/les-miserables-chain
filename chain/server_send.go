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

//获取区块
func sendGetBlocks(toAddress string) {
	payload := utils.GobEncode(GetBlocks{nodeAddress})
	request := append(utils.MessageTobytes(MESSAGE_GETBLOCKS), payload...)
	sendMessage(toAddress, request)
}

// 主节点将自己的所有的区块hash发送给钱包节点
func sendInv(toAddress string, kind string, hashes [][]byte) {

	payload := utils.GobEncode(Inv{nodeAddress, kind, hashes})

	request := append(utils.MessageTobytes(MESSAGE_INV), payload...)

	sendMessage(toAddress, request)

}

func sendBlock(toAddress string, block *Block) {
	payload := utils.GobEncode(BlockData{nodeAddress, block})
	request := append(utils.MessageTobytes(MESSAGE_BLOCK), payload...)
	sendMessage(toAddress, request)
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
