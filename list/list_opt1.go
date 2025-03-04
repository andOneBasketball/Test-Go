package main

import "fmt"

func listFun(str ...string) {
	for i, v := range str {
		fmt.Println(i, v)
	}
}

func main() {
	var list []int
	for i, v := range list {
		fmt.Println(i, v)
	}
	list = append(list, 1, 2, 3, 4, 5)
	fmt.Println(list)
	listFun("")
}
