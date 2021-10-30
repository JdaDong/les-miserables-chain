package chain

//UTXO交易输出
type TXOutput struct {
	Value        int
	ScriptPubKey string
}

//解锁交易输出
func (out *TXOutput) UnlockOutput(unlockOutputAddress string) bool {
	return out.ScriptPubKey == unlockOutputAddress
}
