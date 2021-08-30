package chain_basic

type Chain struct {
	Blocks []*Block
}

//创世区块链
func NewBlockChain() *Chain {
	return &Chain{[]*Block{NewGenesisBlock()}}
}

//区块派生
func (chain *Chain) AddBlock(data string) {
	//1.创建新的区块
	chainLength := len(chain.Blocks) - 1                                   //计算链长度
	newBlock := NewBlock(data, chain.Blocks[chainLength].BlockCurrentHash) //生成区块
	//2.区块链派生
	chain.Blocks = append(chain.Blocks, newBlock)
}
