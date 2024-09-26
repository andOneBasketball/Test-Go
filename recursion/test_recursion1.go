package main

import (
	"fmt"
)

var s = 0

func re(n int) {
	sum := 0
	for i := n; i > 0; i-- {
		sum += i
		fmt.Println(sum, i, n, s)
		re(n - 1)
		s++
	}
	fmt.Println("end ", sum)
}

func main() {
	re(4)
	fmt.Println("s: ", s)
}

