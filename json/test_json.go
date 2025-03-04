package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type ruleInfo struct {
	ACheck map[string][]string    `bson:"a_check" json:"a_check"`
	BCheck map[string]interface{} `bson:"b_check" json:"b_check"`
	Dd     string                 `json: dd`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Ff   string `json:"ff"`
}

// 判断字符串是否在字符串数组内
func InArray(arr []string, item string) bool {
	/*
		判断字符串是否在字符串数组内
		:param arr: 字符串数组
		:param item: 字符串
		:return 判断结果
	*/
	for _, value := range arr {
		if value == item {
			return true
		}
	}
	return false
}

func main() {
	submodulesValue := []byte(`{"a_check":{"enable":["1","2","3"]},"b_check":{"enable":1}}`)
	submodules := ruleInfo{}
	err := json.Unmarshal(submodulesValue, &submodules)
	fmt.Println(submodules)
	fmt.Println(submodules.ACheck)
	for key, value := range submodules.ACheck {
		fmt.Println(key, value, reflect.TypeOf(key), reflect.TypeOf(value))
		for _, vv := range value {
			fmt.Println(vv)
		}
		if InArray(value, "3") {
			fmt.Println("12354\n")
		}
	}
	fmt.Println(submodules.BCheck)
	for key, value := range submodules.BCheck {
		fmt.Println(key, value, reflect.TypeOf(key), reflect.TypeOf(value))
	}

	if len(submodules.Dd) == 0 {
		fmt.Println(submodules.Dd, reflect.TypeOf(submodules.Dd))
	}
	fmt.Println(err)

	var p Person
	p.Name = "Alice"
	p.Age = 30
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b), reflect.TypeOf(b))
}
