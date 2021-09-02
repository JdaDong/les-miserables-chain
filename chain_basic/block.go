package chain_basic

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//区块信息
type Block struct {
	BlockTimestamp   int64
	BlockPreHash     []byte
	BlockData        []byte
	BlockCurrentHash []byte
	BlockNonce       int64
}

//区块hash计算
func (block *Block) SetHash() {
	timeString := strconv.FormatInt(block.BlockTimestamp, 10)
	timestamp := []byte(timeString)
	headers := bytes.Join([][]byte{block.BlockPreHash, block.BlockData, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	block.BlockCurrentHash = hash[:]
}

//生成区块
func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{
		BlockTimestamp:   time.Now().UnixNano() / 1e6, //精确到毫秒
		BlockPreHash:     preBlockHash,
		BlockData:        []byte(data),
		BlockCurrentHash: []byte{},
		BlockNonce:       0,
	}
	block.SetHash()
	return block
}

//创世区块
func NewGenesisBlock() *Block {
	preBlockHash := make([]byte, 32)
	//创世区块的父区块0x0
	for i := 0; i < 32; i++ {
		preBlockHash[i] = 0
	}
	return NewBlock("Hello Genesis.", preBlockHash)
}
