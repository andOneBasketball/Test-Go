package main

import "fmt"

func main() {
	var map1 map[string]*int
	d, _ := map1["abcd"]
	if d == nil {
		fmt.Println("Key not found")
	}
}
