package main

import (
	"fmt"
	"sort"
)

func less(a, b [2]int) bool {
	if a[0] < b[0] {
		return true
	}
	if a[0] == b[0] {
		return a[1] < b[1]
	}
	return false
}

func main() {
	// 整数排序
	ints := []int{5, 2, 8, 3, 1}
	sort.Ints(ints)
	fmt.Println("整数排序:", ints)

	int2s := [][2]int{[2]int{1, 2}, [2]int{3, 5}, [2]int{2, 5}}
	sort.Slice(int2s, func(i, j int) bool {
		return less(int2s[i], int2s[j])
	})
	fmt.Println(int2s)

	// 小数排序（float64）
	floats := []float64{3.14, 1.41, 2.71, 0.99}
	sort.Float64s(floats)
	fmt.Println("小数排序:", floats)

	// 字符串排序
	strs := []string{"banana", "apple", "grape", "kiwi"}
	sort.Strings(strs)
	fmt.Println("字符串排序:", strs)

	// 自定义排序：字符串按长度排序
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	fmt.Println("按字符串长度排序:", strs)
}
