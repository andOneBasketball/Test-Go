package main
// 闭包导致内存逃逸

import "fmt"

func f() int {
	var x int = 1
	return func(x int) int {
		return x + x
	}(x)
}

func main() {
	fmt.Println(f())
	fmt.Println(f())
}
