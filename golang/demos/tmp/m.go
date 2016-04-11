// Package main provides ...
package main

import (
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func main() {
	session, err := mgo.Dial("127.0.0.1:3000") //要连接的服务器和端口
	if err != nil {
		panic(err)

	}
	defer session.Close()
	//获取数据库，获取集合
	c := session.DB("test").C("test")
	result := User{}
	err = c.Find(bson.M{}).One(&result) //payway为数据库中的字段
	if err != nil {
		panic(err)
	}
	fmt.Println("%+v", result)
	// fmt.Println("Phone:", result.message)
}
