package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CnosDB 数据结构
type CnosDB struct {
	URL      string    // CnosDB 的 URL
	Database string    // 数据库名称
	Data     *CnosData // 数据
}

// NewCnosDB 创建 CnosDB 实例
func NewCnosDB(url, database string) *CnosDB {
	return &CnosDB{
		URL:      url,
		Database: database,
		Data:     &CnosData{Measurement: "radar"},
	}
}

// CreateDatabase 在CnosDB中创建一个新的数据库
func CreateDatabase(database string, url string) error {

	// 创建一个POST请求，将创建指令作为URL编码的数据发送
	// 设置请求参数
	url = url + "/query"
	query := "CREATE DATABASE " + database
	data := []byte(fmt.Sprintf("q=%s", query))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error making request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing body:", err)
		}
	}(resp.Body)

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("CnosDB create database error: %s", resp.Status)
	}

	return err
}

// WriteDataToCnosDB 写入数据到 CnosDB
func WriteDataToCnosDB(data string, url string, database string) error {

	mu.Lock()

	_, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Invalid data: %s", err)
	}

	// 构造请求
	req, err := http.NewRequest("POST", url+"/write?db="+database, bytes.NewBufferString(data))

	// 执行curl
	if err != nil {
		return fmt.Errorf("Invalid request: %s", err)
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error making request: %s", err)
	}

	// 关闭响应体
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing body:", err)
		}
	}(resp.Body)

	// 检查响应状态码
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("CnosDB write error: %s", resp.Status)
	}

	mu.Unlock()

	return nil
}
