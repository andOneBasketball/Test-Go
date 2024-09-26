package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	Id     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "test:test123@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//创建表 自动迁移 （把结构体和数据库表进行对应）
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	u1 := UserInfo{Id: 1, Name: "zyj", Gender: "男", Hobby: "唱"}
	db.Create(&u1)
	//查询
	var u UserInfo
	db.First(&u)
	fmt.Println(u)
	//更新
	db.Model(&u).Update("hobby", "唱跳rap篮球")
	//删除
	db.Delete(&u) //将查询出来的第一条数据删除
}

