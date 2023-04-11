package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net"
	"os"
	"radar/pkg"
	"radar/udp"
	"strconv"
	"sync"
)

type Options struct {
	Stderr io.Writer
	Stdout io.Writer

	UdpAddr string
	UdpPort int
	PortNum int
	dbHost  string
	dbName  string
}

// NewOptions returns an Options with default parameters
func NewOptions() *Options {
	return &Options{
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}
}

var opt = NewOptions()
var mu sync.Mutex

func main() {
	mainCmd := Run()

	if err := mainCmd.Execute(); err != nil {
		fmt.Printf("Error : %+v\n", err)
	}
}

func Run() *cobra.Command {

	c := &cobra.Command{
		Use:  "radar",
		Long: "receive radar data and restore to database",
		CompletionOptions: cobra.CompletionOptions{ // 自动补全
			DisableDefaultCmd:   true, // 禁用默认命令
			DisableNoDescFlag:   true,
			DisableDescriptions: true},
		Run: func(cmd *cobra.Command, args []string) {

			err := opt.NewDatabase()
			if err != nil {
				log.Fatal(err)
				return
			}

			if opt.PortNum < 1 {
				fmt.Println("PortNum must be greater than 0")
				return
			} else if opt.PortNum == 1 {
				err = opt.Handle() // 单线程处理
				if err != nil {
					return
				}
			} else {
				err = opt.Handles() // 多线程处理
				if err != nil {
					return
				}
			}
		},
	}

	c.PersistentFlags().StringVar(&opt.UdpAddr, "addr", "127.0.0.1", "Set the udp address")
	c.PersistentFlags().IntVar(&opt.UdpPort, "port", 8000, "Set the udp port")
	c.PersistentFlags().IntVar(&opt.PortNum, "num", 1, "Set the port number")
	c.PersistentFlags().StringVar(&opt.dbHost, "host", "http://localhost:8086", "Set the dbHost")
	c.PersistentFlags().StringVar(&opt.dbName, "name", "radar", "Set the database name")

	return c
}

func (opt *Options) Handle() error {

	for {
		func() {

			data, err := udp.ReceiveData(opt.UdpAddr, opt.UdpPort)
			if err != nil {
				fmt.Printf("Failed to receive data! Error reading: %s\n", err.Error())
				return
			}

			// 将data转化为符合SQL行协议的数据
			for _, v := range data {
				err = ProcessData(v)
				if err != nil {
					fmt.Printf("Error writing: %s\n", err.Error())
					return
				}
			}

		}()
	}

}

func (opt *Options) Handles() error {

	// 创建多个UDP连接并监听不同的端口
	conns := make([]*net.UDPConn, opt.PortNum)
	for i := 0; i < opt.PortNum; i++ {

		addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", i+opt.UdpPort))
		if err != nil {
			fmt.Printf("Failed to receive data! Error reading: %s\n", err.Error())
			break
		}

		conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			panic(err)
		}
		conns[i] = conn
		defer func(conn *net.UDPConn) {
			err := conn.Close()
			if err != nil {
				fmt.Printf("Failed to close connection! Error reading: %s\n", err.Error())
				return
			}
		}(conn)

	}

	// 创建多个数组，分别用于存储每个端口接收到的数据
	dataArrays := make([][]byte, opt.PortNum)
	for i := 0; i < opt.PortNum; i++ {
		dataArrays[i] = make([]byte, 1024)
	}

	// 启动多个并发协程，每个协程负责从一个UDP连接接收数据并存储到对应的数组中
	for i := 0; i < opt.PortNum; i++ {

		// 为每个协程创建一个新的变量，防止协程并发时，i的值被覆盖
		go func(idx int) {
			for {
				// 从UDP连接中读取数据
				n, _, err := conns[idx].ReadFromUDP(dataArrays[idx])
				if err != nil {
					panic(err)
				}
				// 将数据转化为符合SQL行协议的数据
				err = ProcessData(string(dataArrays[idx][:n]))
				if err != nil {
					fmt.Printf("Error writing: %s\n", err.Error())
					return
				}
			}
		}(i)

	}

	// 防止程序退出
	select {}
}

func (opt *Options) NewDatabase() error {

	// 创建新的数据库定义
	db := pkg.NewCnosDB(opt.dbHost, opt.dbName)
	// 在CnosDB中创建数据库
	err := pkg.CreateDatabase(db.Database, db.URL)

	if err != nil {
		fmt.Printf("Error creating database:%s\n", err.Error())
	}
	return err

}

func ProcessData(data string) error {

	cnos := pkg.HandleData(data)
	CnosData := "radar" + "," + "Header=" + cnos.Header + "," + "SendNode=" + cnos.SendNode + "," + "ReceiveNode=" + cnos.ReceiveNode + "," + "Method=" + cnos.Method + "," + "Status=" + cnos.Status + "," + "Tail=" + cnos.Tail + " " + "Order=" + Process(cnos.Order) + "," + "Batch=" + Process(cnos.Batch) + "," + "Distance=" + Process(cnos.Distance) + "," + "Orientation=" + Process(cnos.Orientation) + "," + "Course=" + Process(cnos.Course) + "," + "Speed=" + Process(cnos.Speed) + "," + "Longitude=" + Process(cnos.Longitude) + "," + "Latitude=" + Process(cnos.Latitude) + " " + strconv.FormatInt(cnos.Time, 10)
	// 将数据写入到CnosDB中
	err := pkg.WriteDataToCnosDB(CnosData, opt.dbHost, opt.dbName)
	if err != nil {
		return fmt.Errorf("Error writing: %s\n", err.Error())
	} else {
		fmt.Println("Write data successfully!")
	}
	return nil

}

func Process(cnos string) string {

	mu.Lock()
	result := strconv.FormatInt(pkg.HexToDec64(cnos), 10)
	mu.Unlock()
	return result

}
