package main

import (
	"fmt"
	"time"

	"github.com/golang/groupcache/singleflight"
)

/*
缓存穿透： 当多个请求同时查询缓存不存在的值时，避免重复查询数据库。
限流合并： 将高并发访问同一资源的请求合并成一个。
重复任务的消除： 防止在同一时间触发相同的后台任务。
*/

func main() {
	var g singleflight.Group

	// 模拟多个并发请求
	for i := 0; i < 5; i++ {
		go func(id int) {
			v, err := g.Do("resource_key", func() (interface{}, error) {
				fmt.Printf("Worker %d: 正在执行实际的计算逻辑...\n", id)
				time.Sleep(2 * time.Second) // 模拟耗时操作
				return "计算结果", nil
			})
			if err != nil {
				fmt.Printf("Worker %d: 请求失败，错误: %v\n", id, err)
				return
			}
			fmt.Printf("Worker %d: 请求完成，结果: %v\n", id, v)
		}(i)
	}

	time.Sleep(5 * time.Second) // 等待所有 Goroutine 执行完成
}
