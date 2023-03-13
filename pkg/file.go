package pkg

import (
	"bufio"
	"log"
	"os"
	"strings"
)

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
