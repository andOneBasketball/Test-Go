package main

import "fmt"

func main() {
	var m map[int]int
	fmt.Println(m[0])
	m[0] = 1
	fmt.Println(m[0])
}
