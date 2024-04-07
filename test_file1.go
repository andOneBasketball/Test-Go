package main

import (
	"fmt"
	"os"
)

// 获取文件大小，-1表示文件大小获取失败
func GetFileSize(filePath string) int64 {
	var fileInfo os.FileInfo
	var err error
	fileInfo, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Printf("file %s is missing", filePath)
		return -1
	}

	if err != nil {
		fmt.Printf("get file info from os.Stat error, err: %v", err)
		return -1
	}

	if fileInfo.IsDir() {
		fmt.Printf("file only can be file type, path %s is not match", filePath)
		return -1
	} else {
		return fileInfo.Size()
	}
}

func main() {
	fmt.Printf("%d\n", GetFileSize("./test_chan2.go"))
}