package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	for i := 0; i < len(str); i++ {
		var i8 uint8
		var b1 byte
		i8 = str[i]
		b1 = str[i]
		fmt.Printf("%c %d %b\n", str[i], i8, b1)
	}
}
