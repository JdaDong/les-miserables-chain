package chain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func handleVersion(request []byte, bc *Chain) {
	var buff bytes.Buffer
	var payload Version
	dataBytes := request[12:]

	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	highestHeight := bc.GetHighestHeight()
	foreignerHighestHeight := payload.BestHeight
	if highestHeight > foreignerHighestHeight {
		sendVersion(payload.AddrFrom, bc)
	} else if highestHeight < foreignerHighestHeight {
		sendGetBlocks(payload.AddrFrom)
	}
}

func handleAddr(request []byte, bc *Chain) {

}

func handleGetblocks(request []byte, bc *Chain) {
	var buff bytes.Buffer
	var payload Version
	dataBytes := request[12:]

	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}
	blocks := bc.GetBlockHashes()
	sendInv(payload.AddrFrom, "block", blocks)
}

func handleGetData(request []byte, bc *Chain) {
	var buff bytes.Buffer
	var payload GetData

	dataBytes := request[12:]
	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}
	if payload.Type == "block" {
		block, err := bc.GetBock([]byte(payload.Hash))
		if err != nil {
			return
		}
		sendBlock(payload.AddrFrom, block)
	}
	if payload.Type == "tx" {

	}

}

func handleBlock(request []byte, bc *Chain) {
	var buff bytes.Buffer
	var payload BlockData
	dataBytes := request[12:]

	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}
	block := payload.Block
	err = bc.AddBlock(block)
	if err != nil {
		log.Panic(err)
	}
	if len(transactionArry) > 0 {
		sendGetData(payload.AddrFrom, "block", transactionArry[0])
		transactionArry = transactionArry[1:]
	} else {
		fmt.Println("数据库重置......")
		UTXOSet := &UTXORecord{bc}
		UTXOSet.ResetUTXORecord()
	}

}

func handleTx(request []byte, bc *Chain) {

}

func handleInv(request []byte, bc *Chain) {
	var buff bytes.Buffer
	var payload Inv

	dataBytes := request[12:]
	buff.Write(dataBytes)
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	if payload.Type == "block" {
		blockHash := payload.Items[0]
		sendGetData(payload.AddrFrom, "block", blockHash)
		if len(payload.Items) >= 1 {
			transactionArry = payload.Items[1:]
		}
	}
	if payload.Type == "tx" {

	}

}
