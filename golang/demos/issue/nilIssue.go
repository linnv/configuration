package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}
type String struct {
	data string
}

func (s *String) String() string {
	return s.data

}

func GetString() *String {
	return nil

}

func CheckString(s Stringer) bool {
	fmt.Printf("empty string%+v\n", s)
	fmt.Printf("nil:%+v\n", nil)
	return s == nil
}

func main() {
	println(CheckString(GetString()))
}

//
// type a struct {
// 	Data string
// }
//
// func demox(f func() *string) *string {
// 	// tmp := a{Data: "10"}
// 	// tmp := &a{Data: "10"}
// 	// return tmp
// 	return f()
// }
//
// func demoxx() *string {
// 	// tmp := a{Data: "10"}
// 	// tmp := &a{Data: "10"}
// 	// return tmp
// 	return nil
// }
//
// func demo() interface{} {
// 	// tmp := a{Data: "10"}
// 	// tmp := &a{Data: "10"}
// 	// return tmp
// 	return nil
// }
//
// type tmpType a
//
// func main() {
// 	var d interface{}
// 	d = 19
// 	d = 9.0
// 	d = true
// 	d = `haha`
// 	fmt.Printf("%+v %+v\n", reflect.TypeOf(d), d)
// 	// var x interface{}
// 	// tmp := demoxx()
// 	tmp := demox(demoxx)
//
// 	fmt.Printf("%+v %+v\n", reflect.TypeOf(tmp), tmp)
// 	// tmp := demo(k
// 	println(tmp == nil)
// 	// fmt.Printf("%+v\n", d)
// 	// var m interface{}
// 	// m = demo()
// 	// fmt.Printf("%+v %+v\n", reflect.TypeOf(m), m)
// 	// println(m.Data)
// }
