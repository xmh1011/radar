package main

import (
	"fmt"
	"log"
	"radar/pkg"
	"radar/udp/udp"
	"strconv"
	"time"
)

var UdpAddr = "127.0.0.1"
var UdpPort = 8000
var CnosURL = "http://localhost:8086"
var CnosDatabase = "radar"

// 此代码为正式代码
func main() {
	// 创建新的数据库定义
	db := pkg.NewCnosDB(CnosURL, CnosDatabase)
	// 在CnosDB中创建数据库
	err := pkg.CreateDatabase(db.Database, db.URL)
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		fmt.Println(1)
		func() {
			fmt.Println(2)
			data, err := udp.ReceiveData(UdpAddr, UdpPort)
			if data == nil {
				log.Println("接收数据失败")
			}
			if err != nil {
				log.Fatal(err)
				return
			}
			// 将data转化为符合SQL行协议的数据
			var CnosData string
			for _, v := range data {
				cnos := pkg.HandleData(v)
				CnosData = "radar" + "," + "Header=" + cnos.Header + "," + "SendNode=" + cnos.SendNode + "," + "ReceiveNode=" + cnos.ReceiveNode + "," + "Method=" + cnos.Method + "," + "Status=" + cnos.Status + "," + "Tail=" + cnos.Tail + " " + "Time=" + Process(cnos.Time) + "," + "Order=" + Process(cnos.Order) + "," + "Batch=" + Process(cnos.Batch) + "," + "Distance=" + Process(cnos.Distance) + "," + "Orientation=" + Process(cnos.Orientation) + "," + "Course=" + Process(cnos.Course) + "," + "Speed=" + Process(cnos.Speed) + "," + "Longitude=" + Process(cnos.Longitude) + "," + "Latitude=" + Process(cnos.Latitude) + " " + strconv.FormatInt(time.Now().UnixNano(), 10)
				// fmt.Println("CnosData: ", CnosData)
				// 将数据写入到CnosDB中
				err := pkg.WriteDataToCnosDB(CnosData, CnosURL, CnosDatabase)
				if err != nil {
					log.Fatal(err)
					return
				}
			}
		}()
	}
}

func Process(cnos string) string {
	return strconv.FormatInt(pkg.HexToDec64(cnos), 10)
}

// // 此代码为tcp连接测试代码
// var TcpAddr = "localhost"
// var TcpPort = 8000
//
// func main() {
// 	// data := "radar,Header=FFFF,SendNode=0000,ReceiveNode=0000,Method=00,Status=01,Tail=F7F7 Time=9223372036854775807,Order=0,Batch=43057152,Distance=3172139008,Orientation=27483,Course=20042,Speed=7936,Longitude=2761077320,Latitude=270555670 1678699548373218000"
// 	// 读取文件中的数据
// 	// var FilePath = "/Users/xiaominghao/code/radar/"
// 	// var FileName = "radar.txt"
// 	//
// 	// filename := FilePath + FileName
// 	// data := pkg.ReadDataFromFile(filename)
// 	// fmt.Println("1.data: ", data)
// 	// go func() {
// 	// 	err := pkg.Server(TcpAddr, TcpPort, data)
// 	// 	if err != nil {
// 	//
// 	// 	}
// 	// }()
// 	// fmt.Println("数据发送完毕")
// 	// // if err != nil {
// 	// // 	log.Fatal(err)
// 	// // 	return
// 	// // }
// 	// fmt.Println("即将接收数据")
// 	// fmt.Println("开始接收数据")
//
// 	data, err := pkg.ReceiveDataTCP(TcpAddr, TcpPort)
// 	fmt.Println("2.data: ", data)
// 	if data == nil {
// 		log.Println("接收数据失败")
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	// 将data转化为符合SQL行协议的数据
// 	var CnosData string
// 	for _, v := range data {
// 		fmt.Println(3)
// 		info := pkg.HandleData(v)
// 		CnosData = "radar" + "," + "Header=" + info.Header + "," + "SendNode=" + info.SendNode + "," + "ReceiveNode=" + info.ReceiveNode + "," + "Method=" + info.Method + "," + "Status=" + info.Status + "," + "Tail=" + info.Tail + " " + "Time=" + Process(info.Time) + "," + "Order=" + Process(info.Order) + "," + "Batch=" + Process(info.Batch) + "," + "Distance=" + Process(info.Distance) + "," + "Orientation=" + Process(info.Orientation) + "," + "Course=" + Process(info.Course) + "," + "Speed=" + Process(info.Speed) + "," + "Longitude=" + Process(info.Longitude) + "," + "Latitude=" + Process(info.Latitude) + " " + strconv.FormatInt(time.Now().UnixNano(), 10)
// 		fmt.Println("CnosData: ", CnosData)
// 		// 将数据写入到CnosDB中
// 		err := pkg.WriteDataToCnosDB(CnosData, CnosURL, CnosDatabase)
// 		if err != nil {
// 			log.Fatal(err)
// 			return
// 		}
// 	}
//
// 	fmt.Println("done")
// }

// func main() {
// 	ip := "localhost" // 指定IP地址
// 	port := "8000"    // 指定端口号
//
// 	// 创建TCP连接
// 	addr := ip + ":" + port
// 	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
// 	if err != nil {
// 		fmt.Println("ResolveTCPAddr error:", err)
// 		return
// 	}
// 	conn, err := net.DialTCP("tcp", nil, tcpAddr)
// 	if err != nil {
// 		fmt.Println("DialTCP error:", err)
// 		return
// 	}
// 	defer conn.Close()
//
// 	// 持续接收数据
// 	for {
// 		fmt.Println("Waiting for data...")
// 		data := make([]byte, 1024)
// 		_, err := conn.Read(data)
// 		if err != nil {
// 			fmt.Println("Read error:", err)
// 			return
// 		}
// 		fmt.Println("Received data:", string(data))
// 	}
// }

// // 从指定文件中读取信息时使用此代码
// func main() {
// 	// 读取文件中的数据
// 	var FilePath = "/Users/xiaominghao/code/radar/"
// 	var FileName = "radar.txt"
//
// 	filename := FilePath + FileName
// 	data := pkg.ReadDataFromFile(filename)
// 	// data := pkg.ReceiveData(UdpAddr, UdpPort)
// 	// 将data转化为符合SQL行协议的数据
// 	var CnosData string
// 	for _, v := range data {
// 		info := pkg.HandleData(v)
// 		CnosData = "radar" + "," + "Header=" + info.Header + "," + "SendNode=" + info.SendNode + "," + "ReceiveNode=" + info.ReceiveNode + "," + "Method=" + info.Method + "," + "Status=" + info.Status + "," + "Tail=" + info.Tail + " " + "Time=" + Process(info.Time) + "," + "Order=" + Process(info.Order) + "," + "Batch=" + Process(info.Batch) + "," + "Distance=" + Process(info.Distance) + "," + "Orientation=" + Process(info.Orientation) + "," + "Course=" + Process(info.Course) + "," + "Speed=" + Process(info.Speed) + "," + "Longitude=" + Process(info.Longitude) + "," + "Latitude=" + Process(info.Latitude) + " " + strconv.FormatInt(time.Now().UnixNano(), 10)
// 		fmt.Println("CnosData: ", CnosData)
// 		// 将数据写入到cnosdb中
// 		err := pkg.WriteDataToCnosDB(CnosData, CnosURL, CnosDatabase)
// 		if err != nil {
// 			log.Fatal(err)
// 			return
// 		}
// 	}
// }
