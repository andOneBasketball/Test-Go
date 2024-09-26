package main

import "fmt"

//加法
func add(a, b int) int {
	return a + b
}
 
//减法
func sub(a, b int) int {
	return a - b
}

func main() {
	fmt.Println(add(1, 2))
}