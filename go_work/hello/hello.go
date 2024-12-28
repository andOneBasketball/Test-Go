package main

import (
	"fmt"
	"hello1/fun"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"))
	fun.SayHello()
}
