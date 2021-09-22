package chain

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
