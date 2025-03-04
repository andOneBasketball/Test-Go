package main

import (
	"encoding/json"
	"fmt"
)

// 定义结构体
type MultiRewardPropUserResource struct {
	Count    int64 `json:"count"`
	TurnOn   bool  `json:"turn_on"`
	Status   int8  `json:"status"`
	ExpireAt int64 `json:"expire_at"`
}

func main() {
	// 创建结构体实例
	resource := MultiRewardPropUserResource{
		Count:  100,
		TurnOn: true,
		Status: 1,
	}

	// 序列化为 JSON
	jsonData, err := json.MarshalIndent(resource, "", "  ")
	if err != nil {
		fmt.Printf("Error serializing struct: %v\n", err)
		return
	}

	// 打印 JSON 结果
	fmt.Println("Serialized JSON:")
	fmt.Println(string(jsonData))
}
