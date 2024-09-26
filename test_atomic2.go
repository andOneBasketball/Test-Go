package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"os"
)

func main() {
	var count int32 = 0
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				atomic.AddInt32(&count, 1)
				//fmt.Println(i, "count:", count)
			}
		} (i)
	}

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("count:", atomic.LoadInt32(&count))
		os.Exit(0)
	} ()

	wg.Wait()
}
