package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// 替换为你的数据库连接信息
	dsn := "root:ZxcvAsdf@tcp(10.233.59.153:3306)/weplay-pro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// 初始化生成器
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./dao",   // 生成代码的输出路径
		ModelPkgPath: "./model", // 模型包路径
	})

	g.UseDB(db) // 使用数据库实例

	// 指定表名生成模型代码
	g.ApplyBasic(g.GenerateModel("user_event_log"))

	// 执行生成
	g.Execute()
}
