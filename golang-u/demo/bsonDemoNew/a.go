// Package main provides ...
package demo

import (
	"encoding/json"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	bs := bson.M{"xx": 1}
	bs["a"], bs["b"] = 2, 4
	//@toDelete
	fmt.Printf("bs: %+v\n", bs)
	println("-----------------------------JustDemo end>>>")
	return
}

type JNotTagPerson struct {
	Name  string
	Phone string
	Id    int
}

type JPerson struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Id    int    `json:"Id"`
}

type BPerson struct {
	Name  string `bson:"Name"`
	Phone string `bson:"Phone"`
	Id    int    `bson:"Id"`
}

type BNotTagPerson struct {
	Name  string
	Phone string
	Id    int
}

func StructTagBsonUnMarshalDemo(str []byte) (b *BNotTagPerson, err error) {
	println("<<<structTagBsonUnmarshalDemo---------------------------")
	b = new(BNotTagPerson)
	err = bson.Unmarshal(str, b)
	println("-----------------------------structTagDemo>>>")
	return
}

func StructTagJsonUnMarshalDemo(str []byte) (b *JNotTagPerson, err error) {
	println("<<<structTagUnmarshalJsonDemo---------------------------")
	b = new(JNotTagPerson)
	err = json.Unmarshal(str, b)
	println("-----------------------------structTagDemo>>>")
	return
}

func StructTagBsonDemo(b *BPerson) (str []byte, err error) {
	println("<<<structTagBsonDemo---------------------------")
	ret, y := bson.Marshal(b)
	println("-----------------------------structTagDemo>>>")
	return ret, y
}

func StructTagJsonDemo(b *JPerson) (str []byte, err error) {
	println("<<<structTagJsonDemo---------------------------")
	ret, y := json.Marshal(b)
	println("-----------------------------structTagDemo>>>")
	return ret, y
}
