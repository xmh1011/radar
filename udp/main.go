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
	serverAddr1, err := net.ResolveUDPAddr("udp", "127.0.0.1:8000")
	serverAddr2, err := net.ResolveUDPAddr("udp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("ResolveUDPAddr failed:", err)
		return
	}

	// 定义发送数据
	var FilePath = "./"
	var FileName1 = "radar.txt"
	var FileName2 = "radar2.txt"

	filename1 := FilePath + FileName1
	filename2 := FilePath + FileName2
	data1 := ReadDataFromFile(filename1)
	data2 := ReadDataFromFile(filename2)
	// 将[]string转换为[]byte
	temp := strings.Join(data1, "\n")
	temp2 := strings.Join(data2, "\n")
	// 持续发送数据
	for {
		// 创建UDP连接
		conn1, err := net.DialUDP("udp", nil, serverAddr1)
		conn2, err := net.DialUDP("udp", nil, serverAddr2)

		if err != nil {
			fmt.Println("DialUDP failed:", err)
			return
		}
		_, err = conn1.Write([]byte(temp))
		_, err = conn2.Write([]byte(temp2))
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
