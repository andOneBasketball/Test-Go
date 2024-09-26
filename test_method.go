package main

import "fmt"

type T int64
func (t T) F() {
	fmt.Println("F called")
}

func main() {
    t := T(10)
    t.F()
}