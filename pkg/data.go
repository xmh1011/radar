package pkg

import (
	"strconv"
)

// 将十六进制数据转化为十进制
func HexToDec(hex string) int {
	dec, _ := strconv.ParseInt(hex, 16, 32)
	return int(dec)
}

// 将十六进制数据转化为十进制 int64类型
func HexToDec64(hex string) int64 {
	dec, _ := strconv.ParseInt(hex, 16, 64)
	return dec
}