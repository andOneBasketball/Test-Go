package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)


func main() {
	pipeline := [...]bson.M{
		bson.M{"aa": "bb"},
		bson.M{"dd": 1},
	}
	data, _ := bson.Marshal(pipeline)
	fmt.Printf("%#v\n", data)
}

