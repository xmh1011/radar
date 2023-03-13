package pkg

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func ReceiveData(UdpAddr string, UdpPort int) (data []string, err error) {
	// fmt.Println(3)
	// 创建UDP监听地址
	addr, err := net.ResolveUDPAddr("udp", UdpAddr+":"+strconv.Itoa(UdpPort))
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}
	
	// 创建UDP监听连接
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("Error Listening: %v", err)
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn close failed, err:", err)
		}
	}(conn)
	
	// 接收数据并处理
	for {
		var messages []string
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("conn readFromUDP failed, err:", err)
			break
		}
		// fmt.Printf("data:%v, addr:%v, count:%v\n", string(buffer[:n]), addr, n)
		// 去掉数据中的空格
		msg := strings.Replace(string(buffer[0:n]), " ", "", -1)
		// 将处理后的数据存储到字符串数组中
		messages = append(messages, msg)
		// 将messages中的数据每108位分割一次，存储到data中
		// 将result中的空格删除，变成新的字符串
		for i := 0; i < len(messages); i += Length {
			data = append(data, strings.Join(messages[i:i+Length], ""))
			// data = append(data, messages[i:i+Length]...)
		}
		return data, nil
	}
	return nil, fmt.Errorf("Error: %v", "ReceiveData failed")
}
