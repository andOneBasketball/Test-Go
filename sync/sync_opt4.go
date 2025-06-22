package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	num := 1
	chInt := make(chan int)
	go func() {
		chInt <- num
	}()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			select {
			case data, ok := <-chInt:
				if ok {
					fmt.Println(i, data)
					if data >= 10 {
						close(chInt)
						return
					}
					chInt <- data + 1
				}
			}
		}()
	}
	wg.Wait()
}
