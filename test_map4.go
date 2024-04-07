package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := map[string][]string{
		"a": []string{"123", "123123"},
	}
	fmt.Printf("%#v\n", a)

	b := map[string]interface{}{"a": 123, "b": "qw"}
	bByte, _ := json.Marshal(b)
	fmt.Printf("str: %#v\n", string(bByte))

	c := "{\"abcd\": \"1234\"}"
	fmt.Printf("%#v\n", c)
}
