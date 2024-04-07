package main

import (
    "sync/atomic"
    "time"
	"fmt"
)

func main() {
    // 定义计数器变量
    var count int64

    // 定时器的回调函数
    go func() {
        for {
			time.Sleep(1 * time.Second)
            // 原子地读取计数器变量的值，并将其重置为0
            // currentCount := atomic.SwapInt64(&count, 0)
            // 输出计数器变量的值
			currentCount := atomic.LoadInt64(&count)
            fmt.Printf("24小时内消费的消息数量：%d\n", currentCount)
        }
    }()

    // 消费Kafka消息
    for {
        time.Sleep(100 * time.Millisecond)

        // 原子地增加计数器变量的值
        atomic.AddInt64(&count, 2)
    }
}