package main

import (
	"fmt"
	"time"
	"github.com/karlseguin/ccache/v2"
)

func main() {
	// 创建一个新的缓存实例
	cache := ccache.New(ccache.Configure().MaxSize(3))
	
	// 添加一个键值对到缓存中
	cache.Set("key", uint8(1), time.Second * 10)
	cache.Set("key1", uint8(1), time.Second * 10)
	cache.Set("key2", uint8(1), time.Second * 10)

	// 从缓存中获取值
	item := cache.Get("key")
	if item != nil {
		fmt.Println("从缓存中获取到的值:", item.Value().(uint8))
	} else {
		fmt.Println("缓存中没有找到对应的值")
	}
	item = cache.Get("key")
	if item != nil {
		fmt.Println("从缓存中获取到的值:", item.Value().(uint8))
	} else {
		fmt.Println("缓存中没有找到对应的值")
	}
	//time.Sleep(time.Second * 10)

	cache.Set("key3", uint8(1), time.Second * 10)
	cache.ForEachFunc(func(key string, i *ccache.Item) bool {
		fmt.Printf("%v %#v\n", key, i)
		return true
	})

	// 从缓存中获取值
	item = cache.Get("key")
	if item != nil {
		fmt.Println("从缓存中获取到的值:", item.Value().(uint8))
	} else {
		fmt.Println("缓存中没有找到对应的值")
	}

	// 等待一段时间，使缓存项过期
	// time.Sleep(time.Second * 40)

	item = cache.Get("key")
	if item != nil {
		fmt.Println("11从缓存中获取到的值:", item.Value().(uint8))
	} else {
		fmt.Println("111缓存中没有找到对应的值")
	}

	// 再次尝试获取值
	cache2 := ccache.New(ccache.Configure().MaxSize(1000).ItemsToPrune(100))

	cache2.Set("abcd", uint8(1), time.Second)
	item = cache2.Get("abcd")
	if item != nil {
		fmt.Println("从缓存中获取到的值:", item.Value().(uint8))
	} else {
		fmt.Println("缓存中没有找到对应的值")
	}
}