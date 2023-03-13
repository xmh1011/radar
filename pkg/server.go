package pkg

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
		fmt.Println("server循环")
		// conn 是一个 net.Conn 类型的变量，用于表示客户端的连接
		// conn, err := listener.Accept() // 接收客户端的连接请求
		// if err != nil {
		// 	fmt.Println("server循环失败")
		// 	// 如果接收客户端的连接请求失败，则打印错误信息，然后继续接收下一个客户端的连接请求
		// 	fmt.Println(err)
		// }
		// // 将[]string转换为string
		//
		// temp := strings.Join(data, "\n")
		// fmt.Println("temp:", temp)
		// go func() {
		// 	_, err := io.WriteString(conn, temp) // 向客户端发送当前时间
		// 	if err != nil {
		// 		return
		// 	}
		// }() // 并发处理客户端的连接请求
		// // 在前面添加一个go关键字，使函数在自己的goroutine中执行，而不是在主goroutine中执行
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

// // handleConn 函数接收一个 net.Conn 类型的参数, 然后向客户端发送当前时间
// func handleConn(c net.Conn) {
// 	defer c.Close()
// 	for {
// 		_, err := io.WriteString(c, time.Now().Format("15:04:05\n")) // 向客户端发送当前时间
// 		if err != nil {
// 			return
// 		}
// 		time.Sleep(1 * time.Second) // 每隔一秒，向客户端发送当前时间
// 	}
// }
