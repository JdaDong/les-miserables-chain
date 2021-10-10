package chain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
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

//创世区块奖励
const godMoney = 7

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && tx.Inputs[0].OutputIndex == -1 && len(tx.Inputs[0].TxID) == 0
}
func NewTransaction(from, to string, amount int, chain *Chain) *Transaction {

	//创建输入
	var inputs []TXInput
	//创建输出
	var outputs []TXOutput

	//获取未消费的输出
	acc, spendableOutputs := chain.FindSpendableOutputs(from, amount)
	fmt.Println(spendableOutputs)
	//额度小于转账金额
	if acc < amount {
		log.Panic("Not enough funds")
	}
	//遍历可用的输出
	for txid, outs := range spendableOutputs {
		//txid为交易的hash
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}
		//遍历未花费输出
		for _, out := range outs {
			//交易输入构建
			input := TXInput{
				TxID:        txID,
				OutputIndex: out,
				ScriptSig:   from,
			}
			inputs = append(inputs, input)
		}
	}
	//交易输出
	output := TXOutput{
		Value:        amount,
		ScriptPubKey: to,
	}
	outputs = append(outputs, output)

	//找零
	change := TXOutput{
		Value:        acc - amount,
		ScriptPubKey: from,
	}
	outputs = append(outputs, change)
	//构建交易结构体
	tx := Transaction{nil, inputs, outputs}
	tx.SetIndex()
	fmt.Println("From:", from)
	fmt.Println("To:", to)
	fmt.Println("交易号:", tx.Index)
	fmt.Println("交易输入:", tx.Inputs)
	fmt.Println("交易输出:", tx.Outputs)
	return &tx
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
	fmt.Println("To:", to)
	fmt.Println("交易号:", tx.Index)
	fmt.Println("交易输入:", tx.Inputs)
	fmt.Println("交易输出:", tx.Outputs)
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
