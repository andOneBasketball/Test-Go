package main

import (
	"encoding/json"
	"fmt"
)

type TalentCostData struct {
	Cost int `json:"cost"`
}

type HeroTalentBasicConfig struct {
	Name string `json:"name"`
}

type HeroTalentConfig struct {
	UpgradeCostDataList   []TalentCostData        `json:"upgradeCostDataList"`
	HeroTalentBasicConfig []HeroTalentBasicConfig `json:"heroTalentBasicConfig"`
}

func main() {
	jsonData := `{
		"upgradeCostDataList": [{"cost": 100}, {"cost": 200}],
		"heroTalentBasicConfig": [{"name": "Strength"}, {"name": "Agility"}]
	}`

	var config HeroTalentConfig
	err := json.Unmarshal([]byte(jsonData), &config)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Parsed Struct: %+v\n", config)

	jsonData1 := `[{"cost": 100}, {"cost": 200}]`
	talentConfig := make([]TalentCostData, 0)
	err = json.Unmarshal([]byte(jsonData1), &talentConfig)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Parsed Slice: %+v\n", talentConfig)
}
