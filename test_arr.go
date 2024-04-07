package main

import (
	"fmt"
	"time"
)

func fun1() {
	fmt.Println("fun1")
	time.Sleep(5000 * time.Millisecond)
	go fun2()
}

func fun2() {
	fmt.Println("fun2")
}

func main() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["123"] = "1234"
	a["1"] = "12"
	fmt.Printf("%+v\n", a)
	b := []string{"2", "3", "4", "5"}
	fmt.Println(b, b[:0])

	go fun1()
	time.Sleep(1000 * time.Millisecond)
}
