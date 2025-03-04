package main

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/tidwall/gjson"
)

type TextData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// JSON 数据
	jsonStr := `[{"id":1, "name": "wooden_box", "need_for_open":{"game_count":10}, "count_max":999, "value_for_sort": 90000},{"id":2, "name": "gold_box", "need_for_open":{"gold_key":10}, "count_max":999, "value_for_sort": 130000},{"id":3, "name": "silver_box", "need_for_open":{"silver_key":10}, "count_max":999, "value_for_sort": 110000}]`
	jsonStr1 := `[{"id":1,"type":2,"subType":0,"note":"鱼竿","enNameText":"FIshing Rod","grade":1,"mainAttributeLevel":"6#8   ","mainAttributeText":"Double the probability of collecting items.#Raise the upper limit of fishing weight.","attributeValue":"1-20#2-1000","ascensionItemId":2,"iconResource":"fishingRod_1","sceneResource":"rod_1","timeLimit":0,"mainAttributeDataList":[{"mainAttributeLevel":6,"mainAttributeText":"Double the probability of collecting items.","mainAttributeType":1,"attributeValue":20},{"mainAttributeLevel":8,"mainAttributeText":"Raise the upper limit of fishing weight.","mainAttributeType":2,"attributeValue":1000}]},{"id":2,"type":2,"subType":0,"note":"鱼竿","enNameText":"FIshing Rod","grade":2,"mainAttributeLevel":"6#8   ","mainAttributeText":"Double the probability of collecting items.#Raise the upper limit of fishing weight.","attributeValue":"1-20#2-1000","ascensionItemId":3,"iconResource":"fishingRod_1","sceneResource":"rod_1","timeLimit":0,"mainAttributeDataList":[{"mainAttributeLevel":6,"mainAttributeText":"Double the probability of collecting items.","mainAttributeType":1,"attributeValue":20},{"mainAttributeLevel":8,"mainAttributeText":"Raise the upper limit of fishing weight.","mainAttributeType":2,"attributeValue":1000}]},{"id":3,"type":2,"subType":0,"note":"鱼竿","enNameText":"FIshing Rod","grade":3,"mainAttributeLevel":"6#8   ","mainAttributeText":"Double the probability of collecting items.#Raise the upper limit of fishing weight.","attributeValue":"1-20#2-1000","ascensionItemId":4,"iconResource":"fishingRod_1","sceneResource":"rod_1","timeLimit":0,"mainAttributeDataList":[{"mainAttributeLevel":6,"mainAttributeText":"Double the probability of collecting items.","mainAttributeType":1,"attributeValue":20},{"mainAttributeLevel":8,"mainAttributeText":"Raise the upper limit of fishing weight.","mainAttributeType":2,"attributeValue":1000}]},{"id":4,"type":2,"subType":0,"note":"鱼竿","enNameText":"FIshing Rod","grade":4,"mainAttributeLevel":"6#8   ","mainAttributeText":"Double the probability of collecting items.#Raise the upper limit of fishing weight.","attributeValue":"1-20#2-1000","ascensionItemId":5,"iconResource":"fishingRod_1","sceneResource":"rod_1","timeLimit":0,"mainAttributeDataList":[{"mainAttributeLevel":6,"mainAttributeText":"Double the probability of collecting items.","mainAttributeType":1,"attributeValue":20},{"mainAttributeLevel":8,"mainAttributeText":"Raise the upper limit of fishing weight.","mainAttributeType":2,"attributeValue":1000}]},{"id":5,"type":2,"subType":0,"note":"鱼竿","enNameText":"FIshing Rod","grade":5,"mainAttributeLevel":"6#8   ","mainAttributeText":"Double the probability of collecting items.#Raise the upper limit of fishing weight.","attributeValue":"1-20#2-1000","ascensionItemId":6,"iconResource":"fishingRod_1","sceneResource":"rod_1","timeLimit":0,"mainAttributeDataList":[{"mainAttributeLevel":6,"mainAttributeText":"Double the probability of collecting items.","mainAttributeType":1,"attributeValue":20},{"mainAttributeLevel":8,"mainAttributeText":"Raise the upper limit of fishing weight.","mainAttributeType":2,"attributeValue":1000}]},{"id":6,"type":2,"subType":0,"note":"鱼竿","enNameText":"FIshing Rod","grade":6,"mainAttributeLevel":"6#8   ","mainAttributeText":"Double the probability of collecting items.#Raise the upper limit of fishing weight.","attributeValue":"1-20#2-1000","ascensionItemId":0,"iconResource":"fishingRod_1","sceneResource":"rod_1","timeLimit":0,"mainAttributeDataList":[{"mainAttributeLevel":6,"mainAttributeText":"Double the probability of collecting items.","mainAttributeType":1,"attributeValue":20},{"mainAttributeLevel":8,"mainAttributeText":"Raise the upper limit of fishing weight.","mainAttributeType":2,"attributeValue":1000}]}]`

	// 检查每个 JSON 是否包含指定键
	results := gjson.Parse(jsonStr).Array()
	log.Println(reflect.TypeOf(results))
	for _, obj := range results {
		//log.Printf("jsonStr, %dth obj", i)
		// 检查顶层的键
		if obj.Get("enNameText").Exists() {
			value := obj.Get("enNameText")
			log.Println(value.String(), reflect.TypeOf(value.String()))
		} else if obj.Get("name").Exists() {
			value := obj.Get("name")
			log.Println(value.String(), reflect.TypeOf(value.String()))
		}

		if obj.Get("need_for_open.game_count").Exists() {
			value := obj.Get("need_for_open.game_count")
			log.Println(value.Int(), reflect.TypeOf(value.Int()))
			log.Println(value.String(), reflect.TypeOf(value.String()))
		}
	}

	results = gjson.Parse(jsonStr1).Array()
	for _, obj := range results {
		//log.Printf("obj value:%#v", obj.Value())

		s := struct {
			Grade              int    `json:"grade"`
			MainAttributeLevel string `json:"mainAttributeLevel"`
		}{
			Grade:              int(obj.Get("grade").Int()),
			MainAttributeLevel: obj.Get("mainAttributeLevel").String(),
		}
		log.Printf("struct grade: %d, mainAttributeLevel: %s", s.Grade, s.MainAttributeLevel)

		m := obj.Value().(map[string]interface{})
		log.Printf("map grade: %d, mainAttributeLevel: %s", int(m["grade"].(float64)), m["mainAttributeLevel"])
		//log.Printf("jsonStr1, %dth obj", i)
		// 检查顶层的键
		if obj.Get("enNameText").Exists() {
			value := obj.Get("enNameText")
			log.Println("enNameText: ", value, reflect.TypeOf(value))
		}
	}

	t := make([]TextData, 2)
	json.Unmarshal([]byte(jsonStr), &t)
	for _, obj := range t {
		log.Printf("id: %d, name: %s", obj.Id, obj.Name)
	}

	value := ""
	var configDataValue interface{}
	err := json.Unmarshal([]byte(value), &configDataValue)
	if err != nil {
		log.Printf("json Unmarshal error: %s", err)
	}
	log.Printf("configDataValue: %v", configDataValue)
}
