// Package main provides ...
package main

import (
	"fmt"
	"log"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "log"
	// "strings"

	// "demo/utility"
	"demo/demos/utility"
)

type Person struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Demo struct {
	ID   int    `json:"_id"`
	Name string `json:"name"`
}

func queryStrDemo() bson.M {
	println("<<<queryStrDemo---------------------------")

	query := make(bson.M, 1)
	ids := []int{1, 23}
	// bson.M{"id":bson.M{"$in":ids}}
	query["_id"] = bson.M{"$in": ids}
	println("-----------------------------queryStrDemo>>>")
	return query
}

// func main() {
// 	ret := queryStrDemo()
// 	fmt.Printf("ret: %+v\n", ret)
// 	ids := []int{1, 23}
// 	fmt.Printf("comp w: %+v\n", bson.M{"_id": bson.M{"$in": ids}})
//
// }

func main() {
	// session, err := mgo.Dial("server1.example.com,server2.example.com")
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("testdb").C("demo")

	//insert
	personArray := make([]interface{}, 0)
	// personArray := make([]*Person, 0)
	personArray = append(personArray, &Person{"ax", "ooooo"})
	personArray = append(personArray, &Person{"b", "fefnwfewjfjwef"})
	// err = c.Insert(personArray...)
	// err = c.Insert(&Person{"ax", "+1xxx111111111111"},
	// 	&Person{"b", "+222222222fjefei22"})
	if err != nil {
		log.Fatal(err)

	}

	// str := []string{"b", "new", "    ", "ewf"}
	str := []string{"new", "    ", "fef"}
	retStr := utility.StrArrayToStrWithSpliter(str, "|")
	fmt.Printf("retStr: %+v\n", retStr)
	//find
	result := []*Person{}
	// err = c.FindId(bson.M{"name": "a"}).One(&result)

	//regular find
	// tmp := bson.RegEx{Pattern: "new", Options: "i"}
	// fmt.Printf("tmp%+v\n", tmp)
	// err = c.Find(bson.M{"phone": &bson.RegEx{Pattern: "new", Options: "i"}}).All(&result) //Valid options as of this writing are 'i' for case insensitive matching, 'm' for multi-line matching, 'x' for verbose mode, 'l' to make \w, \W, and similar be locale-dependent, 's' for dot-all mode (a '.' matches everything), and 'u' to make \w, \W, and similar match unicode
	// err = c.Find(bson.M{"phone": &bson.RegEx{Pattern: retStr, Options: "i"}}).All(&result) //Valid options as of this writing are 'i' for case insensitive matching, 'm' for multi-line matching, 'x' for verbose mode, 'l' to make \w, \W, and similar be locale-dependent, 's' for dot-all mode (a '.' matches everything), and 'u' to make \w, \W, and similar match unicode
	// err = c.Find(bson.M{"name": "phone"}).All(&result) //Valid options as of this writing are 'i' for case insensitive matching, 'm' for multi-line matching, 'x' for verbose mode, 'l' to make \w, \W, and similar be locale-dependent, 's' for dot-all mode (a '.' matches everything), and 'u' to make \w, \W, and similar match unicode
	que := bson.M{}
	que["name"] = "b"
	err = c.Find(que).All(&result) //Valid options as of this writing are 'i' for case insensitive matching, 'm' for multi-line matching, 'x' for verbose mode, 'l' to make \w, \W, and similar be locale-dependent, 's' for dot-all mode (a '.' matches everything), and 'u' to make \w, \W, and similar match unicode
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range result {
		fmt.Println("phone:", v.Phone)
	}
}
