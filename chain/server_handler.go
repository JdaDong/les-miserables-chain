package chain

import (
	"bytes"
	"encoding/gob"
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

	}
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
