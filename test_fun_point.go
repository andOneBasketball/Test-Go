package main

import "fmt"

// 定义一个函数类型
type operation func(int, int) int

// 定义一个函数，接受一个函数类型的参数
func calculate(x int, y int, op operation) int {
    return op(x, y)
}

// 定义一个加法函数
func add(x int, y int) int {
    return x + y
}

// 定义一个减法函数
func subtract(x int, y int) int {
    return x - y
}

func main() {
    // 定义一个函数指针，指向加法函数
    var op operation = add

    // 使用函数指针调用加法函数
    result := op(3, 4)
    fmt.Println(result) // 输出 7

    // 将函数指针指向减法函数
    op = subtract

    // 使用函数指针调用减法函数
    result = op(3, 4)
    fmt.Println(result) // 输出 -1

    // 将函数指针作为参数传递给另一个函数
    result = calculate(3, 4, add)
    fmt.Println(result) // 输出 7

    result = calculate(3, 4, subtract)
    fmt.Println(result) // 输出 -1
}