package main

import "fmt"

func main() {
    var m map[string]int = map[string]int{}
	m["a"] = 1
	fmt.Println(m)
	delete(m, "a")
	fmt.Println(m)
}
