package chain

import (
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
	var lastHash [32]byte

	db, err := bolt.Open(persistence.DbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
}

//区块派生
func (chain *Chain) AddBlock(data string) {
	//1.创建新的区块
	chainLength := len(chain.Blocks) - 1                                   //计算链长度
	newBlock := NewBlock(data, chain.Blocks[chainLength].BlockCurrentHash) //生成区块
	//2.区块链派生
	chain.Blocks = append(chain.Blocks, newBlock)
}
