package main

import (
	"fmt"
	"sort"
)

// 自定义结构体
type Person struct {
	Name string
	Age  int
}

// 实现 sort.Interface 接口（Len, Less, Swap）
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	sort.Sort(ByAge(people)) // 使用 sort.Sort 排序
	fmt.Println("按年龄排序后:")
	for _, p := range people {
		fmt.Printf("%s (%d)\n", p.Name, p.Age)
	}
}
