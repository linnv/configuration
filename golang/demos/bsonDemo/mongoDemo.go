package bsonDemo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gopkg.in/mgo.v2/bson"
)

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
func RawDemo() {
	println("//<<-------------------------RawDemo start-----------")
	start := time.Now()
	var a = struct {
		A string `bson:"A "`
		B int64  `bson:"B "`
	}{}
	bs, err := bson.Marshal(a)
	if err != nil {
		panic(err.Error())
	}
	raw := bson.Raw(bs)
	s := bson.Unmarshal(raw)
	fmt.Printf("s: %+v\n", s)
	fmt.Printf("string(s): %+v\n", string(s))
	fmt.Printf("RawDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------RawDemo end----------->>")
}

func JustDemo() {
	println("//<<-------------------------JustDemo start-----------")
	start := time.Now()
	var bdemo bson.M
	// bdemo := bson.M{}
	// if bdemo == nil {
	// 	os.Stdout.Write(append([]byte("damn good"), '\n'))
	// }
	// bdemo["u"] = 3
	if _, ok := bdemo["u"]; ok {
		os.Stdout.Write(append([]byte("good"), '\n'))

	} else {
		os.Stdout.Write(append([]byte("bad"), '\n'))
	}
	fmt.Printf("JustDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------JustDemo end----------->>")
}
