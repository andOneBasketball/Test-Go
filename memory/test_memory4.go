package main

//编译阶段不能确定接口的动态类型（interface{}）

type Inter1 interface {
	A()
}

type Inter2 interface {
	B()
}

type User struct {
	Name string
	Id   int
}

func (u *User) A() {
}

func (u User) B() {
}

func Test1() interface{} {
	var a interface{}
	a = &User{}
	return a
}

func Test2() Inter1 {
	var b Inter1
	b = &User{}
	return b
}

func Test3() Inter2 {
	var b Inter2
	b = User{}
	return b
}

func main() {
	Test1()
	Test2()
	Test3()
}
