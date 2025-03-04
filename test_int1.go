package main

import "fmt"

func main() {
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)
	a, b = b, a

	fmt.Println(a, b, c)

	aa := []int{1, 3, 4}
	fmt.Println(aa)
	var bb []int
	bb = aa
	fmt.Println(bb)
}
