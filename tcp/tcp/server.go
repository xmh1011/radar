package tcp

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

// 向固定端口发送数据
func Server(TcpAddr string, TcpPort int, data []string) error {
	for {
		fmt.Println("Server")
		_, err := net.Listen("tcp", TcpAddr+":"+strconv.Itoa(TcpPort))
		fmt.Println(TcpAddr + ":" + strconv.Itoa(TcpPort)) // 创建一个监听器，监听端口 8000
		if err != nil {
			log.Fatal(err)
		}
		
		conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{IP: net.ParseIP(TcpAddr), Port: TcpPort})
		fmt.Println(TcpAddr + ":" + strconv.Itoa(TcpPort))
		if err != nil {
			// fmt.Println("Error connecting:", err)
			return fmt.Errorf("Error connecting: %v", err)
		}
		defer conn.Close()
		
		// 将[]string转换为[]byte
		temp := strings.Join(data, "\n")
		
		// 发送数据
		_, err = conn.Write([]byte(temp))
		if err != nil {
			// fmt.Println("Error sending data:", err)
			return fmt.Errorf("Error sending data: %v", err)
		}
		fmt.Println("Data sent successfully.")
		return nil
	}
}
