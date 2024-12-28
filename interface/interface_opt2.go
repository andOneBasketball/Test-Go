package main

import (
	"log"
)

func main() {
	var i interface{} = 10
	log.Printf("i is of type %T and value %v", i, i)
	i = "hello"
	log.Printf("i is of type %T and value %v", i, i)
	j, ok := i.(int)
	log.Printf("j is of type %T and value %v, ok is %v", j, j, ok)

	i = 20.5
	log.Printf("i is of type %T and value %v", i, i)
	i = map[string]interface{}{"key": 10}
	log.Printf("i is of type %T and value %v", i, i)
}
