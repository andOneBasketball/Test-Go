package main

import (
	"fmt"
	"time"

	"github.com/karlseguin/ccache/v2"
)

func main() {
	cache := ccache.New(ccache.Configure())

	// 添加一个缓存项，并设置过期时间为1秒
	cache.Set("key1", "value1", 0)

	// 在缓存项过期前获取值
	val := cache.Get("key1")
	fmt.Println("Value before expiration:", val)

	// 等待1秒，让缓存项过期
	time.Sleep(time.Second)

	// 缓存项过期后尝试获取值
	val = cache.Get("key1")
	fmt.Println("Value after expiration:", val)
}
