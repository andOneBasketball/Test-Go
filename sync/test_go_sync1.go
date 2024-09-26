/*
按顺序打印1~10
*/

package main

import (
	"fmt"
	"sync"
)

func createGoroutine(num int, maxNum int, ch chan int) {
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 1; i <= num; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case data, ok := <- ch:
					if !ok {
						fmt.Println(i, " goroutine exit, maxNum ", maxNum)
						return
					} else {
						if data > maxNum {
							fmt.Println(i, " goroutine exit, maxNum ", maxNum)
							close(ch)
							return
						}
						fmt.Printf("%d: %d\n", i, data)
						ch <- data + 1
					}
				default:
					//fmt.Println(i, " default goroutine exit, maxNum ", maxNum)
					break
				}
			}
		}(i)
	}
	wg.Wait()
}


func main() {
	maxNum := 10
	ch := make(chan int, maxNum)
	go func() {
		ch <- 1
	}()
	createGoroutine(5, maxNum, ch)
	fmt.Println("main exit")
}