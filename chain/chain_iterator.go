package chain

import (
	"github.com/boltdb/bolt"
	"les-miserables-chain/persistence"
	"log"
)

//链迭代器
type ChainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

//迭代器生成
func (chain *Chain) Iterator() *ChainIterator {
	return &ChainIterator{
		CurrentHash: chain.LastHash,
		DB:          chain.DB,
	}

}

//遍历迭代器
func (ci *ChainIterator) Next() *ChainIterator {
	var nextHash []byte
	err := ci.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(persistence.BlockBucket))
		currentBlockBytes := b.Get(ci.CurrentHash)

		currentBlock := DeserializeBlock(currentBlockBytes)
		nextHash = currentBlock.BlockPreHash
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return &ChainIterator{
		CurrentHash: nextHash,
		DB:          ci.DB,
	}
}
