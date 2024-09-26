package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("err1")
	err2 := errors.New("err2")
	fmt.Println(err1)
	fmt.Println(err2)
}