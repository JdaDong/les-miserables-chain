package chain

import (
	"bytes"
	"les-miserables-chain/utils"
)

//UTXO交易输出
type TXOutput struct {
	Value        int
	ScriptPubKey []byte
}

////解锁交易输出
//func (out *TXOutput) UnlockOutput(unlockOutputAddress string) bool {
//	return out.ScriptPubKey == unlockOutputAddress
//}

//交易输出锁定
func (out *TXOutput) Lock(address string) {
	publicKeyHash := utils.Base58Decode([]byte(address))
	out.ScriptPubKey = publicKeyHash[1 : len(publicKeyHash)-4] //公钥数据部分
}

//新建交易输出
func NewTxOutput(value int, address string) *TXOutput {
	txOutput := &TXOutput{
		Value:        value,
		ScriptPubKey: nil,
	}
	txOutput.Lock(address)
	return txOutput
}

//解锁
func (out *TXOutput) UnLockScriptPubKeyWithAddress(address string) bool {
	publicKeyHash := utils.Base58Decode([]byte(address))
	publicKeyRipemd160Hash := publicKeyHash[1 : len(publicKeyHash)-4]
	return bytes.Compare(out.ScriptPubKey, publicKeyRipemd160Hash) == 0
}
