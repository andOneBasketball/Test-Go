package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.Split("hello world", ",")
	fmt.Printf("%#v, len=%d\n", s, len(s))
	s = strings.Split("hello,world", ",")
	fmt.Printf("%#v, len=%d\n", s, len(s))

	var ssss string
	if ssss == "" {
		fmt.Println(123)
	}
}
