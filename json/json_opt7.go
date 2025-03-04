package main

import (
	"encoding/json"
	"log"
)

func main() {
	type body struct {
		PointCouponType  int `json:"point_coupon_type"`   // 券ID
		LuckDrawConfigID int `json:"luck_draw_config_id"` // 箱子ID
		LuckDrawWay      int `json:"luck_draw_way"`       // 开箱方式, 1: 使用券；2: 使用箱子
	}

	jsonData, _ := json.Marshal(body{
		LuckDrawConfigID: 1,
		LuckDrawWay:      2,
	})
	log.Printf("jsonData: %s", string(jsonData))
}
