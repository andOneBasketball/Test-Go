package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserEventLog struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	EventType  int       `gorm:"not null" json:"event_type"`
	UserID     int64     `gorm:"not null" json:"user_id"`
	ValueNum   float64   `json:"value_num"`
	CreateTime time.Time `gorm:"default:current_timestamp" json:"create_time"`
}

func main() {
	dsn := "root:a1b3c3d4@tcp(172.29.164.230:3306)/weplay-pro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Sample data for batch insert
	logs := []UserEventLog{
		{EventType: 1, UserID: 1001, ValueNum: 10.5},
		{EventType: 2, UserID: 1002, ValueNum: 20.0},
		{EventType: 3, UserID: 1003, ValueNum: 5.5},
	}

	// Insert data in batch
	result := db.Create(&logs)
	if result.Error != nil {
		log.Println("Failed to insert logs:", result.Error)
	} else {
		fmt.Println("Batch insert successful")
		log.Printf("RowsAffected = %d", result.RowsAffected)

	}
}
