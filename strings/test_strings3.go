package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "-Hello, Wor-ld!\\u00001123"
	str1 := "-Hello, Wor-ld!\u00001123"
	str2 := `-Hello, Wor-ld!\u00001123`
	fmt.Printf("str: %s str1: %s str2: %s\n", str, str1, str2)
	fmt.Printf("%s\n%s\n", strings.Replace(str, `\u0000`, " ", -1), strings.Replace(str, "\\u0000", " ", -1))
}
