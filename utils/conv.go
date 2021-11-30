package utils

import (
	"bytes"
	"encoding/gob"
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

//消息命令转字节数组
func CommandTobytes(command string) []byte {
	var bytes [12]byte //命令长度硬编码为12
	for i, cmd := range command {
		bytes[i] = byte(cmd)
	}
	return bytes[:]
}

//结构体序列化
func GobEncode(data interface{}) []byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(data)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
