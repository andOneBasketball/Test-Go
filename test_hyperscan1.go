package main

import (
	"fmt"
	"github.com/flier/gohs/hyperscan"
)

func main() {
	// 创建一个模式数据库
	db, err := hyperscan.NewPatternDB([]string{"foo", "bar"}, []uint32{0, 1})
	if err != nil {
		fmt.Println("Error creating pattern database:", err)
		return
	}

	// 创建一个扫描器
	s, err := hyperscan.NewScanner(db)
	if err != nil {
		fmt.Println("Error creating scanner:", err)
		return
	}

	// 要扫描的输入
	input := "foobaz"

	// 执行扫描并指定回调函数
	err = s.Scan([]byte(input), func(id uint32, from, to uint64, flags uint, context interface{}) error {
		fmt.Printf("Matched pattern ID: %d, from: %d, to: %d\n", id, from, to)
		return nil
	}, nil)

	if err != nil {
		fmt.Println("Error scanning input:", err)
		return
	}
}