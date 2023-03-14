package udp

import (
	"net"
	"strconv"
	"strings"
)

func ReceiveData(UdpAddr string, UdpPort int) (data []string, err error) {
	
	// 创建UDP监听地址
	addr, err := net.ResolveUDPAddr("udp", UdpAddr+":"+strconv.Itoa(UdpPort))
	if err != nil {
		return nil, err
	}
	
	// 创建UDP监听连接
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}
	
	// 接收数据并处理
	for {
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		err = conn.Close()
		if err != nil {
			return nil, err
		}
		
		// 去掉数据中的空格
		msg := strings.Replace(string(buffer[0:n]), " ", "", -1)
		// 将处理后的数据存储到字符串数组中
		data = append(data, msg)
		
		return data, err
	}
	
}
