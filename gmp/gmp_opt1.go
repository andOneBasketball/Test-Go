package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	numP := runtime.GOMAXPROCS(0)
	fmt.Printf("当前P的数量: %d\n", numP)

	var wg sync.WaitGroup
	var mu sync.Mutex
	maxNum := 260
	wg.Add(maxNum)
	for i := 0; i < maxNum; i++ {
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			fmt.Printf("%d ", num)
			mu.Unlock()
		}(i + 1)
	}
	wg.Wait()
}
