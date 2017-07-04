// Package main provides ...
package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("server1.example.com,server2.example.com")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")
	err = c.Insert(&Person{"a", "+1111111111111"},
		&Person{"b", "+22222222222"})
	if err != nil {
		log.Fatal(err)

	}
	result := Person{}
	err = c.FindId(bson.M{"name": "a"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("phone:", result.Phone)
}
