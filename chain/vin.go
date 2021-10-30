package chain

import (
	"bytes"
	"les-miserables-chain/utils"
)

//UTXO交易输入
type TXInput struct {
	TxID        []byte //交易hash id
	OutputIndex int    //交易输出索引
	ScriptSig   string //交易输入-数字签名
	PublicKey   []byte //公钥
}

////解锁交易输入
//func (in *TXInput) UnlockInput(unlockInputAddress string) bool {
//	return in.ScriptSig == unlockInputAddress
//}

//解锁交易输入
func (in *TXInput) UnlockPublicKeyHash(publicKeyHash []byte) bool {
	publicKeyRipemd160Hash := utils.GetRipemd160(in.PublicKey)
	return bytes.Compare(publicKeyRipemd160Hash, publicKeyHash) == 0
}
