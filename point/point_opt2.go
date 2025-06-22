package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var validVar int = 42
	validPtr := unsafe.Pointer(&validVar)

	testCases := []struct {
		name string
		ptr  unsafe.Pointer
	}{
		{"Valid Pointer", validPtr},
		{"Nil Pointer", nil},
		{"Zero Pointer", unsafe.Pointer(uintptr(0))},
		{"One Pointer", unsafe.Pointer(uintptr(1))},
	}

	for _, c := range testCases {
		fmt.Printf("%s: %v\n", c.name, isPointerValid(c.ptr))
	}
}

func isPointerValid(p unsafe.Pointer) bool {
	// 基本检查
	if p == nil {
		return false
	}

	// 零地址检查
	if uintptr(p) == 0 {
		return false
	}

	// 反射检查
	v := reflect.ValueOf(p)
	if !v.IsValid() || v.IsNil() {
		return false
	}

	// 尝试访问检查（安全方式）
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	_ = *(*byte)(p)

	return true
}
