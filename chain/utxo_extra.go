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
