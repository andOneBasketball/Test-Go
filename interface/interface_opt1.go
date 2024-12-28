package main

import (
	"fmt"
	"log"
)

type MyStruct struct {
	Field int
}

func main() {
	var i interface{} = MyStruct{Field: 12}

	if ms, ok := i.(MyStruct); ok {
		// 使用 %p 显示内存地址
		fmt.Printf("Address of ms: %p\n", &ms)
		fmt.Printf("Address of i: %p\n", &i)
		fmt.Println(ms.Field) // 输出: 12

		// 尝试修改 ms 的 Field 值
		ms.Field = 123

		// 无法修改，依然输出12
		if ms2, ok := i.(MyStruct); ok {
			fmt.Println(ms2.Field) // 输出: 12
		}
	}

	// 如果我们想要通过接口修改原始值，需要确保接口持有的是引用类型（如指针）
	var ip interface{} = &MyStruct{Field: 12}
	if _, ok := ip.(int); !ok {
		log.Println("ip is not a int data type")
	}
	if _, ok := ip.(MyStruct); !ok {
		log.Println("ip is not a MyStruct data type")
	}

	if msp, ok := ip.(*MyStruct); ok {
		// msp 和 ip 指向同一个地址
		fmt.Printf("Address of ms: %p\n", msp)
		fmt.Printf("Address of ip: %p\n", ip)
		log.Println(msp.Field)
		msp.Field = 1234567
		log.Println(msp.Field)
		log.Printf("%#v", ip)
	}

	// 现在通过断言再次获取值，验证修改
	if ms, ok := ip.(*MyStruct); ok {
		fmt.Println(ms.Field) // 输出: 1234567
	}
}
