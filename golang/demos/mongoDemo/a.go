package mongoDemo

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type JPerson struct {
	Name     string
	Phone    string
	DateTime int64 `json:"DateTime"`
	Id       int   `bson:"_id"`
}

const (
	// ADDRESSPORT = "192.168.10.35:3000"
	// DB          = "ssp_dev"
	// COLLECTION  = "demo"
	ADDRESSPORT = "127.0.0.8:27017,127.0.0.1:27017"
	DB          = "demo"
	COLLECTION  = "acud"
)

func SuntengFindOperationDemo() (err error) {
	println("//<<-------------------------OrOperationDemo start-----------")
	session, err := mgo.Dial(ADDRESSPORT)
	if err != nil {
		return
	}
	defer session.Close()
	// session.SetMode(mgo.Monotonic, true)
	c := session.DB(DB).C(COLLECTION)
	// for i := 0; i < 1; i++ {
	// 	err = c.Insert(&JPerson{Name: "1json",
	// 		Phone:    "sort demo",
	// 		DateTime: time.Now().Unix(), Id: i})
	// }
	// var result []interface{}
	var result []*JPerson
	// limit, offset := 0, 0
	// sort := "Id"
	// doQeury := bson.M{"IsDeleted": false}
	// doQeury := bson.M{"Name": "1json"}

	err = c.Find(nil).All(&result)
	// err = c.Find(doQeury).All(&result)
	// err = c.Find(nil).Sort("Id").All(&result)  not work
	if err != nil {
		fmt.Println("not found")
		return
	}
	fmt.Println("found")
	fmt.Printf("len(result): %+v\n", len(result))
	for i := 0; i < len(result); i++ {
		//@toDelete
		fmt.Printf("  result[i]: %+v\n", result[i])
	}

	println("//---------------------------OrOperationDemo end----------->>")
	return
}

func SuntengUpdateOperationDemo() (err error) {
	println("//<<-------------------------update operation start-----------")
	// session, err := mgo.Dial("192.168.10.35:3000")
	session, err := mgo.Dial(ADDRESSPORT)
	if err != nil {
		return
	}
	defer session.Close()
	// session.SetMode(mgo.Monotonic, true)
	c := session.DB(DB).C(COLLECTION)

	// var result []*JPerson
	// err = c.Find(nil).All(&result)
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }
	// for k, v := range result {
	// 	fmt.Printf("%+v: %+v\n", k, v)
	// 	err = c.UpdateId(v.Id, bson.M{"$set": bson.M{"ChildStatus.A": 1}})
	// 	if err != nil {
	// 		panic(err.Error())
	// 		return
	// 	}
	// }

	change, err := c.UpdateAll(nil, bson.M{"$set": bson.M{"Child.Add": 3}})
	//@toDelete
	fmt.Printf("  change: %+v\n", change)
	// for i := 0; i < 30; i++ {
	// 	err = c.Insert(&JPerson{"upgrade", "upgrade demo", i})
	// }
	// var result []interface{}
	// // limit, offset := 0, 0
	// // sort := "Id"
	// doQeury := bson.M{"IsDeleted": false}
	//
	// err = c.Find(doQeury).Sort("-_id").All(&result)
	// // err = c.Find(nil).Sort("Id").All(&result)  not work
	// if err != nil {
	// 	fmt.Println("not found")
	// 	return
	// }
	// fmt.Println("found")
	// for i := 0; i < len(result); i++ {
	// 	//@toDelete
	// 	fmt.Printf("  result[i]: %+v\n", result[i])
	// }

	println("//---------------------------update operation end----------->>")
	return
}

func AggregateQueryDemo() {
	println("//<<-------------------------AggregateQueryDemo start-----------")
	start := time.Now()

	session, err := mgo.Dial(ADDRESSPORT)
	if err != nil {
		return
	}
	defer session.Close()
	c := session.DB(DB).C(COLLECTION)

	project := bson.M{
		"$project": bson.M{
			"name": true, "a": true, "b": true, "suma": true, "sumb": true, "_id": true,
			"bdiva": bson.M{"$cond": []interface{}{bson.M{"$eq": []interface{}{"$suma", 0}}, 0, bson.M{"$divide": []interface{}{"$sumb", "$suma"}}}},
		},
	}

	group := bson.M{
		"$group": bson.M{
			"_id":  "$name",
			"sumb": bson.M{"$sum": "$b"},
			"suma": bson.M{"$sum": "$a"},
		},
	}

	sort := bson.M{
		"$sort": bson.M{
			"bdiva": -1,
		},
	}

	//order is important
	operations := []bson.M{group, project, sort}
	// operations := []bson.M{project, group, sort}
	pipe := c.Pipe(operations)

	ret := []interface {
	}{}
	err = pipe.All(&ret)
	if err != nil {
		panic(err.Error())
		return
	}
	for k, v := range ret {
		fmt.Printf("%+v: %+v\n", k, v)
	}
	fmt.Printf(" %v microseconds\n", time.Since(start)/1000000)
	println("//---------------------------AggregateQueryDemo end----------->>")
}

func BsonRawDemo() {
	println("//<<-------------------------BsonRawDemo start-----------")
	start := time.Now()

	var s = &struct {
		Name string `bson:"Name"`
	}{Name: "xxxx"}

	var js = &struct {
		Name string `json:"xxname" bson:"Name"`
	}{}
	// bs, err := bson.Marshal(s)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// var value interface{} = s
	// var value interface{} = bson.M{"some": "value"}
	data, err := bson.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	var doc bson.Raw
	err = bson.Unmarshal(data, &doc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("doc: %+v\n", doc)
	bs, err := bson.Marshal(doc)
	fmt.Printf("string(bs): %+v\n", string(bs))

	err = doc.Unmarshal(&js)

	bbs, err := json.Marshal(js)
	fmt.Printf("string(bbs): %+v\n", string(bbs))

	// a := make(map[string]string)
	// doc.Unmarshal(&a)
	// fmt.Printf("a: %+v\n", a)
	//
	// fmt.Printf("bs: %+v\n", bs)
	// raw := bson.Raw{0x08, []byte{0x01}} // true
	// // var a = make(map[string]interface{})
	// // err := raw.Unmarshal(a)
	// err := raw.Unmarshal(&struct{}{})
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Printf("a: %+v\n", a)
	fmt.Printf("BsonRawDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------BsonRawDemo end----------->>")
}
