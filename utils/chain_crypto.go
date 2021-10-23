package utils

import (
	"bytes"
	"crypto/sha256"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

//获取SHA256字节数组
func GetSha256(data []byte) []byte {
	//等同于sha256.Sum256
	hasher := sha256.New()
	hasher.Write(data)
	res := hasher.Sum(nil)
	return res
}

//获取RIPEMD160字节数组
func GetRipemd160(data []byte) []byte {
	hasher := ripemd160.New()
	hasher.Write(data)
	res := hasher.Sum(nil)
	return res
}

var encodeStd = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) []byte {
	var result []byte
	x := big.NewInt(0).SetBytes(input)
	base := big.NewInt(int64(len(encodeStd)))
	zero := big.NewInt(0)
	mod := &big.Int{}
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, encodeStd[mod.Int64()])
	}
	ReverseBytes(result)
	for b := range input {
		if b == 0x00 {
			result = append([]byte{encodeStd[0]}, result...)
		} else {
			break
		}
	}
	return result
}
func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0
	for b := range input {
		if b == 0x00 {
			zeroBytes++
		}
	}
	payload := input[zeroBytes:]
	for _, b := range payload {
		charIndex := bytes.IndexByte(encodeStd, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))

	}
	decoded := result.Bytes()
	decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)
	return decoded
}
func ReverseBytes(data []byte) {
	left := 0
	right := len(data) - 1
	for left <= right {
		data[left], data[right] = data[right], data[left]
		left++
		right--
	}
}
