package main

import (
	"fmt"
	"time"
	"sync"
	"os"
)

func main() {
	var count int = 0
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				mu.Lock()
				count++
				//fmt.Println(i, "count:", count)
				mu.Unlock()
			}
		} (i)
	}

	go func() {
		time.Sleep(1 * time.Second)
		mu.Lock()
		fmt.Println("sleep done,count:", count)
		mu.Unlock()
		os.Exit(0)
	}()

	wg.Wait()
}
