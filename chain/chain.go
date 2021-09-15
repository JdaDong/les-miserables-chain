package chain

import (
	"fmt"
	"github.com/boltdb/bolt"
	"les-miserables-chain/persistence"
	"log"
)

//链结构体
type Chain struct {
	LastHash []byte   //链的最新高度区块hash
	DB       *bolt.DB //数据库对象
}

//创世区块链
func NewBlockChain() *Chain {
	var lastHash []byte

	db, err := bolt.Open(persistence.DbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(persistence.BlockBucket))
		//判断bucket是否存在
		if b == nil {
			fmt.Println("Creating the genesis block.....")
			genesisBlock := NewGenesisBlock()
			//bucket不存在，创建一个桶
			b, err := tx.CreateBucket([]byte(persistence.BlockBucket))
			if err != nil {
				log.Panic(err)
			}
			//创世区块存储到bucket中
			err = b.Put(genesisBlock.BlockCurrentHash, Serialize(genesisBlock))
			if err != nil {
				log.Panic(err)
			}
			//存储最新的出块hash
			err = b.Put([]byte("last"), genesisBlock.BlockCurrentHash)
			if err != nil {
				log.Panic(err)
			}
			lastHash = genesisBlock.BlockCurrentHash
		} else {
			lastHash = b.Get([]byte("last"))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return &Chain{
		LastHash: lastHash,
		DB:       db,
	}
}

//区块派生
func (chain *Chain) AddBlock(data string) {

}
