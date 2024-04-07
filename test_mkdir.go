package main

import (
	"fmt"
	"os"
)

func main() {
	localDirpath := "./path/to/dir"

	// 判断目录是否存在
	if _, err := os.Stat(localDirpath); os.IsNotExist(err) {
		// 目录不存在，创建目录
		if err := os.MkdirAll(localDirpath, 0755); err != nil {
			fmt.Println("创建目录失败：", err)
			return
		}
	}

	fmt.Println("目录已存在或创建成功")
}