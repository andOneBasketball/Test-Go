package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	count := 0
	mu := sync.Mutex{}
	go func() {
		for {
			mu.Lock()
			count++
			mu.Unlock()
			//fmt.Println("go Count:", count)
		}
	} ()
	time.Sleep(1 * time.Second)
	fmt.Println("Count:", count)
}
