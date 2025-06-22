package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   *int
}

func main() {
	// 创建 User 结构体实例
	user := &User{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
	}

	// 打印指针地址和结构体内容
	fmt.Printf("Pointer address: %p\n", user) // 打印指针地址
	fmt.Printf("User struct: %+v\n", user)    // 打印结构体内容
	user.Name = "Bob"
	fmt.Printf("User struct: %#v\n", *user)
	users := []*User{user}
	fmt.Printf("Users struct: %+v\n", users)

	// 格式化为 JSON 字符串输出
	userJSON, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("User as JSON:")
	fmt.Println(string(userJSON))

	var aaa *int
	aaa = nil
	user.Age = aaa
	bb, _ := json.Marshal(user)
	log.Printf("%v", string(bb))
	b := 10
	aaa = &b
	user.Age = aaa
	bb, _ = json.Marshal(user)
	log.Printf("%v", string(bb))
	if aaa != nil && *aaa == 0 {
		fmt.Println("aaa is nil")
	}
}
