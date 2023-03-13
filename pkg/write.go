package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// CnosDB 数据结构
type CnosDB struct {
	URL      string // CnosDB 的 URL
	Database string // 数据库名称
}

// NewCnosDB 创建 CnosDB 实例
func NewCnosDB(url, database string) *CnosDB {
	return &CnosDB{
		URL:      url,
		Database: database,
	}
}

// 写入数据到 CnosDB
func WriteDataToCnosDB(data string, url string, database string) error {
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
		return err
	}
	
	// 关闭响应体
	defer resp.Body.Close()
	
	// 检查响应状态码
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("CnosDB write error: %s", resp.Status)
	}
	
	return nil
}
