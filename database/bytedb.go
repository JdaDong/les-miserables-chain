package database

import (
	"fmt"
	"os"
)

//数据库文件
var DbFile string

//数据仓库
var BlockBucket string

//UTXO数据桶
var UTXOBucket string

func GenerateDatabase(nodeID string) {
	DbFile = fmt.Sprintf("database/NODE-%s.database", nodeID)
	BlockBucket = "blocks"
	UTXOBucket = "utxo"
}

func DbExist() bool {
	_, err := os.Stat(DbFile)
	return err == nil || !os.IsNotExist(err)
}

func DeleteDbFile() error {
	return os.Remove(DbFile)
}
