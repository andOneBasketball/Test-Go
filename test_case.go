package main

import (
	"fmt"
)

func stringPrint(str string) {
	if str != nil {
		fmt.Println(str)
	}
}

func main() {
	a := 2
	switch a {
	case 2:
	case 3:
		fmt.Println("3")
	case 1:
		fmt.Println("1")
	}
	stringPrint("abcwefe")
	stringPrint(nil)
}