package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func main() {
	//定义map
	m := make(map[string]*student)

	//定义student数组
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	//将数组依次添加到map中
	// 1.22 开始会为每次循环创建块级作用域
	for _, stu := range stus {
		m[stu.Name] = &stu
		//break
	}

	//打印map
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
}
