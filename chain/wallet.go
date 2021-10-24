package chain

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"les-miserables-chain/utils"
	"log"
)

var addressHex = []byte{byte(0x00)}

const addressChecksumLen = 4

type Wallet struct {
	PrivateKey ecdsa.PrivateKey //私钥
	PublicKey  []byte           //公钥
}

//获取私钥-公钥对
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
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
	privateKey, publicKey := NewKeyPair()
	return &Wallet{privateKey, publicKey}
}

//获取地址
func (w *Wallet) GetAddress() []byte {
	//双哈希生成公钥hash
	sha256Hash := utils.GetSha256(w.PublicKey)      //公钥第一次sha256
	ripemd160Hash := utils.GetRipemd160(sha256Hash) //公钥第二次ripemd160

	//Base58Check编码
	dataWithVersion := append(addressHex, ripemd160Hash...) //给公钥hash加上版本前缀
	checkSum := CheckSum(dataWithVersion)                   //根据公钥hash和前缀生成4位校验码
	payload := append(dataWithVersion, checkSum...)
	address := utils.Base58Encode(payload)
	return address
}

//获取4位校验码
func CheckSum(data []byte) []byte {
	firstHash := utils.GetSha256(data)
	secondHash := utils.GetSha256(firstHash)
	return secondHash[:addressChecksumLen]
}

func CheckAddress(addr []byte) bool {
	addressDecoded := utils.Base58Decode(addr)
	fmt.Println(string(addressDecoded))

	checkSumBytes := addressDecoded[len(addressDecoded)-addressChecksumLen:]
	dataWithVersion := addressDecoded[:len(addressDecoded)-addressChecksumLen]
	checkBytes := CheckSum(dataWithVersion)
	if bytes.Compare(checkSumBytes, checkBytes) == 0 {
		return true
	}
	return false
}
