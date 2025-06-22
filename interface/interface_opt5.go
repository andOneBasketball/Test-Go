package main

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
	s := S{}
	p := &s
	var m interface{}
	f(s) //A
	g(s) //B
	f(p) //C
	g(p) //D
	f(m)
	g(&m)
}
