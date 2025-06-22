package main

import (
	"log"
	"time"
)

func main() {
	local, err := time.LoadLocation("Asia/Shanghai")
	log.Printf("local: %v, err: %v", local, err)
	// 创建一个定时器，每隔 2 秒触发一次
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop() // 在程序结束时停止定时器

	// 启动一个循环，监听定时器通道
	for {
		<-ticker.C
		log.Println("hello")
	}
}
