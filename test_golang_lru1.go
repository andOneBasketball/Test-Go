package main

import (
	"fmt"
	"github.com/hashicorp/golang-lru"
	"reflect"
)

func main() {
	// 创建一个能容纳 100 个元素的 LRU 缓存
	cache, err := lru.New(3)
	if err != nil {
		// 处理错误
	}

	// 向缓存中添加数据
	cache.Add("key1", uint8(1))
	cache.Add("key2", uint8(2))
	cache.Add("key3", uint8(3))
	if value, ok := cache.Get("key1"); ok {
		fmt.Println("Value for key1:", value)
	}
	cache.Add("key4", uint8(4))
	cache.Add("key3", uint8(11))

	len := cache.Len()
	fmt.Println("cache len:", len)

	// 遍历缓存中的所有键值对
	for i, k := range cache.Keys() {
		if v, ok := cache.Get(k); !ok || v != k || v != i+len {
			fmt.Printf("key: %v, value: %v, typeof: %s\n", k, v, reflect.TypeOf(v))
			if v == uint8(1) {
				v1, ok := v.(uint8)
				fmt.Println(int32(v1), reflect.TypeOf(v1), ok)
				fmt.Println("fun")
			}
		}
	}

	// 从缓存中获取数据
	if value, ok := cache.Get("key1"); ok {
		fmt.Println("Value for key1:", value)
	}
}
