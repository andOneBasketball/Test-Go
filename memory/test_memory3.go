package main

// 创建过大的切片

import (
	"fmt"
)

func main() {
	a := make([]int, 1)
	b := make([]int, 1000000000)
	fmt.Println(len(a), cap(b))
}