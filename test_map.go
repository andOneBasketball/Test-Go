package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// 文件指纹信息数据结构
type MInfo struct {
	Product            string `json:"product"`
	Devid              string `json:"devid"`
}

func main() {
	f1 := MInfo{
		Product: "AAa",
		Devid: "BBb",
	}
	var f2 [5]MInfo
	f2[0] = f1
	for _, f := range f2 {
		m, _ := json.Marshal(f)
		fmt.Println(string(m))
	}
	msg, _ := json.Marshal(f1)
	fmt.Println(string(msg))

	m1 := make(map[string]interface{})
	m1 = map[string]interface{}{
		"abcd": f1,
		"a": f1,
		"d": "abcd",
	}
	fmt.Println(len(m1), m1)
	m1["c"] = m1["a"]
	delete(m1, "a")
	if d, ok := m1["d"].(string); ok {
		m1["d"] = strings.ToUpper(d)
	}
	//m1["d"] = string(d)
	m1["e"] = "12353"
	fmt.Println(len(m1), m1)

	for key, _ := range m1 {
		fmt.Printf("key: %s\n", key)
	}
}
