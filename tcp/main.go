package main

import (
	"fmt"
	"net"
	"radar/pkg"
	"strconv"
	"strings"
	"time"
)

var TcpAddr = "localhost"
var TcpPort = 8000

// 测试代码，模拟通过tcp协议不断发送数据
func main() {
	
	// 定义发送数据
	var filePath = "../radar/"
	var fileName = "radar.txt"
	
	file := filePath + fileName
	data, err := pkg.ReadDataFromFile(file)
	if err != nil {
		fmt.Printf("Read data from file %s failed: %s\n", file, err)
		return
	}
	
	for {
		conn, err := net.Dial("tcp", TcpAddr+":"+strconv.Itoa(TcpPort))
		if err != nil {
			fmt.Println("Error connecting:", err)
			time.Sleep(time.Second) // 等待1秒后重试连接
			continue
		}
		
		fmt.Println("Connected to server:", conn.RemoteAddr())
		
		// 将[]string转换为[]byte
		temp := strings.Join(data, "\n")
		for {
			_, err = conn.Write([]byte(temp)) // 发送数据
			fmt.Println(temp)
			fmt.Println("发送数据成功")
			if err != nil {
				fmt.Println("Error sending data:", err)
				break
			}
			
			time.Sleep(time.Second) // 每秒发送一次数据
		}
		
		err = conn.Close()
		if err != nil {
			return
		} // 关闭连接
	}
}
