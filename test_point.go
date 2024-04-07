package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = flag.String("name", "defaultName", "Name to greet")
	fmt.Println(name)
	fmt.Println(*name)
	a := 1
	b := &a
	fmt.Printf("%d %d\n", *b+1, b)
}