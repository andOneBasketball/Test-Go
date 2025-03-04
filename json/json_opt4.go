package main

import (
	"fmt"

	"github.com/tidwall/sjson"
)

type EnergyInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// 初始 JSON 字符串
	jsonData := `{"name": "Alice", "skills": []}`

	// 设置 JSON 数组中的值
	updatedJSON, err := sjson.Set(jsonData, "skills.0", "JavaScript")
	if err != nil {
		fmt.Println("Error setting JSON value:", err)
		return
	}

	// 添加更多数组项
	updatedJSON, _ = sjson.Set(updatedJSON, "skills.1", "Python")
	updatedJSON, _ = sjson.Set(updatedJSON, "skills.2", "Solidity")

	// 打印结果
	fmt.Println(updatedJSON)

	updatedJSON, _ = sjson.Delete(updatedJSON, "skills.1")
	fmt.Println(updatedJSON)

	energyInfos := []EnergyInfo{
		{1, "Solar Energy"},
		{2, "Wind Energy"},
		{3, "Hydro Energy"},
	}
	updatedJSON, _ = sjson.Set(updatedJSON, "energy_info", energyInfos)
	fmt.Println(updatedJSON)

	updatedJSON, err = sjson.Delete(updatedJSON, "skills.100")
	if err != nil {
		fmt.Println("Error deleting JSON value:", err)
	}
	fmt.Println(updatedJSON)
}
