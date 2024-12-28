package main

import (
	"fmt"

	"github.com/tidwall/sjson"
)

func main() {
	jsonStr := `{"name": "Alice", "details": {"age": 25, "city": "Wonderland"}}`

	// 修改嵌套对象中的值
	updatedJSON, err := sjson.Set(jsonStr, "details.city", "Dreamland")
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated JSON:", updatedJSON)

	// 添加新的嵌套字段
	updatedJSON, err = sjson.Set(jsonStr, "details.country", "Fantasyland")
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated JSON with new field:", updatedJSON)

	// 添加全新的嵌套对象
	updatedJSON, err = sjson.Set(jsonStr, "details.address", map[string]interface{}{
		"street": "Elm St.",
		"zip":    "12345",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated JSON with new object:", updatedJSON)
}
