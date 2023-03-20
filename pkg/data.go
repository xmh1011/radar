package pkg

import (
	"encoding/hex"
	"math"
	"strconv"
	"time"
)

const Length = 108

type CnosData struct {
	Measurement string
	Header      string // 报头
	Tag         string // 标识
	Len         string // 长度
	SendNode    string // 发送节点
	ReceiveNode string // 接收节点
	Time        int64  // 时间
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
	cnos.Len = data[10:12] + data[8:10]
	cnos.SendNode = data[14:16] + data[12:14]
	cnos.ReceiveNode = data[18:20] + data[16:18]
	// TODO: 处理时间
	cnos.Time = GetTime(data[20:52])
	cnos.Time = ParseTime(data[20:52])
	cnos.Order = data[54:56] + data[52:54]
	cnos.Batch = data[62:64] + data[60:62] + data[58:60] + data[56:58]
	cnos.Distance = data[70:72] + data[68:70] + data[66:68] + data[64:66]
	cnos.Orientation = data[74:76] + data[72:74]
	cnos.Course = data[78:80] + data[76:78]
	cnos.Speed = data[82:84] + data[80:82]
	cnos.Longitude = data[90:92] + data[88:90] + data[86:88] + data[84:86]
	cnos.Latitude = data[98:100] + data[96:98] + data[94:96] + data[92:94]
	cnos.Method = data[100:102]
	cnos.Status = data[102:104]
	cnos.Tail = data[104:]
	return *cnos
}

func GetTime(s string) int64 {
	return ParseTime(s)
}

// HexToDec64 将32位的十六进制数据转化为十进制
func HexToDec64(hex string) int64 {
	var dec int64
	for i := 0; i < len(hex); i++ {
		dec += int64(HexToDec(hex[i:i+1])) * int64(math.Pow(16, float64(len(hex)-i-1)))
	}
	return dec
}

func ParseTime(s string) int64 {
	
	data, err := hex.DecodeString(s)
	
	if err != nil {
		panic(err)
	}
	
	// 将字节数组转换为SYSTEMTIME结构体
	if len(data) != 16 {
		panic("invalid SYSTEMTIME data")
	}
	st := SystemTime{
		wYear:         uint16(data[0]) | uint16(data[1])<<8,
		wMonth:        uint16(data[2]) | uint16(data[3])<<8,
		wDayOfWeek:    uint16(data[4]) | uint16(data[5])<<8,
		wDay:          uint16(data[6]) | uint16(data[7])<<8,
		wHour:         uint16(data[8]) | uint16(data[9])<<8,
		wMinute:       uint16(data[10]) | uint16(data[11])<<8,
		wSecond:       uint16(data[12]) | uint16(data[13])<<8,
		wMilliseconds: uint16(data[14]) | uint16(data[15])<<8,
	}
	
	// 将SYSTEMTIME结构体转换为time.Time类型
	utc := time.Date(int(st.wYear), time.Month(st.wMonth), int(st.wDay), int(st.wHour), int(st.wMinute), int(st.wSecond), int(st.wMilliseconds)*1000000, time.UTC)
	
	// 将utc转化为纳秒级时间戳输出
	return utc.UnixNano()
	
}

type SystemTime struct {
	wYear         uint16
	wMonth        uint16
	wDayOfWeek    uint16
	wDay          uint16
	wHour         uint16
	wMinute       uint16
	wSecond       uint16
	wMilliseconds uint16
}

func HexToDec(hex string) int {
	dec, _ := strconv.ParseInt(hex, 16, 32)
	return int(dec)
}
