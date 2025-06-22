package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan string, 0)
	ch2 := make(chan string, 0)
	num := 10
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		ch1 <- "num"
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < num; i++ {
			select {
			case msg, ok := <-ch1:
				if !ok {
					close(ch2)
					return
				}
				if msg == "num" {
					fmt.Printf("1 %d\n", i)
					ch2 <- "ch"
				}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < num; i++ {
			select {
			case msg, ok := <-ch2:
				if !ok {
					return
				}
				if msg == "ch" {
					fmt.Printf("ch %d\n", i)
					if i != num-1 {
						ch1 <- "num"
					} else {
						close(ch1)
					}
				}
			}
		}
	}()
	wg.Wait()
}
