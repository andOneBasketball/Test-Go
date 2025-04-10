package main

import (
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	// 示例 JSON 数据
	jsonData := `{
		"name": "John Doe",
		"age": 30,
		"address": {
			"street": 123,
			"city": "New York"
		},
		"phones": [
			{"type": "home", "number": "123-456-7890"},
			{"type": "work", "number": "987-654-3210"}
		]
	}`

	// 使用 gjson 解析 JSON 数据
	j := gjson.New(jsonData)
	name := j.Get("name")
	age := j.Get("age")
	street := j.Get("address.street").Uint()
	city := j.Get("address.city")
	phoneNumbers := j.Get("phones.#.number")

	// 打印解析结果
	g.Log().Info(ctx, "Name:", name.String())
	g.Log().Line().Info(ctx, "Age:", age.Int())
	g.Log().Error(ctx, "Street:", street)
	g.Log().Warning(ctx, "City:", city.String())
	g.Log().Fatal(ctx, "Phone Numbers:", phoneNumbers.Array())

	decodeState, _ := gurl.Decode(`{"data_id":63,"app_id":"5906751"}`)
	fmt.Println(decodeState)
	jState := gjson.New(decodeState)
	dataId := jState.Get("data_id").Uint()
	appId := jState.Get("app_id").String()
	fmt.Println("data_id:", dataId)
	fmt.Println("app_id:", appId)
}
