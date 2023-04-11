package tcp

import (
	"fmt"
	"net"
	"strings"
)

// 测试代码，做为测试能否通过tcp协议接收数据

// 从固定端口通过TCP协议接收数据
func ReceiveDataTCP(TcpAddr string, TcpPort int) (data []string, err error) {
	fmt.Println("ReceiveDataTCP")
	
	// 连接TCP服务器
	// _, err = net.ResolveTCPAddr("tcp", TcpAddr+":"+strconv.Itoa(TcpPort))
	// _, err = net.Dial("tcp", TcpAddr+":"+strconv.Itoa(TcpPort))
	listener, err := net.Listen("tcp", "localhost:8000")
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	// 3. 关闭监听通道
	defer listener.Close()
	fmt.Println("server is Listening")
	for {
		// 2. 进行通道监听
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// 启动一个协程去单独处理该连接
		data, err = handle(conn, data)
		fmt.Println("data:", data)
		return data, err
		// go func(conn net.Conn) {
		// 	defer conn.Close()
		// 	buffer := make([]byte, 124)
		// 	fmt.Println("创建buffer成功")
		// 	n, err := conn.Read(buffer)
		// 	conn.Close()
		// 	if err != nil {
		// 		fmt.Println("读取数据失败：", err)
		// 	}
		// 	// 处理数据
		// 	fmt.Println("buffer:", string(buffer[:n]))
		//
		// 	var messages []string
		// 	// 去掉数据中的空格
		// 	msg := strings.Replace(string(buffer[0:n]), " ", "", -1)
		// 	// 将处理后的数据存储到字符串数组中
		// 	messages = append(messages, msg)
		// 	fmt.Println("messages:", messages)
		// 	// 将messages中的数据每108位分割一次，存储到data中
		// 	// 将result中的空格删除，变成新的字符串
		// 	for i := 0; i < len(messages); i += Length {
		// 		data = append(data, strings.Join(messages[i:i+Length], ""))
		// 		// data = append(data, messages[i:i+Length]...)
		// 	}
		// }(conn)
		// return data, nil
	}
	
	// // defer listener.Close()
	// fmt.Println("1")
	// // fmt.Println("2")
	// // if err != nil {
	// // 	return nil, fmt.Errorf("Error: %v", err)
	// // }
	// // defer conn.Close()
	//
	// for {
	// 	// 读取服务器发送的数据
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		fmt.Errorf(err.Error())
	// 	}
	// 	go handle(conn, data)
	// 	// 	go func(conn net.Conn) {
	// 	// 		// conn, err := listener.Accept() // 接收客户端的连接请求
	// 	// 		fmt.Println("for循环")
	// 	// 		buffer := make([]byte, 124)
	// 	// 		fmt.Println("创建buffer成功")
	// 	// 		n, err := conn.Read(buffer)
	// 	// 		conn.Close()
	// 	// 		if err != nil {
	// 	// 			fmt.Println("读取数据失败：", err)
	// 	// 		}
	// 	// 		// 处理数据
	// 	// 		fmt.Println("buffer:", string(buffer[:n]))
	// 	//
	// 	// 		var messages []string
	// 	// 		// 去掉数据中的空格
	// 	// 		msg := strings.Replace(string(buffer[0:n]), " ", "", -1)
	// 	// 		// 将处理后的数据存储到字符串数组中
	// 	// 		messages = append(messages, msg)
	// 	// 		fmt.Println("messages:", messages)
	// 	// 		// 将messages中的数据每108位分割一次，存储到data中
	// 	// 		// 将result中的空格删除，变成新的字符串
	// 	// 		for i := 0; i < len(messages); i += Length {
	// 	// 			data = append(data, strings.Join(messages[i:i+Length], ""))
	// 	// 			// data = append(data, messages[i:i+Length]...)
	// 	// 		}
	// 	// 		// return nil, fmt.Errorf("Error: %v", "ReceiveData failed")
	// 	// 	}
	// 	// }(conn)
	// 	// return data, nil
	// 	// 创建TCP监听连接
	// 	// _, err = net.ListenTCP("tcp", addr)
	// 	// fmt.Println("创建TCP监听连接成功")
	// 	// if err != nil {
	// 	// 	fmt.Println("报错：Error Listening: %v", err)
	// 	// 	return nil, fmt.Errorf("Error Listening: %v", err)
	// 	// }
	// 	// fmt.Printf("Listening on %s:%d")
	// 	// fmt.Printf("断电")
	// 	// fmt.Println(TcpAddr + ":" + strconv.Itoa(TcpPort))
	// 	// var conn *net.TCPConn
	// 	// defer func(conn *net.TCPConn) {
	// 	// 	err := conn.Close()
	// 	// 	if err != nil {
	// 	// 		fmt.Println("conn close failed, err:", err)
	// 	// 	}
	// 	// }(conn)
	// 	// if err != nil {
	// 	// 	fmt.Println("连接服务器失败：", err)
	// 	// 	return nil, fmt.Errorf("连接服务器失败：", err)
	// 	// }
	// 	// fmt.Println("连接服务器成功")
	// 	// defer conn.Close()
	// }
}

func handle(conn net.Conn, data []string) ([]string, error) {
	defer conn.Close()
	buffer := make([]byte, 124)
	fmt.Println("创建buffer成功")
	n, err := conn.Read(buffer)
	conn.Close()
	if err != nil {
		fmt.Println("读取数据失败：", err)
	}
	// 处理数据
	fmt.Println("buffer:", string(buffer[:n]))
	
	var messages []string
	// 去掉数据中的空格
	msg := strings.Replace(string(buffer[0:n]), " ", "", -1)
	// 将处理后的数据存储到字符串数组中
	messages = append(messages, msg)
	fmt.Println("messages:", messages)
	// 将messages中的数据每108位分割一次，存储到data中
	// 将result中的空格删除，变成新的字符串
	data = messages
	// data = append(data, messages[i:i+Length]...)
	fmt.Println("go handle data:", data)
	return data, nil
}
