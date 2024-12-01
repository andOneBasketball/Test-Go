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

	var arr []int
	arr[1] = 2
	fmt.Println(arr)

}
