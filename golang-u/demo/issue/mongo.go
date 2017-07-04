// Package main provides ...
package main

import (
	"encoding/json"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
)

type Server struct {
	A string `json:"aname"`

	ServerIP string
}
type Serverslice struct {
	Servers []Server
}

func um() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Printf("%+v", s)
}
func jsonTag() {
	type Server struct {
		// ID 不会导出到JSON中
		ID int `json:"-"`

		// ServerName 的值会进行二次JSON编码
		ServerName  string `json:"serverName"`
		ServerName2 string `json:"serverName2,string"`

		// 如果 ServerIP 为空，则不输出到JSON串中
		// ServerIP string `json:"serverIP,omitempty"`
		ServerIP string `json:"-"`
	}

	s := Server{
		ID:          3,
		ServerName:  `Go "1.0" `,
		ServerName2: `Go "1.0" `,
		ServerIP:    ``,
	}
	b, _ := json.Marshal(s)
	os.Stdout.Write(b)
}

type User struct {
	Id   int    `json:"_id"`
	Name string `json:"name","-"`
}

func main() {
	session, err := mgo.Dial("127.0.0.1") //要连接的服务器和端口
	if err != nil {
		panic(err)
	}
	defer session.Close()
	//获取数据库，获取集合
	c := session.DB("test").C("test")
	//新增一个记录
	newTest := User{1199, "sunteng"}
	tmp, err := bson.Marshal(newTest)
	fmt.Printf("%+v\n", tmp)
	c.Insert(&newTest)
	checkErr(err)

	//单条记录查询
	result := User{}
	err = c.Find(bson.M{}).One(&result)
	checkErr(err)
	fmt.Printf("%+v\n", result)

	multiResult := []User{}
	//查找所有名字为 abc 的记录
	err = c.Find(&bson.M{"Name": "abc"}).All(&multiResult)
	checkErr(err)
	for i := 0; i < len(multiResult); i++ {
		fmt.Printf("%+v\n", multiResult[i])
	}

}
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
