package main

import (
    "fmt"
    "os"
)

func main() {
    // 获取所有命令行参数
    args := os.Args

    // 打印程序名称
    fmt.Println("程序名称：", args[0])

    // 打印其他参数
    for i := 1; i < len(args); i++ {
        fmt.Printf("参数 %d, %s\n", i, args[i])
    }
}
