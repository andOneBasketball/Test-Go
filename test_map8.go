package main

import (
	"fmt"
	"log"
)

func main() {
	var m map[string]int = map[string]int{}
	m["a"] = 1
	fmt.Println(m)
	delete(m, "a")
	fmt.Println(m)

	m1 := make(map[string]int)
	m1["b"]++
	fmt.Println(m1)
	m1["b"]++
	fmt.Println(m1)
	m1["a"]++
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	fmt.Println(m1["b"], m1["c"])

	tokenFishFlagMap := make(map[int]bool)
	log.Printf("%v %v", tokenFishFlagMap[9], tokenFishFlagMap[10])
}
