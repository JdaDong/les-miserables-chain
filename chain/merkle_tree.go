package chain

import "crypto/sha256"

//默克尔树结点
type MerkleTreeNode struct {
	Left  *MerkleTreeNode
	Right *MerkleTreeNode
	Data  []byte
}

//默克尔树
type MerkleTree struct {
	RootNode *MerkleTreeNode
}

//新建默克尔树
func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleTreeNode
	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}
	for _, datum := range data {
		node := NewMerkleTreeNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}
	for i := 0; i < len(data)/2; i++ {
		var newLevel []MerkleTreeNode
		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleTreeNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}
		if len(newLevel)%2 != 0 {
			newLevel = append(newLevel, newLevel[len(newLevel)-1])
		}
		nodes = newLevel
	}
	tree := MerkleTree{&nodes[0]}
	return &tree
}

//新建默克尔树结点
func NewMerkleTreeNode(left, right *MerkleTreeNode, data []byte) *MerkleTreeNode {
	node := MerkleTreeNode{}
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]

	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		node.Data = hash[:]
	}
	node.Left = left
	node.Right = right
	return &node
}
