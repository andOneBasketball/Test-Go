package main

import (
	"fmt"
)

func main() {
	ch := make(chan interface{})

	go func() {
		ch <- true
		close(ch)
	}()

	value, ok := <-ch
	if ok {
		str, ok := value.(bool)
		if ok {
			fmt.Println("获取到字符串:", str)
		} else {
			fmt.Println("无法转换为字符串")
		}
	} else {
		fmt.Println("通道已关闭")
	}
}