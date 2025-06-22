package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()

	num := 10
	wg := sync.WaitGroup{}
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case data, ok := <-ch:
				if ok {
					goNum := runtime.NumGoroutine()
					fmt.Println(data, goNum)
					if data+1 > num {
						close(ch)
						return
					}
					ch <- data + 1
				}
			}
			//fmt.Println(fmt.Sprintf("end %d", i))
		}()
	}
	goNum := runtime.NumGoroutine()
	fmt.Println(goNum)
	wg.Wait()
}
