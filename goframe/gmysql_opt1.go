package main

import (
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()

	// 假设已初始化数据库
	db, _ := gdb.New(gdb.ConfigNode{
		Type: "mysql",
		Link: "root:password@tcp(127.0.0.1:3306)/test",
	})

	// 批量数据，模拟错误数据（假设 email 字段是 UNIQUE，重复值会导致错误）
	saveData := []g.Map{
		{"name": "user_1", "email": "user1@example.com"},
		{"name": "user_2", "email": "user2@example.com"},
		{"name": "user_3", "email": "user1@example.com"}, // ❌ 这条数据会触发唯一索引冲突
		{"name": "user_4", "email": "user4@example.com"},
	}

	// 批量插入（每 2 条为一批）
	_, err := db.Model("users").Ctx(ctx).Data(saveData).Batch(2).Insert()
	if err != nil {
		fmt.Println("批量插入失败:", err)
	} else {
		fmt.Println("批量插入成功")
	}
}
