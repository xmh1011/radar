package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var Length = 108

func main() {
	// 设置UDP地址
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("ResolveUDPAddr failed:", err)
		return
	}
	
	// 定义发送数据
	var FilePath = "/Users/xiaominghao/code/radar/"
	var FileName = "radar.txt"
	
	filename := FilePath + FileName
	data := ReadDataFromFile(filename)
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
		defer conn.Close()
		_, err = conn.Write([]byte(temp))
		if err != nil {
			fmt.Println("Write failed:", err)
			return
		}
		fmt.Println("Send data successfully!")
		
		time.Sleep(1 * time.Second) // 暂停1秒钟再发送下一次数据
	}
}

// 从txt文件中读取数据
func ReadDataFromFile(filename string) (data []string) {
	var result string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		result = scanner.Text()
	}
	// fmt.Println("result: ", result)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// // 将result中的数据按照空格分割
	// temp := strings.Split(result, " ")
	// // 将temp中的数据每54位分割一次，存储到data中
	// fmt.Println("temp: ", temp)
	// fmt.Println("len(temp): ", len(temp))
	// 将result中的空格删除，变成新的字符串
	result = strings.Replace(result, " ", "", -1)
	// 将result中的数据，每108位分割一次，存储到data中
	for i := 0; i < len(result); i += Length {
		// data = append(data, strings.Join(result[i:i+Length], ""))
		data = append(data, result[i:i+Length])
	}
	return data
}
