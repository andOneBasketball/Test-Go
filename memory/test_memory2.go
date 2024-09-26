package main
// 返回局部变量的指针

import "fmt"

func f() *int {
	var x int = 1
	return &x
}

func main() {
	fmt.Println(f())
}
