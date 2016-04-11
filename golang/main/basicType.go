// Package main provides ...
package main

import (
	"encoding/json"
	"fmt"
	"net"
)

const (
	ActionAdd = iota + 1 // 0 增加
)

type Response1 struct {
	Page   int
	Fruits []string
	B      int
}

func main() {

	var a int
	a = ActionAdd
	fmt.Printf("a: %+v\n", a)
	str := `{"page": 1,"b":true,"fruits": ["apple", "peach"]}`
	// str := `{"page": 1,"B":false ,"fruits": ["apple", "peach"]}`
	res := &Response1{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		fmt.Printf("err.Error(): %+v\n", err.Error())
		return
	}
	fmt.Println(res)
	net.ListenPacket()
	//  byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// var dat map[string]interface{}
	// 	if err := json.Unmarshal(byt, &dat); err != nil {
	//         panic(err)
	//     }
	//     fmt.Println(dat)
}
