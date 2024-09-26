package main

import (
	"fmt"
	_ "kpkg"
	_ "kpk"
)

var a = test()

func test() int {
	fmt.Println("test() called\n")
	return 10
}

func init() {
	a = 20
	fmt.Printf("init() called %d\n", a)
}

func main() {
	fmt.Println(a)
}