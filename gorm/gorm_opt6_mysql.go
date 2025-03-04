package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:a1b3c3d4@tcp(172.29.164.230:3306)/weplay-pro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 获取底层 *sql.DB
	sqlDB, _ := db.DB()
	printDBStats("Before WithContext", sqlDB)
	ctxG, _ := context.WithTimeout(context.Background(), time.Second*30)
	dbWithctxG := db.WithContext(ctxG)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
			dbWithCtx := db.WithContext(ctx)

			// 验证 dbWithCtx 是否创建了新的连接
			printDBStats(fmt.Sprintf("After %d WithContext, sqlDB address: %p", i, dbWithCtx), sqlDB)

			// 查询示例（实际代码视具体表而定）
			var result []map[string]interface{}
			dbWithCtx.Table("user_attribute").Find(&result)

			// 再次打印状态
			printDBStats(fmt.Sprintf("After %d Query with Context, sqlDB address: %p", i, dbWithCtx), sqlDB)
		}()

		go func() {
			defer wg.Done()
			var result []map[string]interface{}
			dbWithctxG.Table("user_attribute").Find(&result)
			printDBStats(fmt.Sprintf("Second after %d Query with Context, sqlDB address: %p", i, dbWithctxG), sqlDB)
		}()
	}
	wg.Wait()
}

func printDBStats(label string, db *sql.DB) {
	stats := db.Stats()
	fmt.Printf("[%s] Open: %d, In Use: %d, Idle: %d\n", label, stats.OpenConnections, stats.InUse, stats.Idle)
}
