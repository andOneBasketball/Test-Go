package main

import (
	"fmt"
)

func main() {
	var value interface{}
	value = 100
	//fmt.Printf("value type: %v\n", value.(type))  // value.(type) 必须结合switch一起使用
	switch value.(type) {
	case int32:
		fmt.Println(value.(int32))
	case string:
		fmt.Println(value.(string))
	case int:
		fmt.Println(value.(int))
	}
}
