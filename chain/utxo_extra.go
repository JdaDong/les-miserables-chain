package chain

import (
	"encoding/hex"
	"les-miserables-chain/database"
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
