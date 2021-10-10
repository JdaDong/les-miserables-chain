package chain

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"les-miserables-chain/persistence"
	"log"
	"math/big"
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
			//创世区块集成交易
			coinbaseTx := NewCoinBaseTX("levy", "In a soldier's stance, I aimed my hand at the mongrel dogs who teach")
			genesisBlock := NewGenesisBlock(coinbaseTx)
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
func (chain *Chain) MineBlock(transactions []*Transaction) {
	err := chain.DB.Update(func(tx *bolt.Tx) error {
		//创建区块
		newBlock := NewBlock(transactions, chain.LastHash)
		//获取当前表
		b := tx.Bucket([]byte(persistence.BlockBucket))
		if b != nil {
			//存储区块数据
			err := b.Put(newBlock.BlockCurrentHash, Serialize(newBlock))
			if err != nil {
				log.Panic(err)
			}
			//存储最新出块的hash
			err = b.Put([]byte("last"), newBlock.BlockCurrentHash)
			if err != nil {
				log.Panic(err)
			}
			//更新最新出块的hash
			chain.LastHash = newBlock.BlockCurrentHash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

}

//查询地址下的未花费交易集合
func (chain *Chain) FindUnspentTransactions(address string) []Transaction {
	//存储未花费的交易
	var unspentTxs []Transaction
	//存储花费的交易
	spentTxs := make(map[string][]int)
	blockchainIterator := chain.Iterator()
	var hashInt big.Int

	for {
		err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
			//获取当前区块
			b := tx.Bucket([]byte(persistence.BlockBucket))
			blockBytes := b.Get(blockchainIterator.CurrentHash)
			block := DeserializeBlock(blockBytes)

			for _, transaction := range block.Transactions {
				//将交易ID转换为16进制
				index := hex.EncodeToString(transaction.Index)
				//Outputs的label
			Outputs:
				//遍历交易输出
				for outIdx, out := range transaction.Outputs {
					//判断是否已经被花费？
					if spentTxs[index] != nil {
						for _, spentOut := range spentTxs[index] {
							if spentOut == outIdx {
								continue Outputs
							}
						}
					}
					//如果是交易输出的解锁对象，则加入未花费交易
					if out.UnlockOutput(address) {
						unspentTxs = append(unspentTxs, *transaction)
					}
				}
				//判断是否是coinbase交易
				if transaction.IsCoinbase() == false {
					//遍历交易输入
					for _, in := range transaction.Inputs {
						//如果是交易输入解锁对象，则加入已花费交易中
						if in.UnlockInput(address) {
							//inTxID := hex.EncodeToString(in.TxID)
							spentTxs[index] = append(spentTxs[index], in.OutputIndex)
						}
					}
				}

			}
			fmt.Println()
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
		blockchainIterator = blockchainIterator.Next()
		hashInt.SetBytes(blockchainIterator.CurrentHash)
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
	return unspentTxs
}

//查询可用的未花费输出信息
func (chain *Chain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	//未花费交易下的未花费交易输出
	unspentOutputs := make(map[string][]int)
	//未花费交易
	unspentTxs := chain.FindUnspentTransactions(address)
	//未花费交易输出的value总量
	accumulated := 0
Work:
	//遍历未花费交易
	for _, tx := range unspentTxs {
		//获取未花费交易的交易ID
		txID := hex.EncodeToString(tx.Index)
		//遍历该未花费交易下的未花费交易输出
		for outIdx, out := range tx.Outputs {
			if out.UnlockOutput(address) && accumulated < amount {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)
				if accumulated >= amount {
					break Work
				}
			}
		}
	}
	return accumulated, unspentOutputs
}
