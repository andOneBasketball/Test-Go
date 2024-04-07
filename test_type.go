package main

import (
	"fmt"
)

type feedbackResData struct {
	Path     string `json:"path"`
	Code     int32  `json:"code"`
	Descript string `json:"descript"`
}

func main() {
	var a feedbackResData
	a.Path = "abcd"
	a.Code = 200
	a.Descript = "wefw"
	fmt.Println(a)

	max := 3
	if max > 2 {
		max--
		fmt.Println(max)
	}
}
