package main

import "fmt"


func main() {
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(*f())
	fmt.Println(*f())
	fmt.Println(f()==f())
}


func f() *int {
	v := 1
	v++
	return &v
}