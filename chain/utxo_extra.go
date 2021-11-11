package chain

import (
	"bytes"
	"encoding/hex"
	"les-miserables-chain/database"
	"les-miserables-chain/utils"
	"log"

	"github.com/boltdb/bolt"
)

type UTXORecord struct {
	Blockchain *Chain
}

//重置数据桶
func (utxoRecord *UTXORecord) ResetUTXORecord() {
	err := utxoRecord.Blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.UTXOBucket))
		if b != nil {
			err := tx.DeleteBucket([]byte(database.UTXOBucket))
			if err != nil {
				log.Panic(err)
			}
		}
		b, _ = tx.CreateBucket([]byte(database.UTXOBucket))
		if b != nil {
			txOutputsMap := utxoRecord.Blockchain.FindUTXOMap()
			for hash, outs := range txOutputsMap {
				txHash, _ := hex.DecodeString(hash)
				_ = b.Put(txHash, outs.Serialize())
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

//获取地址余额(大量数据)
func (utxoRecord *UTXORecord) GetBalance(address string) int {
	UTXOS := utxoRecord.findUTXO(address)
	var amount int
	for _, utxo := range UTXOS {
		amount += utxo.OutPut.Value
	}
	return amount
}

func (utxoRecord *UTXORecord) findUTXO(address string) []*UTXO {
	var utxos []*UTXO
	err := utxoRecord.Blockchain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.UTXOBucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			txOutputs := DeserializeTXOutputs(v)
			for _, utxo := range txOutputs.UTXOS {
				if utxo.OutPut.UnLockScriptPubKeyWithAddress(address) {
					utxos = append(utxos, utxo)
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return utxos
}
func (utxoRecord *UTXORecord) FindUnPackageSpendableUTXOs(from string, txs []*Transaction) []*UTXO {
	var unUTXOs []*UTXO
	spentTXOutputs := make(map[string][]int)

	for _, tx := range txs {
		if tx.IsCoinbase() == false {
			for _, in := range tx.TxInputs {
				publicKeyHash := utils.Base58Decode([]byte(from))
				publicKeyRipemd160Hash := publicKeyHash[1 : len(publicKeyHash)-4]
				if in.UnlockPublicKeyHash(publicKeyRipemd160Hash) {
					txHashHex := hex.EncodeToString(in.TxID)
					spentTXOutputs[txHashHex] = append(spentTXOutputs[txHashHex], in.OutputIndex)
				}
			}
		}
	}

	//遍历交易
	for _, tx := range txs {
	Work:
		//1.遍历交易输出
		for index, out := range tx.TxOutputs {
			if out.UnLockScriptPubKeyWithAddress(from) {
				//2.如果花费的交易输出为空,则直接构建一个utxo
				if len(spentTXOutputs) == 0 {
					utxo := &UTXO{
						TxHash: tx.TxHash,
						Index:  index,
						OutPut: out,
					}
					unUTXOs = append(unUTXOs, utxo)
				} else {
					for hash, outPutsArray := range spentTXOutputs {
						txHashHex := hex.EncodeToString(tx.TxHash)
						if hash == txHashHex {
							var isSpentUTXO bool
							for _, outIndex := range outPutsArray {
								if index == outIndex {
									isSpentUTXO = true
									continue Work
								}
								if isSpentUTXO == false {
									utxo := &UTXO{
										TxHash: tx.TxHash,
										Index:  index,
										OutPut: out,
									}
									unUTXOs = append(unUTXOs, utxo)
								}
							}
						} else {
							utxo := &UTXO{
								TxHash: tx.TxHash,
								Index:  index,
								OutPut: out,
							}
							unUTXOs = append(unUTXOs, utxo)
						}
					}
				}
			}
		}
	}
	return unUTXOs
}

//在未打包交易中，查找有限的可花费uxto
func (utxoRecord *UTXORecord) FindSpendableUTXOs(from string, amount int, txs []*Transaction) (int, map[string][]int) {
	//查找未打包交易中的可花费utxo
	unPackageUTXOS := utxoRecord.FindUnPackageSpendableUTXOs(from, txs)

	spentableUTXO := make(map[string][]int)
	var money int = 0
	//遍历所有可花费utxo，如果钱足够，则返回有限的可花费uxto信息
	for _, UTXO := range unPackageUTXOS {
		money += UTXO.OutPut.Value
		txHashHex := hex.EncodeToString(UTXO.TxHash)
		spentableUTXO[txHashHex] = append(spentableUTXO[txHashHex], UTXO.Index)
		if money >= amount {
			return money, spentableUTXO
		}
	}
	//如果钱不够
	utxoRecord.Blockchain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.UTXOBucket))
		if b != nil {
			c := b.Cursor()
		UTXOBREAK:
			for k, v := c.First(); k != nil; k, v = c.Next() {
				txOutputs := DeserializeTXOutputs(v)
				for _, UTXO := range txOutputs.UTXOS {
					money += UTXO.OutPut.Value
					txHashHex := hex.EncodeToString(UTXO.TxHash)
					spentableUTXO[txHashHex] = append(spentableUTXO[txHashHex], UTXO.Index)
					if money >= amount {
						break UTXOBREAK
					}
				}
			}
		}
		return nil
	})
	if money < amount {
		log.Panic("交易失败，余额不足!")
	}
	return money, spentableUTXO
}

