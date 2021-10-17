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
func (chain *Chain) UnUTXOs(address string, txs []*Transaction) []*UTXO {
	//UTXO未花费交易输出集合
	var unUTXOs []*UTXO
	//已花费交易输出
	spentTXOutputs := make(map[string][]int)

	//1.遍历同一区块的前几个交易输入
	for _, tx := range txs {
		if tx.IsCoinbase() == false {
			for _, in := range tx.Inputs {
				if in.UnlockInput(address) {
					key := hex.EncodeToString(in.TxID)
					spentTXOutputs[key] = append(spentTXOutputs[key], in.OutputIndex)
				}
			}
		}
	}

	//2.遍历同一区块的前几个交易输出
	for _, tx := range txs {
	Work1:
		for index, out := range tx.Outputs {
			if out.UnlockOutput(address) {
				if len(spentTXOutputs) == 0 {
					utxo := &UTXO{tx.Index, index, out}
					unUTXOs = append(unUTXOs, utxo)
				} else {

					for hash, indexArray := range spentTXOutputs {
						txHashStr := hex.EncodeToString(tx.Index)
						if hash == txHashStr {
							var isUnSpentUTXO bool
							for _, outIndex := range indexArray {
								if index == outIndex {
									isUnSpentUTXO = true
									continue Work1
								}
								if isUnSpentUTXO == false {
									utxo := &UTXO{tx.Index, index, out}
									unUTXOs = append(unUTXOs, utxo)
								}
							}
						} else {
							utxo := &UTXO{tx.Index, index, out}
							unUTXOs = append(unUTXOs, utxo)
						}
					}
				}
			}
		}
	}

	//区块链迭代器 遍历前面所有区块的交易
	blockIterator := chain.Iterator()

	for {

		block := blockIterator.NextBlock()

		//从区块最后一笔交易遍历
		for i := len(block.Transactions) - 1; i >= 0; i-- {
			tx := block.Transactions[i]
			if tx.IsCoinbase() == false {
				for _, in := range tx.Inputs {
					if in.UnlockInput(address) {

						key := hex.EncodeToString(in.TxID)

						spentTXOutputs[key] = append(spentTXOutputs[key], in.OutputIndex)
					}

				}
			}

		work:
			for index, out := range tx.Outputs {

				if out.UnlockOutput(address) {

					if len(spentTXOutputs) != 0 {

						var isSpentUTXO bool

						for txHash, indexArray := range spentTXOutputs {

							for _, i := range indexArray {
								if index == i && txHash == hex.EncodeToString(tx.Index) {
									isSpentUTXO = true
									continue work
								}
							}
						}

						if isSpentUTXO == false {

							utxo := &UTXO{tx.Index, index, out}
							unUTXOs = append(unUTXOs, utxo)

						}
					} else {
						utxo := &UTXO{tx.Index, index, out}
						unUTXOs = append(unUTXOs, utxo)
					}

				}

			}

		}

		//fmt.Println(unUTXOs)

		var hashInt big.Int
		hashInt.SetBytes(block.BlockPreHash)
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}

	}
	//返回所有交易输出集合
	return unUTXOs

}

//多方转账-可用UTXO交易输出
func (chain *Chain) SpendableUTXOs(from string, amount int, txs []*Transaction) (int, map[string][]int) {
	//1.获取转账源地址的未花费交易输出
	utxos := chain.UnUTXOs(from, txs)

	//可用UTXO交易输出
	spendableUTXO := make(map[string][]int)

	//余额
	var value int

	//2.遍历未花费UTXO交易输出，一旦余额充足就不再遍历
	for _, utxo := range utxos {
		value += utxo.OutPut.Value
		hash := hex.EncodeToString(utxo.TxHash)
		spendableUTXO[hash] = append(spendableUTXO[hash], utxo.Index)
		if value >= amount {
			break
		}
	}
	//3.遍历完后，余额还是不足，退出
	if value < amount {
		fmt.Printf("%s地址中的余额不足\n", from)
		os.Exit(1)
	}
	return value, spendableUTXO
}
