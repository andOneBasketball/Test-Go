package main

import (
	"encoding/json"
	"fmt"
)

// 结构体定义
type MultiRewardPropUserResource struct {
	Count        int64 `json:"count"`
	TurnOn       bool  `json:"turn_on"`
	TicketStatus bool  `json:"ticket_status"` // 修正反引号错误
}

func main() {
	// 原始 JSON 字符串
	jsonStr := `{"count":1,"turn_on":false,"market_box_id":0,"mpb_id":51,"mnpb_id":50}`

	// 解析目标结构体
	var data MultiRewardPropUserResource

	// JSON 解析
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("JSON 解析失败:", err)
		return
	}

	// 输出解析结果
	fmt.Printf("Count: %d, TurnOn: %t, TicketStatus: %t\n", data.Count, data.TurnOn, data.TicketStatus)

	sss := false
	fmt.Printf("sss: %t\n", sss)
}
