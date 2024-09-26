package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		select {
		case i == 2:
			fmt.Println("i is 2")
			break
		default:
			fmt.Println("ok, i is ", i)
		}
	}
}