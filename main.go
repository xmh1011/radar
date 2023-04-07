package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"radar/pkg"
	"radar/udp"
	"strconv"
)

type Options struct {
	Stderr io.Writer
	Stdout io.Writer
	
	UdpAddr string
	UdpPort int
	host    string
	dbName  string
}

func NewOptions() *Options {
	return &Options{
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}
}

var opt = NewOptions()

func main() {
	mainCmd := Run()
	
	if err := mainCmd.Execute(); err != nil {
		fmt.Printf("Error : %+v\n", err)
	}
}

// 此代码为正式代码
func Run() *cobra.Command {
	
	c := &cobra.Command{
		Use:  "radar",
		Long: "receive radar data and restore to database",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   true,
			DisableNoDescFlag:   true,
			DisableDescriptions: true},
		Run: func(cmd *cobra.Command, args []string) {
			// 创建新的数据库定义
			db := pkg.NewCnosDB(opt.host, opt.dbName)
			// 在CnosDB中创建数据库
			err := pkg.CreateDatabase(db.Database, db.URL)
			
			if err != nil {
				fmt.Printf("Error creating database:")
				log.Fatal(err)
				return
			}
			
			for {
				func() {
					
					data, err := udp.ReceiveData(opt.UdpAddr, opt.UdpPort)
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
						CnosData = "radar" + "," + "Header=" + cnos.Header + "," + "SendNode=" + cnos.SendNode + "," + "ReceiveNode=" + cnos.ReceiveNode + "," + "Method=" + cnos.Method + "," + "Status=" + cnos.Status + "," + "Tail=" + cnos.Tail + " " + "Order=" + Process(cnos.Order) + "," + "Batch=" + Process(cnos.Batch) + "," + "Distance=" + Process(cnos.Distance) + "," + "Orientation=" + Process(cnos.Orientation) + "," + "Course=" + Process(cnos.Course) + "," + "Speed=" + Process(cnos.Speed) + "," + "Longitude=" + Process(cnos.Longitude) + "," + "Latitude=" + Process(cnos.Latitude) + " " + strconv.FormatInt(cnos.Time, 10)
						// 将数据写入到CnosDB中
						err := pkg.WriteDataToCnosDB(CnosData, opt.host, opt.dbName)
						if err != nil {
							fmt.Printf("Error writing:")
							log.Fatal(err)
							return
						}
					}
					
				}()
			}
		},
	}
	
	c.PersistentFlags().StringVar(&opt.UdpAddr, "addr", "127.0.0.1", "Set the udp address")
	c.PersistentFlags().IntVar(&opt.UdpPort, "port", 8000, "Set the udp port")
	c.PersistentFlags().StringVar(&opt.host, "host", "http://localhost:8086", "Set the database host")
	c.PersistentFlags().StringVar(&opt.dbName, "dbName", "radar", "Set the database name")
	
	return c
}

func Process(cnos string) string {
	return strconv.FormatInt(pkg.HexToDec64(cnos), 10)
}
