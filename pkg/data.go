package pkg

import (
	"strconv"
)

type CnosData struct {
	Measurement string
	Header      string // 报头
	Tag         string // 标识
	Len         string // 长度
	SendNode    string // 发送节点
	ReceiveNode string // 接收节点
	Time        string // 时间
	Order       string // 顺序
	Batch       string // 批次
	Distance    string // 距离
	Orientation string // 方向
	Course      string // 航向
	Speed       string // 速度
	Longitude   string // 经度
	Latitude    string // 纬度
	Method      string // 方法
	Status      string // 状态
	Tail        string // 报尾   // 数据长度
}

// 处理数据函数，将数据分割成各个部分
func HandleData(data string) CnosData {
	cnos := &CnosData{}
	cnos.Header = data[0:4]
	cnos.Tag = data[4:8]
	cnos.Len = data[8:12]
	cnos.SendNode = data[12:16]
	cnos.ReceiveNode = data[16:20]
	cnos.Time = data[20:52]
	cnos.Order = data[52:56]
	cnos.Batch = data[56:64]
	cnos.Distance = data[64:72]
	cnos.Orientation = data[72:76]
	cnos.Course = data[76:80]
	cnos.Speed = data[80:84]
	cnos.Longitude = data[84:92]
	cnos.Latitude = data[92:100]
	cnos.Method = data[100:102]
	cnos.Status = data[102:104]
	cnos.Tail = data[104:]
	return *cnos
}

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
