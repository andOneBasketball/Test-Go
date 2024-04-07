package main

import "fmt"

func modifyMap(m map[int]map[int]string) {
    m[0][0] = "modified"
}

func main() {
    // 创建一个二维哈希 map
    m := make(map[int]map[int]string)
    m[0] = make(map[int]string)
    m[0][0] = "original"

    // 调用函数修改 map
    modifyMap(m)

    // 输出修改后的值
    fmt.Println(m[0][0]) // 输出：modified

    a := map[string]map[string]int{
        "abcd": {
            "qewf": 1,
            "wef": 2,
        },
    }
    fmt.Printf("%#v\n", a)
}