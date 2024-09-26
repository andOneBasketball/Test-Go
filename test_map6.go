package main

import (
	"fmt"
)

func main() {
	m1 := make(map[string]string)
	m1["john@example.com"] = "Join"
	m1["jane@example.com"] = "Jane"
	m1["bob@example.com"] = "Bob"
	fmt.Println(&m1["jane@example.com"])
	for key, value := range m1 {
		fmt.Print(key, value, &key, &value,"\n")
	}
}
