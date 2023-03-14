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
		func() {
			
			data, err := udp.ReceiveData(UdpAddr, UdpPort)
			if data == nil {
				fmt.Printf("Failed to receive data!")
			}
			if err != nil {
				fmt.Printf("Error reading:")
				log.Fatal(err)
				return
			}
			if data != nil {
				fmt.Println("Received data successfully!")
			}
			// 将data转化为符合SQL行协议的数据
			var CnosData string
			for _, v := range data {
				cnos := pkg.HandleData(v)
				CnosData = "radar" + "," + "Header=" + cnos.Header + "," + "SendNode=" + cnos.SendNode + "," + "ReceiveNode=" + cnos.ReceiveNode + "," + "Method=" + cnos.Method + "," + "Status=" + cnos.Status + "," + "Tail=" + cnos.Tail + " " + "Time=" + Process(cnos.Time) + "," + "Order=" + Process(cnos.Order) + "," + "Batch=" + Process(cnos.Batch) + "," + "Distance=" + Process(cnos.Distance) + "," + "Orientation=" + Process(cnos.Orientation) + "," + "Course=" + Process(cnos.Course) + "," + "Speed=" + Process(cnos.Speed) + "," + "Longitude=" + Process(cnos.Longitude) + "," + "Latitude=" + Process(cnos.Latitude) + " " + strconv.FormatInt(time.Now().UnixNano(), 10)
				// 将数据写入到CnosDB中
				err := pkg.WriteDataToCnosDB(CnosData, CnosURL, CnosDatabase)
				if err != nil {
					fmt.Printf("Error writing:")
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
