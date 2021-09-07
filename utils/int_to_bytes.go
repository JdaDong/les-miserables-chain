package utils

import "strconv"

//strconv转换方式
func IntToHex(num int64) []byte {
	numString := strconv.FormatInt(num, 10)
	numHex := []byte(numString)
	return numHex
}

//buff转换方式
func IntToHexBuff(num int64) []byte {
	return nil
}
