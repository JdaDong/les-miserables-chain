package chain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"les-miserables-chain/utils"
	"math"
	"math/big"
)

const (
	targetOffset = 30            //难度偏移量
	maxNonce     = math.MaxInt64 //最大nonce值
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

//准备数据
func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join([][]byte{
		pow.block.BlockPreHash,
		pow.block.BlockData,
		utils.IntToHex(pow.block.BlockTimestamp),
		utils.IntToHex(int64(targetOffset)),
		utils.IntToHex(nonce),
	},
		[]byte{},
	)
	return data
}

//新增权益证明对象
func NewProof(block *Block) *ProofOfWork {
	//难度系数
	target := big.NewInt(7)
	//位运算左移计算
	target.Lsh(target, uint(256-targetOffset))
	pow := &ProofOfWork{block, target}
	return pow
}

//pow计算
func (pow *ProofOfWork) ProofWork() (int64, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := int64(0)

	//最大maxNonce范围内循环
	for nonce < maxNonce {
		data := pow.prepareData(nonce) //预处理数据
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:]) //预处理数据哈希值

		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println("\n\n")
	return nonce, hash[:]
}

//验证nonce值
func (pow *ProofOfWork) Validate() bool {

	var hashInt big.Int

	data := pow.prepareData(pow.block.BlockNonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
