package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. 计算 n 次方
	base := 2.0
	exponent := 3.0
	power := math.Pow(base, exponent)
	fmt.Printf("%.2f 的 %.2f 次方 = %.2f\n", base, exponent, power) // 2^3 = 8

	// 2. 计算 n 次根号（通过 math.Pow 实现）
	num := 27.0
	n := 3.0 // 开 3 次根号
	root := math.Pow(num, 1/n)
	fmt.Printf("%.2f 的 %d 次根号 = %.2f\n", num, n, root) // ³√27 = 3

	// 3. 测试 math 包的其他常用方法
	// (1) 绝对值
	absVal := math.Abs(-5.5)
	fmt.Printf("Abs(-5.5) = %.2f\n", absVal) // 5.5

	// (2) 平方根
	sqrtVal := math.Sqrt(16)
	fmt.Printf("Sqrt(16) = %.2f\n", sqrtVal) // 4

	// (3) 立方根
	cbrtVal := math.Cbrt(64)
	fmt.Printf("Cbrt(64) = %.2f\n", cbrtVal) // 4

	// (4) 最大值 & 最小值
	maxVal := int(math.Max(3, 7))
	minVal := math.Min(3, 7)
	fmt.Printf("Max(3, 7) = %d, Min(3, 7) = %d\n", maxVal, minVal) // 7, 3

	// (5) 三角函数（弧度制）
	sinVal := math.Sin(math.Pi / 2)                                // π/2 ≈ 1.5708 → sin(π/2) = 1
	cosVal := math.Cos(math.Pi)                                    // π ≈ 3.1416 → cos(π) = -1
	fmt.Printf("Sin(π/2) = %.2f, Cos(π) = %.2f\n", sinVal, cosVal) // 1, -1

	// (6) 对数
	logVal := math.Log(math.E)                                                             // ln(e) = 1
	log10Val := math.Log10(100)                                                            // log10(100) = 2
	log2Val := math.Log2(8)                                                                // log2(8) = 3
	fmt.Printf("Ln(e) = %.2f, Log10(100) = %d, Log2(8) = %d\n", logVal, log10Val, log2Val) // 1, 2, 3

	// (7) 取整 & 舍入
	ceilVal := math.Ceil(3.2)   // 向上取整 → 4
	floorVal := math.Floor(3.8) // 向下取整 → 3
	roundVal := math.Round(3.5) // 四舍五入 → 4
	truncVal := math.Trunc(3.9) // 截断小数 → 3
	fmt.Printf("Ceil(3.2) = %d, Floor(3.8) = %d, Round(3.5) = %d, Trunc(3.9) = %d\n",
		int(ceilVal), int(floorVal), int(roundVal), int(truncVal)) // 4, 3, 4, 3

	// (8) 常量
	fmt.Printf("Pi = %.5f, E = %.5f\n", math.Pi, math.E) // π ≈ 3.14159, e ≈ 2.71828
}
