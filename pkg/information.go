package pkg

const Length = 108

type Information struct {
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
func HandleData(data string) Information {
	info := &Information{}
	info.Header = data[0:4]
	info.Tag = data[4:8]
	info.Len = data[8:12]
	info.SendNode = data[12:16]
	info.ReceiveNode = data[16:20]
	info.Time = data[20:52]
	info.Order = data[52:56]
	info.Batch = data[56:64]
	info.Distance = data[64:72]
	info.Orientation = data[72:76]
	info.Course = data[76:80]
	info.Speed = data[80:84]
	info.Longitude = data[84:92]
	info.Latitude = data[92:100]
	info.Method = data[100:102]
	info.Status = data[102:104]
	info.Tail = data[104:]
	return *info
}
