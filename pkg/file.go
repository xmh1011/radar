package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// 从txt文件中读取数据
func ReadDataFromFile(filename string) (data []string, err error) {
	var result string
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open file failed: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		result = scanner.Text()
	}
	
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan file failed: %v", err)
	}
	
	result = strings.Replace(result, " ", "", -1)
	
	// 将result中的数据，每108位分割一次，存储到data中
	for i := 0; i < len(result); i += Length {
		data = append(data, result[i:i+Length])
	}
	
	return data, nil
}
