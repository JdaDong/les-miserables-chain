package chain

type Version struct {
	Version    int    // 消息版本
	BestHeight int    // 当前节点区块的高度
	AddrFrom   string //当前节点的地址
}
