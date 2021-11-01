# 公链的设计与实现

## 区块block

区块结构

- 时间戳 Timestamp
- 父区块哈希 PreBlockHash
- 当前区块哈希 Hash
- 交易数据 Data
- nonce值

### 创建区块

function NewBlock(data string,preBlockHash []byte)*Block

### 生成hash值

function (block *Block) SetHash()

1. 时间戳转换为字节数组
2. 将区块其他属性进行SHA256函数计算
3. 生成当前区块的hash值






## 交易签名

### 转账交易

```go
func CreateTransaction(from, to string, amount int, chain *Chain, txs []*Transaction) *Transaction 
```