func (utxoRecord *UTXORecord) Update() {
	block := utxoRecord.Blockchain.Iterator().NextBlock()
	var ins []*TXInput
	outsMap := make(map[string]*TXOutputs)
	//查找需要删除的交易输入数据
	for _, tx := range block.Transactions {
		for _, in := range tx.TxInputs {
			ins = append(ins, in)
		}
	}

	//遍历最新区块
	for _, tx := range block.Transactions {
		utxos := []*UTXO{}
		//获取未花费utxo
		for index, out := range tx.TxOutputs {
			isSpent := false
			for _, in := range ins {
				if in.OutputIndex == index && bytes.Compare(tx.TxHash, in.TxID) == 0 {
					isSpent = true
					continue
				}
			}
			if isSpent == false {
				utxo := &UTXO{tx.TxHash, index, out}
				utxos = append(utxos, utxo)
			}
		}
		if len(utxos) > 0 {
			txHashHex := hex.EncodeToString(tx.TxHash)
			outsMap[txHashHex] = &TXOutputs{utxos}
		}
	}
	//更新数据桶
	err := utxoRecord.Blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.UTXOBucket))
		if b != nil {
			//遍历交易输入
			for _, in := range ins {
				txOutputsBytes := b.Get(in.TxID)
				if len(txOutputsBytes) == 0 {
					continue
				}
				txOutputs := DeserializeTXOutputs(txOutputsBytes)
				UTXOS := []*UTXO{}
				toDelete := false
				for _, utxo := range txOutputs.UTXOS {
					if in.OutputIndex == utxo.Index && bytes.Compare(utxo.OutPut.ScriptPubKey, utils.GetRipemd160(in.PublicKey)) == 0 {
						toDelete = true
					} else {
						UTXOS = append(UTXOS, utxo)
					}
				}
				if toDelete {
					_ = b.Delete(in.TxID)
					if len(UTXOS) > 0 {
						prevTXOutputs := outsMap[hex.EncodeToString(in.TxID)]
						prevTXOutputs.UTXOS = append(prevTXOutputs.UTXOS, UTXOS...)
						outsMap[hex.EncodeToString(in.TxID)] = prevTXOutputs
					}
				}
			}
			for keyHash, outPuts := range outsMap {
				keyHashBytes, _ := hex.DecodeString(keyHash)
				_ = b.Put(keyHashBytes, outPuts.Serialize())
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}
