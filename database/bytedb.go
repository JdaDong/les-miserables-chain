package database

import "os"

//数据库文件
const DbFile = "database/blockchain.database"

//数据仓库
const BlockBucket = "blocks"

func DbExist() bool {
	_, err := os.Stat(DbFile)
	return err == nil || !os.IsNotExist(err)
}

func DeleteDbFile() error {
	return os.Remove(DbFile)
}
