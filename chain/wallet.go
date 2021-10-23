package chain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey //私钥
	PublicKey  []byte           //公钥
}

//获取私钥-公钥对
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()                              //实现了secp256r1的椭圆曲线
	private, err := ecdsa.GenerateKey(curve, rand.Reader) //根据曲线生成私钥
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...) //合并为无前缀公钥
	return *private, pubKey
}

//初始化钱包
func NewWallet() *Wallet {
	privateKey, publicKey := newKeyPair()
	return &Wallet{privateKey, publicKey}
}
