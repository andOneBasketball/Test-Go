package main

import (
	"fmt"
	"time"
)

func main() {
	messageChan := make(chan string, 5)   // 创建一个缓冲区为5的通道
	ticker := time.NewTicker(time.Second) // 创建一个定时器，每秒触发一次
	defer ticker.Stop()                   // 确保程序退出时停止定时器

	go func() {
		for i := 0; i < 10; i++ {
			messageChan <- fmt.Sprintf("message %d", i)
			time.Sleep(500 * time.Millisecond) // 模拟消息生成的时间间隔
		}
		close(messageChan) // 关闭通道，通知结束
	}()

	for {
		select {
		case msg, ok := <-messageChan:
			if !ok {
				fmt.Println("Channel closed, exiting.")
				return
			}
			fmt.Println("Received:", msg)
			for {
			}
		case <-ticker.C:
			fmt.Println("Timer tick:", time.Now().Format(time.RFC3339))
		}
	}
}
