package utils

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

//strconv转换方式
func IntToHex(num int64) []byte {
	numString := strconv.FormatInt(num, 10)
	numHex := []byte(numString)
	return numHex
}

////buff转换方式
//func IntToHexBuff(num int64) []byte {
//	return nil
//}

//时间戳转换成时间
func ConvertToTime(stamp int64) string {
	format := time.Unix(stamp, 0).Format("2006-01-02 15:04:05")
	return format
}

//json字符串转数组
func JsonToArray(jsonString string) []string {
	var sArr []string
	if err := json.Unmarshal([]byte(jsonString), &sArr); err != nil {
		log.Panic(err)
	}
	return sArr
}
