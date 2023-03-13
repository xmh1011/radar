package pkg

import (
	"io"
	"log"
	"net"
	"os"
)

// 建立一个tcp链接，并向该端口发送数据
func Dial() {
	conn, err := net.Dial("tcp", ":8000") // 创建一个连接，连接到端口 8000
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // 忽略错误
		log.Println("done")
		done <- struct{}{} // 通知主goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // 等待后台goroutine完成
}

// mustCopy是一个io.Copy的包装器，如果io.Copy返回一个错误，那么就会调用log.Fatal函数，终止程序的执行
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
