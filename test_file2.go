package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    // 打开文件
    data, err := ioutil.ReadFile("./test_file1.go")
    if err != nil {
        fmt.Println("读取文件失败:", err)
        return
    }

    // 输出文件内容
    fmt.Println(string(data))
}
