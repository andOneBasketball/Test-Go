package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i uint8
	var f bool
	var s string
	s1 := "abcefghiabcefghiabcefghi1"

	fmt.Println("Size of int:", unsafe.Sizeof(i), "bytes")
	fmt.Println("Size of float64:", unsafe.Sizeof(f), "bytes")
	fmt.Println("Size of string:", unsafe.Sizeof(s), "bytes")
	fmt.Println("Size of s1 string:", unsafe.Sizeof(s1), "bytes")
}