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
	Inputs  []*TXInput
	Outputs []*TXOutput
}

//创世区块奖励
const godMoney = 7

//判断是否是创世交易
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && tx.Inputs[0].OutputIndex == -1 && len(tx.Inputs[0].TxID) == 0
}

//转账-创建新的交易
func CreateTransaction(from, to string, amount int, chain *Chain, txs []*Transaction) *Transaction {
	//1.获取刚好能用的金额和合规的UTXO输出
	money, validateUTXO := chain.SpendableUTXOs(from, amount, txs)

	var txInputs []*TXInput
	var txOutputs []*TXOutput

	//2.遍历可用UTXO交易输出
	for txHash, indexArray := range validateUTXO {
		//2.1.构建交易输入
		txHashBytes, _ := hex.DecodeString(txHash)
		for _, index := range indexArray {
			txInput := &TXInput{txHashBytes, index, from}
			txInputs = append(txInputs, txInput)
		}
	}
	//转账
	txOutput := &TXOutput{amount, to}
	txOutputs = append(txOutputs, txOutput)

	//找零
	txOutput = &TXOutput{money - amount, from}
	txOutputs = append(txOutputs, txOutput)

	//构建交易
	tx := &Transaction{[]byte{}, txInputs, txOutputs}

	tx.SetIndex()

	return tx

}

//创建coinbase交易
func NewCoinBaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	//创世输入
	txin := &TXInput{
		TxID:        []byte{},
		OutputIndex: -1,
		ScriptSig:   data,
	}
	//创世输出
	txout := &TXOutput{
		Value:        godMoney,
		ScriptPubKey: to,
	}
	//创世交易
	tx := Transaction{
		Index:   nil,
		Inputs:  []*TXInput{txin},
		Outputs: []*TXOutput{txout},
	}
	tx.SetIndex()
	fmt.Println("To:", to)
	fmt.Printf("交易号:%x\n", tx.Index)
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
