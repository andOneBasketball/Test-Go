package main

import "fmt"

func f1() (int, bool) {
	return 3, false
}

func F2() (int, bool) {
	//data, f := f1()
	return f1()
}


func main() {
	a, b := F2()
	fmt.Printf("%v %v", a, b)
}
