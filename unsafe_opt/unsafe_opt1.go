// main.go
package main

import (
	"fmt"
	"unsafe"
	"unsafe_opt/mypackage"
)

func main() {
	p := mypackage.Person{Age: 25}
	//p.name = "Bob" // ❌ 编译错误：cannot refer to unexported field 'name'

	// 获取结构体指针
	ptr := unsafe.Pointer(&p)

	// 计算 name 字段的偏移量（需要知道内存布局）
	// 假设 name 是第一个字段，偏移量是 0
	namePtr := (*string)(unsafe.Pointer(uintptr(ptr) + 0)) // 实际要用 unsafe.Offsetof

	heightPtr := (*int)(unsafe.Pointer(uintptr(ptr) + 24)) // 手动计算偏移量，极其脆弱

	// 修改私有字段
	*namePtr = "Bob"

	*heightPtr = 20

	fmt.Println(p)
}
