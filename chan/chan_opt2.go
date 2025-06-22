package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan string, 10)
	done := make(chan struct{}) // 终止信号
	ch <- "ch"

	go func() {
		defer wg.Done()
		for {
			select {
			case msg, ok := <-ch:
				if !ok {
					return
				}
				if msg == "ch" {
					fmt.Println("a")
				}
				select {
				case ch <- "num": // 尝试发送
				case <-done: // 收到终止信号
					return
				}
			case <-done: // 直接退出
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case msg, ok := <-ch:
				if !ok {
					return
				}
				if msg == "num" {
					fmt.Println(1)
				}
				select {
				case ch <- "ch":
				case <-done:
					return
				}
			case <-done:
				return
			}
		}
	}()

	time.Sleep(1 * time.Second)
	close(done) // 发送终止信号
	wg.Wait()   // 等待 Goroutine 退出
	close(ch)   // 此时 Channel 已无人使用，关闭安全
}
