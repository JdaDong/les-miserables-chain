package chain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

//UTXO交易数据
type Transaction struct {
	Index   []byte
	Inputs  []TXInput
	Outputs []TXOutput
}

//UTXO交易输入
type TXInput struct {
	TxID        []byte
	OutputIndex int
	ScriptSig   string
}

//UTXO交易输出
type TXOutput struct {
	Value        int
	ScriptPubKey string
}

const godMoney = 7

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && tx.Inputs[0].OutputIndex == -1 && len(tx.Inputs[0].TxID) == 0
}

//coinbase交易
func NewCoinBaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	//创世输入
	txin := TXInput{
		TxID:        []byte{},
		OutputIndex: -1,
		ScriptSig:   data,
	}
	//创世输出
	txout := TXOutput{
		Value:        godMoney,
		ScriptPubKey: to,
	}
	//创世交易
	tx := Transaction{
		Index:   nil,
		Inputs:  []TXInput{txin},
		Outputs: []TXOutput{txout},
	}
	tx.SetIndex()
	return &tx
}

//生成交易ID
func (tx *Transaction) SetIndex() {
	var encoded bytes.Buffer
	var hash [32]byte
	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.Index = hash[:]
}

//解锁交易输入
func (in *TXInput) UnlockInput(unlockInputAddress string) bool {
	return in.ScriptSig == unlockInputAddress
}

//解锁交易输出
func (out *TXOutput) UnlockOutput(unlockOutputAddress string) bool {
	return out.ScriptPubKey == unlockOutputAddress
}
