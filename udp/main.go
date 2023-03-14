package main

import (
	"fmt"
	"net"
	"radar/pkg"
	"strconv"
	"strings"
	"time"
)

var UdpAddr = "localhost"
var UdpPort = 8000

func main() {
	// 设置UDP地址
	serverAddr, err := net.ResolveUDPAddr("udp", UdpAddr+":"+strconv.Itoa(UdpPort))
	if err != nil {
		fmt.Println("ResolveUDPAddr failed:", err)
		return
	}
	
	// 定义发送数据
	var filePath = "../radar/"
	var fileName = "radar.txt"
	
	file := filePath + fileName
	data, err := pkg.ReadDataFromFile(file)
	if err != nil {
		fmt.Printf("Read data from file %s failed: %s\n", file, err)
		return
	}
	// 将[]string转换为[]byte
	temp := strings.Join(data, "\n")
	
	// 持续发送数据
	for {
		// 创建UDP连接
		conn, err := net.DialUDP("udp", nil, serverAddr)
		if err != nil {
			fmt.Println("DialUDP failed:", err)
			return
		}
		defer func(conn *net.UDPConn) {
			err := conn.Close()
			if err != nil {
				fmt.Println("Close failed:", err)
				return
			}
		}(conn)
		_, err = conn.Write([]byte(temp))
		if err != nil {
			fmt.Println("Write failed:", err)
			return
		}
		fmt.Println("Send data successfully!")
		
		time.Sleep(1 * time.Second) // 暂停1秒钟再发送下一次数据
	}
}
