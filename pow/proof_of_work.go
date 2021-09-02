package pow

import (
	"les-miserables-chain/chain_basic"
	"math/big"
)

//难度偏移量
const TargetOffset = 20

type ProofOfWork struct {
	block  *chain_basic.Block
	target *big.Int
}

//新增权益证明
func NewProofOfWork(block *chain_basic.Block) *ProofOfWork {
	//难度系数
	target := big.NewInt(7)
	//位运算左移计算
	target.Lsh(target, uint(256-TargetOffset))
	pow := &ProofOfWork{block, target}
	return pow
}

func (pow *ProofOfWork) Run() (int64, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

}
