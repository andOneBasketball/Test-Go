package main

import (
	"fmt"
	_ "init_opt/p1"
	_ "init_opt/p2"
)

func init() {
	fmt.Println("init_opt1")
}

func main() {
	fmt.Println("hello")
}
