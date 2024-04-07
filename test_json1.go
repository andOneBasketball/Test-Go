package main

import (
	"encoding/json"
	"fmt"
)

func main() {
    type Test struct {
        Input string
    }
    regularString := map[string]interface{}{"Input": "asd\nqwe"}
	regularString["abc"] = "12314"
    out, err := json.Marshal(regularString)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(out))

	t := Test{}
	json.Unmarshal(out, &t)
	fmt.Printf("%s\n", t.Input)

	moduleData := map[string]interface{}{
		"type": "module",
	}
	moduleDataValue := map[string]interface{}{
		"abc": "123",
	}
	moduleDataValue["a"] = "123141232"
	moduleData["data"] = moduleDataValue
	fmt.Printf("%#v\n", moduleData)
	out, err = json.Marshal(moduleData)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(out))
}