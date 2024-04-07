package main

import (
	"flag"
	"fmt"
)

type s1 struct {
	str1 string
	i1   int
}

func main() {
	var appConf string
	flag.StringVar(&appConf, "config", "app/config.json", "app config")
	flag.Parse()
	fmt.Printf("%s\n", appConf)

	var ss *s1
	ss = &s1{
		str1: "hello",
		i1: 10,
	}
	fmt.Printf("%#v\n", ss)
}
