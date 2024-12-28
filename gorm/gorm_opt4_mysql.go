package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 配置你的数据库连接
	dsn := "root:ZxcvAsdf@tcp(10.233.59.153:3306)/weplay-pro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 查询当前连接数
	var threadsConnected int
	row := db.Raw("SHOW STATUS LIKE 'Threads_connected';").Row()
	err = row.Scan(&threadsConnected, &threadsConnected) // 忽略第一列的名称
	if err != nil {
		panic(err)
	}

	fmt.Printf("当前连接数: %d\n", threadsConnected)
}
