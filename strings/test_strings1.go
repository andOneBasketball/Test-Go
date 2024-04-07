package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World!"
	str1 := []string{"abcd", "1234"}

	if strings.HasPrefix(str, "e") {
		fmt.Println("字符串以'Hello'开头")
	} else {
		fmt.Println("字符串不以'Hello'开头")
	}

	fmt.Printf("%d\n", len(str))
	for i, ch := range str {
		fmt.Printf("%d %c\n", i, ch)
	}

	for i, ch := range str1 {
		fmt.Printf("%d %s\n", i, ch)
	}
}
