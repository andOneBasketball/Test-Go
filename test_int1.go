package main

import "fmt"

func main() {
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)
	a, b = b, a

	fmt.Println(a, b, c)
}