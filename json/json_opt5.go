package main

import (
	"encoding/json"
	"fmt"
)

type Body struct {
	LuckyDrawType int    `json:"luckyDrawType"`
	Timestamp     int64  `json:"timestamp"`
	Signature     string `json:"signature"`
	GameLevel     *int   `json:"gameLevel"` // 使用指针来允许 null 值
}

func main() {
	// 示例1：设置 GameLevel 为 nil
	bodyWithNilGameLevel := Body{
		LuckyDrawType: 1,
		Timestamp:     1674015600,
		Signature:     "some_signature",
		GameLevel:     nil, // GameLevel 为 null
	}

	// 示例2：设置 GameLevel 为具体值
	gameLevelValue := 5
	bodyWithGameLevel := Body{
		LuckyDrawType: 2,
		Timestamp:     1674015601,
		Signature:     "another_signature",
		GameLevel:     &gameLevelValue, // 使用指针设置值
	}

	// 输出 JSON
	jsonData1, err := json.Marshal(bodyWithNilGameLevel)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
	}
	fmt.Println("Body with nil GameLevel:", string(jsonData1))

	jsonData2, err := json.Marshal(bodyWithGameLevel)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
	}
	fmt.Println("Body with GameLevel set:", string(jsonData2))
}
