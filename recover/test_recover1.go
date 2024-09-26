package main

import (
	"fmt"
	"time"
)

func f() {
	defer func() {
		if err := recover(); err!= nil {
			fmt.Println("Recovered in f", err)
		}
	}()

	a := 10
	d := 0
	c := a / d
	fmt.Println(a, c, d)
}

func main() {
	go f()
	time.Sleep(5 * time.Second)
	fmt.Println("main function ends")
}