package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Result struct {
	Count int `bson:"count"`
}

type ProcessRuleDocument struct {
	Id         bson.ObjectId `bson:"_id"`
	AInfo    interface{}        `bson:"a_info"`
	BType   uint32        `bson:"b_type"`
}


func main() {
	host := "x.x.x.41:27017"
	mongo, _ := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{host},
		Username: "root",
		Password: "xxxxxx",
	})

	collection := mongo.DB("db_hello").C("tb_hello")

	iter := collection.Find(bson.M{"product": "A", "attr": 101}).Batch(10).Iter()
	var doc ProcessRuleDocument
	for iter.Next(&doc) {
		fmt.Printf("%#v\n", doc)
		//fmt.Printf("origin cmd: %s, after: %s\n", doc.CmdInfo, strings.Replace(doc.CmdInfo, "\\", "\\\\\\", -1))
	}
}
