package main

import "fmt"

func f2() (int, bool) {
	//data, f := f1()
	return 3, false
}

func main() {
	a, b := f2()
	fmt.Printf("%v %v", a, b)
}
