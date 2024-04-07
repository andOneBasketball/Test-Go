package main

import (
	"fmt"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	//"reflect"
	//"time"
)

// 样本管理集合
type SampleResultDocument struct {
	Id                 bson.ObjectId `bson:"_id" json:"objectId"`
	Md5                string        `bson:"md5" json:"md5"`
	CreateTime         uint32        `bson:"create_time" json:"createTime"`
}

type statistics struct {
	Md5   string `bson:"_id"`
	Count uint64 `bson:"count"`
}

func main() {
	var host = "x.x.x.x:27017"
	var mongo, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{host},
		Username: "root",
		Password: "xxxxx",
	})
	if err != nil {
		fmt.Println("connect error: ", err)
		return
	}
	
	/*
	setValue := bson.M{"command_feedback.rule_update": "200, success"}
	//setValue := bson.M{"operator": "200, success"}
	value := bson.M{"$set": setValue}
	client := mongo.DB("db_hello_check").C("tb_hello")
	//docID, _ := primitive.ObjectIDFromHex("60460c94637182e982aa642a")
	res, errs := client.Upsert(bson.M{"update_time": 1666343687, "create_time": 1615203483}, value)
	fmt.Println(res, errs)
	*/
	

	/*
	client := mongo.DB("db_hello_check").C("tb_hello")

	var sort_filed = "md5"
	result := new(statistics)
	fmt.Println(result.Count)
	var pipe []bson.M
	group := bson.M{
		"_id": "$" + sort_filed,
		"count": bson.M{
			"$sum": 1,
		},
	}
	pipe = append(pipe, bson.M{"$group": group})
	pipe = append(pipe, bson.M{"$sort": bson.M{"count": -1}})
	pipe = append(pipe, bson.M{"$limit":10})
	//client.Pipe(pipe).All(&result)
	//fmt.Println(result)
	iter := client.Pipe(pipe).Iter()
	count := 0
	for iter.Next(&result) {
		fmt.Println(result.Md5)
		count++
	}
	fmt.Println(count)
	*/
	

	//fmt.Println(result)

	/*
		// Ensure cursor logic is working by forcing a small batch.
		pipe.Batch(2)

		// Smoke test for AllowDiskUse.
		pipe.AllowDiskUse()

		iter := pipe.Iter()
		result := struct{ N int }{}
		for i := 2; i < 7; i++ {
			ok := iter.Next(&result)
			c.Assert(ok, Equals, true)
			c.Assert(result.N, Equals, ns[i])
		}

		c.Assert(iter.Next(&result), Equals, false)
		c.Assert(iter.Close(), IsNil)
	*/


	sampleResult := SampleResultDocument{}
	client := mongo.DB("db_hello_check").C("tb_hello")
	/*
	sort := []string{"create_time"}
	fmt.Printf("%v\n", sort...)
	iter := client.Find(bson.M{}).Sort([]string{"create_time"}...).Batch(10).Iter()
	var doc SampleResultDocument
	for iter.Next(&doc) {
		fmt.Printf("%v\n", doc)
		time.Sleep(1000 * time.Millisecond)
	}
	*/

	/*
	docID, _ := primitive.ObjectIDFromHex("5fd869dd02287580463a8845")
	cErr := client.Find(bson.M{"_id": docID}).One(&sampleResult)
	if cErr != nil {
		fmt.Printf("find error: %v", cErr)
		return
	}
	fmt.Println(sampleResult, "\n", sampleResult.HidsTrust)
	*/

	
	cErr := client.Find(bson.M{"md5": "bf619eac0cdf3111d496ea9344137e8b"}).One(&sampleResult)
	if cErr != nil {
		fmt.Println("find error.")
		return
	}
	if sampleResult.Md5 == "" {
		fmt.Println("1234")
	}
	fmt.Printf("sampleResult:%#v md5:%v action:%d\n", sampleResult, sampleResult.Md5, int32(sampleResult.CreateTime))
	if sampleResult.Md5 == "" {
		fmt.Println("123445654")
	}
}
