package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	// 假设给定的int64类型数据为时间戳，表示2023年10月19日 12:00:00
	timestamp := int64(1697990400)

	// 将int64类型数据转换为time.Time类型
	t := time.Unix(timestamp, 0)

	duration := int64(time.Now().Sub(t).Seconds())

	var fileLastModifyTimeLimit, a int64
	fileLastModifyTimeLimit = 10
	a = 5
	a |= 3
	f := fmt.Sprintf("%d", a)

	fmt.Println("转换结果：", t, duration, reflect.TypeOf(duration), fileLastModifyTimeLimit, a, f)
}