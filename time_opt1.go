package main

import (
	"fmt"
	"time"
)

type UserEventLog struct {
	CreateTime time.Time `gorm:"type:datetime(6)"`
}

func main() {
	// 模拟从数据库读取的时间值
	log := UserEventLog{
		CreateTime: time.Now(), // 假设当前时间是数据库查询结果
	}

	// 显式格式化为带微秒的字符串
	formattedTime := log.CreateTime.Format("2006-01-02 15:04:05.000000")
	fmt.Println("Formatted CreateTime:", formattedTime)
}
