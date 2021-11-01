package chain

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
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
	txOut := NewTxOutput(godMoney, address)
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
	wallets, err := NewWallets()
	if err != nil {
		log.Panic(err)
	}
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
			txInput := &TXInput{txHashBytes, index, nil, wallet.PublicKey}
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
	//交易签名
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
	//如果是创世交易，不进行签名
	if tx.IsCoinbase() {
		return
	}
	//判断交易是否为空
	for _, in := range tx.TxInputs {
		if prevTXs[hex.EncodeToString(in.TxID)].TxHash == nil {
			log.Panic("Previous transaction is not correct")
		}
	}
	txCopy := tx.TransactionCopy()
	for inID, vin := range txCopy.TxInputs {
		prevTXs := prevTXs[hex.EncodeToString(vin.TxID)]
		txCopy.TxInputs[inID].ScriptSig = nil
		txCopy.TxInputs[inID].PublicKey = prevTXs.TxOutputs[vin.OutputIndex].ScriptPubKey
		txCopy.TxHash = txCopy.Hash()
		txCopy.TxInputs[inID].PublicKey = nil
		r, s, err := ecdsa.Sign(rand.Reader, &privateKey, txCopy.TxHash)
		if err != nil {
			log.Panic(err)
		}
		signature := append(r.Bytes(), s.Bytes()...)
		tx.TxInputs[inID].ScriptSig = signature
	}
}

//拷贝新的Transaction用于数字签名
func (tx *Transaction) TransactionCopy() Transaction {
	var inputs []*TXInput
	var outputs []*TXOutput

	for _, vin := range tx.TxInputs {
		inputs = append(inputs, &TXInput{vin.TxID, vin.OutputIndex, nil, nil})
	}

	for _, vout := range tx.TxOutputs {
		outputs = append(outputs, &TXOutput{vout.Value, vout.ScriptPubKey})
	}
	txCopy := Transaction{tx.TxHash, inputs, outputs}

	return txCopy
}

func (tx *Transaction) Hash() []byte {
	txCopy := tx

	txCopy.TxHash = []byte{}

	hash := sha256.Sum256(txCopy.Serialize())
	return hash[:]
}

func (tx *Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}
