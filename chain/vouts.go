package chain

import (
	"bytes"
	"encoding/gob"
	"log"
)

type TXOutputs struct {
	UTXOS []*UTXO
}

// 将区块序列化成字节数组
func (txOutputs *TXOutputs) Serialize() []byte {

	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(txOutputs)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// 反序列化
func (txOutputs *TXOutputs) Deserialize(txOutputsBytes []byte) *TXOutputs {

	var result TXOutputs

	decoder := gob.NewDecoder(bytes.NewReader(txOutputsBytes))
	err := decoder.Decode(&result)
	if err != nil {
		log.Panic(err)
	}

	return &result
}
