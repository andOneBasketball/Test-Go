package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)


func Map(d bson.D) bson.M {
	m := make(bson.M, len(d))
	for _, e := range d {
		m[e.Name] = e.Value
	}
	return m
}


func main() {
	setValue := bson.M{"online_time": 0, "config_version": "",
		"update_time": 0, "dev_version": "", "extend": ""}
	d := bson.M{}
	value := bson.D{bson.DocElem{Name: "online_time", Value: 0}}
	fmt.Println(setValue)
	fmt.Println(value)
	fmt.Println(len(setValue), len(d))
}