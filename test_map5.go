package main

import (
	"fmt"
)

func main() {
	a := make(map[string]interface{})
	a["md5"] = "123214"
	if a["md5"] == "123214" {
		fmt.Printf("%s\n", a["md5"])
	} else {
		fmt.Printf("aaaaa")
	}

	b := map[string]string{
		"acd": "123",
		"123": "fewfew",
	}
	fmt.Printf("%#v\n", b)
	for k, v := range b {
		fmt.Printf("k:%v v:%v\n", k, v)
		b[k] = "31543"
	}
	fmt.Printf("%#v\n", b)
}

