package main

import (
	"fmt"
)

type FileMd5Info struct {
	a  string
	b  string
}

func f2(s FileMd5Info) FileMd5Info {
	s.b = "afwe"
	return s
}

func f1(s FileMd5Info) FileMd5Info {
	s.a = "abcd"
	s = f2(s)
	return s
}

func main() {
	s := FileMd5Info{a: "1234", b: "2345"}
	var a []FileMd5Info
	a = append(a, s)
	s = f1(s)
	a = append(a, s)
	fmt.Printf("%v\n", a)
}
