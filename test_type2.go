import (
    "reflect"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Alice", Age: 20}

    // 判断结构体中是否存在 Name 字段
    _, ok := reflect.TypeOf(p).FieldByName("Name")
    if ok {
        fmt.Println("Name exists")
    } else {
        fmt.Println("Name does not exist")
    }

    // 判断结构体中是否存在 Address 字段
    _, ok = reflect.TypeOf(p).FieldByName("Address")
    if ok {
        fmt.Println("Address exists")
    } else {
        fmt.Println("Address does not exist")
    }
}
