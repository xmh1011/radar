package main

import (
	"fmt"
	"net"
)

const numPorts = 2

func main() {
	// 创建8个UDP连接并监听不同的端口
	conns := make([]*net.UDPConn, numPorts)
	for i := 0; i < numPorts; i++ {
		addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", i+8000))
		if err != nil {
			panic(err)
		}
		conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			panic(err)
		}
		conns[i] = conn
		defer conn.Close()
	}

	// 创建8个数组，分别用于存储每个端口接收到的数据
	dataArrays := make([][]byte, numPorts)
	for i := 0; i < numPorts; i++ {
		dataArrays[i] = make([]byte, 1024)
	}

	// 启动8个并发协程，每个协程负责从一个UDP连接接收数据并存储到对应的数组中
	for i := 0; i < numPorts; i++ {
		go func(idx int) {
			for {
				n, _, err := conns[idx].ReadFromUDP(dataArrays[idx])
				if err != nil {
					panic(err)
				}
				fmt.Printf("Received %d bytes on port %d: %s\n", n, idx+1, string(dataArrays[idx][:n]))
			}
		}(i)
	}

	// 防止程序退出
	select {}
}
