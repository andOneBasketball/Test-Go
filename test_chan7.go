package main

import (
	"fmt"
)

func main() {
	//ch := make(chan int, 1)
	/*
		ch := make(chan int) // 死锁
		ch <- 10
		ch <- 20
		fmt.Println(<-ch)
	*/

	var ch chan *int
	ch = make(chan *int, 10)
	ch <- new(int)
	close(ch)
	fmt.Println(len(ch))

}
