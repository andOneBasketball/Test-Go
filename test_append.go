package main

import (
	"fmt"
)

type statistics struct {
	Md5   string `bson:"_id"`
	Count uint64 `bson:"count"`
}

type data struct {
	s  []*statistics
	dd uint32
}

func main() {
	d := new(data)
	d.dd = 1
	a := statistics{"qwe", 22}
	d.s = append(d.s, &a)
	fmt.Println(*d)
	fmt.Println(*d.s[0])

	var f uint32 = 123
	fmt.Println(f)
}
