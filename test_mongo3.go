package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 样本管理集合  tb_ids_md5
type Result struct {
	Count int `bson:"count"`
}


func main() {
	host := "x.x.x.141:27017"
	mongo, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{host},
		Username: "root",
		Password: "xxxx",
	})

	collection := mongo.DB("db_hello").C("tb_hello")

	matchStage := bson.M{
		"$match": bson.M{
			"trust":        bson.M{"$in": []interface{}{"2", "", nil}},
		},
	}

	groupStage := bson.M{
		"$group": bson.M{
			"_id": "$md5",
		},
	}

	groupStage2 := bson.M{
		"$group": bson.M{
			"_id":   nil,
			"count": bson.M{"$sum": 1},
		},
	}

	projectStage := bson.M{
		"$project": bson.M{
			"_id":   0,
			"count": 1,
		},
	}

	pipe := collection.Pipe([]bson.M{matchStage, groupStage, groupStage2, projectStage})

	var result Result
	err = pipe.One(&result)
	if err != nil {
		fmt.Println("Failed to execute aggregation:", err)
		return
	}

	fmt.Println("Count:", result.Count)
}
