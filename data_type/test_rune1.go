package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "hello"
	var r rune = rune(str[2])

	fmt.Printf("%s %c %s %s\n", str, r, reflect.TypeOf(str[2]), reflect.TypeOf(r))
}
