package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("开始 select 循环")

	select {
	// 没有任何 case 就绪的情况下，会阻塞在 select 这里
	default:
		// 如果加了 default，就不会阻塞，会变成忙循环
		fmt.Println("default 执行了（不会阻塞）")
		time.Sleep(1 * time.Second)
	}
}
