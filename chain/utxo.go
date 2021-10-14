package chain

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
)

//交易输出UXTO模型
type UTXO struct {
	TxHash []byte    //交易hash
	Index  int       //输出索引
	OutPut *TXOutput //交易输出
}

//获取未花费交易的UXTO信息
func (chain *Chain) UnUTXOs(address string) []*UTXO {
	//UTXO交易输出集合
	var unUTXOos []*UTXO
	//已花费交易输出
	spentTXOutputs := make(map[string][]int)

	//区块链迭代器
	blockIterator := chain.Iterator()

	for {
		//从最新区块开始遍历
		block := blockIterator.NextBlock()
		//遍历区块交易数据
		for _, tx := range block.Transactions {
			//遍历非创世交易
			if tx.IsCoinbase() == false {
				//遍历交易输入
				for _, in := range tx.Inputs {
					if in.UnlockInput(address) {
						key := hex.EncodeToString(in.TxID)
						spentTXOutputs[key] = append(spentTXOutputs[key], in.OutputIndex)
					}
				}
			}
			//遍历交易输出
			for index, out := range tx.Outputs {
				if out.UnlockOutput(address) {
					if len(spentTXOutputs) != 0 {
						//遍历已花费交易输出
						for txHash, indexArry := range spentTXOutputs {
							for _, i := range indexArry {
								//如果索引相等，表示某该地址的交易输出
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
		//遍历到0区块为止
		var hashInt big.Int
		hashInt.SetBytes(block.BlockPreHash)
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
	//返回所有交易输出集合
	return unUTXOos

}

func (chain *Chain) SpendableUTXOs(from string, amount int) (int, map[string][]int) {
	//获取转账源地址的未花费交易输出
	utxos := chain.UnUTXOs(from)

	//可用UTXO的map
	spendableUTXO := make(map[string][]int)

	//余额
	var value int

	//遍历未花费UTXO交易输出，一旦余额充足就不再遍历
	for _, utxo := range utxos {
		value += utxo.OutPut.Value
		hash := hex.EncodeToString(utxo.TxHash)
		spendableUTXO[hash] = append(spendableUTXO[hash], utxo.Index)
		if value >= amount {
			break
		}
	}
	//遍历完后，余额还是不足，退出
	if value < amount {
		fmt.Printf("%s地址中的余额不足\n", from)
		os.Exit(1)
	}
	return value, spendableUTXO
}
