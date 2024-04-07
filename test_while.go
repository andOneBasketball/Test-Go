package main

import (
	"fmt"
)

func main() {
	a := 1
	for true {
		if a > 2 {
			break
		}
		a++
		fmt.Println(a)
	}
}
