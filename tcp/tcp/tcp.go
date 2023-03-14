package tcp

import (
	"net"
	"strings"
)

// 测试代码，做为测试能否通过tcp协议接收数据

// 从固定端口通过TCP协议接收数据
func ReceiveDataTCP(TcpAddr string, TcpPort int) (data []string, err error) {
	
	// 连接TCP服务器
	listener, err := net.Listen("tcp", "localhost:8000")
	
	if err != nil {
		return nil, err
	}
	
	// 关闭监听通道
	defer listener.Close()
	
	for {
		// 进行通道监听
		conn, err := listener.Accept()
		if err != nil {
			return nil, err
		}
		
		data, err = handle(conn, data)
		if err != nil {
			return nil, err
		}
		return data, err
	}
	
}

func handle(conn net.Conn, data []string) ([]string, error) {
	
	defer conn.Close()
	buffer := make([]byte, 124)
	
	n, err := conn.Read(buffer)
	conn.Close()
	if err != nil {
		return nil, err
	}
	
	// 去掉数据中的空格
	msg := strings.Replace(string(buffer[0:n]), " ", "", -1)
	// 将处理后的数据存储到字符串数组中
	data = append(data, msg)
	
	return data, nil
}
