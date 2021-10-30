package chain

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

//UTXO交易数据
type Transaction struct {
	TxHash    []byte
	TxInputs  []*TXInput
	TxOutputs []*TXOutput
}

//创世区块奖励
const godMoney = 7

//判断是否是创世交易
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.TxInputs) == 1 && tx.TxInputs[0].OutputIndex == -1 && len(tx.TxInputs[0].TxID) == 0
}

//创建coinbase交易
func NewCoinBaseTX(address string) *Transaction {

	//创世输入
	txIn := &TXInput{
		TxID:        []byte{},
		OutputIndex: -1,
		ScriptSig:   nil,
		PublicKey:   []byte{},
	}
	//创世输出
	txOut := NewTxOutput(10, address)
	//创世交易
	tx := Transaction{
		TxHash:    []byte{},
		TxInputs:  []*TXInput{txIn},
		TxOutputs: []*TXOutput{txOut},
	}
	//生成交易hash
	tx.SetTxHash()

	fmt.Println("=======生成创世交易=======")
	fmt.Printf("创世地址:%s", address)
	fmt.Printf("交易号:%x\n", tx.TxHash)
	fmt.Println("交易输入:", tx.TxInputs)
	fmt.Println("交易输出:", tx.TxOutputs)
	fmt.Println("=========================")

	return &tx
}

//转账-创建新的交易
func CreateTransaction(from, to string, amount int, chain *Chain, txs []*Transaction) *Transaction {
	wallets, _ := NewWallets()
	wallet := wallets.WalletMap[from]
	//1.获取刚好能用的金额和合规的UTXO输出
	money, validateUTXO := chain.SpendableUTXOs(from, amount, txs)

	var txInputs []*TXInput
	var txOutputs []*TXOutput

	//2.遍历可用UTXO交易输出
	for txHash, indexArray := range validateUTXO {
		//2.1.构建交易输入
		txHashBytes, _ := hex.DecodeString(txHash)
		for _, index := range indexArray {
			txInput := &TXInput{txHashBytes, index, nil, wallet.GetAddress()}
			txInputs = append(txInputs, txInput)
		}
	}
	//转账
	txOutput := NewTxOutput(amount, to)
	txOutputs = append(txOutputs, txOutput)

	//找零
	txOutput = NewTxOutput(money-amount, from)
	txOutputs = append(txOutputs, txOutput)

	//构建交易
	tx := &Transaction{[]byte{}, txInputs, txOutputs}

	tx.SetTxHash()
	chain.SignTransaction(tx, wallet.PrivateKey)
	return tx

}

//生成交易Hash ID
func (tx *Transaction) SetTxHash() {
	var encoded bytes.Buffer
	var hash [32]byte
	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.TxHash = hash[:]
}

//交易签名
func (tx *Transaction) Sign(privateKey ecdsa.PrivateKey, prevTXs map[string]Transaction) {
	if tx.IsCoinbase() {
		return
	}
	for _, in := range tx.TxInputs {
		if prevTXs[hex.EncodeToString(in.TxID)].TxHash == nil {
			log.Panic("Previous transaction is not correct")
		}
	}
	txCopy :=tx.T
}


func(tx *Transaction)TrimmedCopy()Transaction{
	var inputs []*TXInput
	var outputs []*TXOutput
	for ,
}