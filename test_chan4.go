package main

import (
	"fmt"
	"time"
)

func f() {
	fmt.Println(time.Now().Unix())
	i := make(chan int)
	go func() {
		i <- 10
		i <- 20
		close(i)
	}()
	
	time.Sleep(10 * time.Second)
	fmt.Println(time.Now().Unix())
	

	go func() {
		for {
			select {
			case x, ok := <-i:
				if !ok {
					//fmt.Println("channel closed")
					break
				}
				fmt.Println("channel data ", x)
			default:
				break
			}
		}
	}()
}

func main() {
	go f()
	for {}
}

