package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGetSha256(t *testing.T) {
	data := []byte("hello")
	fmt.Println(string(data))
	result := GetSha256(data)
	fmt.Println(len(result))
	fmt.Println(hex.EncodeToString(result))
	fmt.Println(sha256.Sum256(data))
}
func TestGetRipemd160(t *testing.T) {
	data := []byte("hello")
	fmt.Println(string(data))
	result := GetRipemd160(data)
	fmt.Println(result)
	fmt.Println(hex.EncodeToString(result))
}

func TestBase58Encode(t *testing.T) {
	data := []byte("hello")
	fmt.Println(string(data))
	result := Base58Encode(data)
	fmt.Println(len(result))
	fmt.Println(hex.EncodeToString(result))
}

func TestBase58Decode(t *testing.T) {
	data := []byte("hello")
	fmt.Println(string(data))
	result := Base58Encode(data)
	fmt.Println(len(result))
	fmt.Println(hex.EncodeToString(result))
	result = Base58Decode(result)
	fmt.Println(string(result))
}

func TestReverseBytes(t *testing.T) {
	data := []byte("hello")
	fmt.Println(data)
	ReverseBytes(data)
	fmt.Println(data)
}
