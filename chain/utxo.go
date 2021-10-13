package chain

import (
	"encoding/hex"
	"math/big"
)

//交易输出UXTO模型
type UTXO struct {
	TxHash []byte    //交易hash
	Index  int       //输出索引
	OutPut *TXOutput //交易输出
}

//获取未花费交易的UXTO信息
func (chain *Chain) UnUTXOs(address string) []*UTXO {
	var unUTXOos []*UTXO
	spentTXOutputs := make(map[string][]int)

	blockIterator := chain.Iterator()

	for {
		block := blockIterator.NextBlock()

		for _, tx := range block.Transactions {
			if tx.IsCoinbase() == false {
				for _, in := range tx.Inputs {
					if in.UnlockInput(address) {
						key := hex.EncodeToString(in.TxID)
						spentTXOutputs[key] = append(spentTXOutputs[key], in.OutputIndex)
					}
				}
			}
			for index, out := range tx.Outputs {
				if out.UnlockOutput(address) {
					if len(spentTXOutputs) != 0 {
						for txHash, indexArry := range spentTXOutputs {
							for _, i := range indexArry {
								if index == i && txHash == hex.EncodeToString(tx.Index) {
									continue
								} else {
									utxo := &UTXO{
										TxHash: tx.Index,
										Index:  index,
										OutPut: &out,
									}
									unUTXOos = append(unUTXOos, utxo)
								}
							}
						}
					} else {
						utxo := &UTXO{
							TxHash: tx.Index,
							Index:  index,
							OutPut: &out,
						}
						unUTXOos = append(unUTXOos, utxo)
					}

				}
			}
		}
		var hashInt big.Int
		hashInt.SetBytes(block.BlockPreHash)
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
	return unUTXOos

}
