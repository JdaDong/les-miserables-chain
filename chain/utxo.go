package chain

//交易输出UXTO模型
type UTXO struct {
	TxHash []byte    //交易hash
	Index  int       //输出索引
	OutPut *TXOutput //交易输出
}
