package main

import (
	"fmt"
	"math"

	"gopkg.in/mgo.v2"
)

func main() {
	rate, years := 0.95, []float64{3, 5, 10, 20, 30}
	for i := 0; i < len(years); i++ {
		fmt.Printf("after %d years actual value is: %+v\n", int(years[i]), math.Pow(rate, years[i]))
	}
	select {}

	// after 3 years actual value is: 0.8573749999999999
	// after 5 years actual value is: 0.7737809375
	// after 10 years actual value is: 0.5987369392383789
	// after 20 years actual value is: 0.35848592240854216
	// after 30 years actual value is: 0.2146387639429375

	// b := now.BeginningOfDay().Unix()
	// e := now.EndOfDay().Unix()
	// //@toDelete
	// fmt.Printf("b: %+v\n", b)
	// //@toDelete
	// fmt.Printf("e: %+v\n", e)
	// fmt.Println("hi")
	//
	// time.Sleep(30 * time.Second)
}

func suntengFindOperationDemo() (err error) {
	println("//<<-------------------------OrOperationDemo start-----------")
	// session, err := mgo.Dial("192.168.100.44:27017")
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		return
	}
	defer session.Close()
	// session.SetMode(mgo.Monotonic, true)
	c := session.DB("testdb").C("demo")
	// q := bson.M{}
	// for i := 0; i < 30; i++ {
	// 	err = c.Insert(&JPerson{"json", "sort demo", i})
	// }
	var result []interface{}
	// limit, offset := 0, 0
	// sort := "Id"
	// doQeury := ""
	err = c.Find(nil).All(&result)
	if err != nil {
		fmt.Println("not found")
		return
	}
	fmt.Println("found")
	for k, v := range result {
		fmt.Printf("%+v: %+v\n", k, v)
	}

	// return c.Find(query).Select(bson.M{"_id": 1, "CASUserInfo.Name": 1, "Independent": 1, "Type": 1, "Status": 1, "Seller": 1, "CreateTime": 1, "OwnerId": 1}).Sort(sort...).Skip(offset).Limit(limit).All(&result)

	println("//---------------------------OrOperationDemo end----------->>")
	return
}
