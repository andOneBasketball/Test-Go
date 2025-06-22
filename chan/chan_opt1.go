package main

import (
	"fmt"
	"sync"
)

// 合适的关闭 chan 的时机

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int, 10)
	maxNum := 20
	//count := 0
	ch <- 1

	go func() {
		defer wg.Done()
		for {
			select {
			case msg, ok := <-ch:
				if !ok {
					return
				}
				if msg > maxNum {
					close(ch)
					return
				}
				if msg%2 == 0 {
					fmt.Println(fmt.Sprintf("a 1 %d", msg))
					ch <- msg + 1
				} else if msg%2 == 1 {
					fmt.Println(fmt.Sprintf("1 1 %d", msg))
					ch <- msg + 1
				}
			default:
				//fmt.Println("no data 1")
				break
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case msg, ok := <-ch:
				if !ok {
					return
				}
				if msg > maxNum {
					close(ch)
					return
				}
				if msg%2 == 0 {
					fmt.Println(fmt.Sprintf("a 2 %d", msg))
					ch <- msg + 1
				} else if msg%2 == 1 {
					fmt.Println(fmt.Sprintf("1 2 %d", msg))
					ch <- msg + 1
				}
			default:
				break
			}
		}
	}()
	wg.Wait()
}
