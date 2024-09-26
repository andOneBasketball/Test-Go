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

func f3(s *[]FileMd5Info) {
	*s = append(*s, FileMd5Info{a: "1234", b: "2345"})
}

func f4(s []FileMd5Info) {
	s[0] = FileMd5Info{a: "14", b: "25"}
	//s = append(s, FileMd5Info{a: "1234", b: "2345"})
}

func main() {
	s := FileMd5Info{a: "1234", b: "2345"}
	var a []FileMd5Info
	a = append(a, s)
	s = f1(s)
	a = append(a, s)
	fmt.Printf("%v\n", a)
	f4(a)
	fmt.Printf("after %v\n", a)
}
