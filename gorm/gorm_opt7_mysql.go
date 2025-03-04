package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:a1b3c3d4@tcp(172.29.164.230:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetConnMaxIdleTime(59000 * time.Millisecond)
	sqlDB.SetConnMaxLifetime(59000 * time.Millisecond)

	wg := &sync.WaitGroup{}
	// Simulate frequent queries in a loop to observe behavior
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				db.Exec("SELECT SLEEP(0.1)") // Simulate a query
				stats := sqlDB.Stats()
				fmt.Printf("Open: %d, In Use: %d, Idle: %d, WaitCount: %d\n",
					stats.OpenConnections, stats.InUse, stats.Idle, stats.WaitCount)
				time.Sleep(1 * time.Second) // Add delay to mimic realistic load
			}
		}()
	}

	wg.Wait()
}
