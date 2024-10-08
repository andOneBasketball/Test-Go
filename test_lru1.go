package main

import (
	"fmt"
	"github.com/golang/groupcache/lru"
)

func main() {
	cache := lru.New(3)
	cache.Add("bill", 20)
	cache.Add("dable", 19)
	cache.Add("dable1", 19)
	v, ok := cache.Get("bill")
	if ok {
		fmt.Printf("bill's age is %v\n", v)
	}
	cache.Add("cat", "18")
	
	fmt.Printf("cache length is %d\n", cache.Len())
	_, ok = cache.Get("dable")
	if !ok {
		fmt.Printf("dable was evicted out\n")
	}

	cache.ForEach(func(key interface{}, value interface{}) {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	})
}