package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"reflect"
)

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}

func main() {
    // 打开文件
    file, err := os.Open("./tmp/report.json")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // 创建一个 Scanner 对象
    scanner := bufio.NewScanner(file)

	var m runtime.MemStats

    // 逐行读取文件内容
	i := 0
    for scanner.Scan() {
		i++
        line := scanner.Text()
        //fmt.Println(line)
		if line == "" {
			fmt.Println("space line")
			continue
		}
		if i >= 5 {
			break
		}
		var data map[string]interface{}
		json.Unmarshal([]byte(line), &data)
		if stat, ok := data["stat"].(map[string]interface{}); ok {
			data["file_size"] = stat["size"]
			data["file_last_access_time"] = stat["atime"]
			if mode, ok := stat["mode"].(float64); ok {
				fmt.Printf("%v\n", mode)
				data["file_access"] = fmt.Sprintf("%d", int(mode))
			}
		}
		delete(data, "stat")
		data["dev_path"] = data["file"]
		delete(data, "file")
		if md5Data, ok := data["md"].(string); ok {
			data["md5"] = strings.ToLower(md5Data)
			delete(data, "md")
		}

		data["ss"] = "test"
		jsonStr, _ := json.Marshal(data)
		fmt.Printf("%v %s\n", string(jsonStr), reflect.TypeOf(jsonStr))
		runtime.ReadMemStats(&m)
    	fmt.Printf("Alloc = %v B\n", m.Alloc)
    }

    // 检查是否有错误发生
    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
