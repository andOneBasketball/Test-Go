package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	maxGoroutines := 10
	ch := make(chan int, 11)
	ch <- 0
	for i := 0; i < maxGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case num, ok := <- ch:
					if ok {
						// fmt.Println(i, " received ", num)
						ch <- num + 1
					} else {
						//fmt.Println(i, " closed")
						return
					}
				}
			}
		} (i)
	}

	go func() {
		time.Sleep(1 * time.Second)
		select {
		case num, ok := <- ch:
			if ok {
				fmt.Println("sleep done, num:", num)
				break
			}
		}
		close(ch)
	} ()

	wg.Wait()
}
