// Package main provides ...
package main

import "fmt"

type Aer interface {
	Light() int
	Sname() string
	Data() interface{}
}

type B struct {
	A int
}

func (this B) Sname() string {
	return "name works"
}

func (this B) Data() interface{} {
	return this
}

func (this B) Light() int {
	println("light of B")
	return this.A
}

func (this B) AddLight() int {
	println("light of B")
	return this.A
}

func multiParametersDemo(v ...interface{}) {
	println("<<<multiParametersDemo---------------------------")
	for k, v := range v {
		fmt.Printf("%+v: %+v\n", k, v)
	}

	println("-----------------------------multiParametersDemo>>>")
	return
}

// func (this B) Hight(error) {
//
// }

func main() {
	// multiParametersDemo("2", 3, 334, true, 3.4)

	// c := &B{A: 9}
	// md, err := json.Marshal(c)
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }

	// t := reflect.ValueOf(c).Interface()
	// fmt.Printf("t: %+v\n", t)
	//
	// fmt.Printf("value typeof kind(): %+v\n", reflect.ValueOf(reflect.ValueOf(c).Interface()).Type())
	// fmt.Printf("value typeof kind(): %+v\n", reflect.ValueOf(t).Type())
	// fmt.Printf("valueof kind(): %+v\n", reflect.ValueOf(reflect.ValueOf(c).Interface()).Kind())
	// var t interface{}
	// if err = json.Unmarshal(md, &t); err == nil {
	// 	fmt.Printf("t: %+v\n", t)
	// 	// b = true
	// 	return
	// }

	var a Aer
	a = B{A: 9}
	a.Light()
	// f := a.(Aer)
	// f := t.(*B)   //type assertion must depen on a dynamil type like interface{}
	// fmt.Printf("t: %+v\n", t.A)
	// fmt.Printf("f: %+v\n", f.Light())
	// i := 10
	// s := strconv.Itoa(i)
	// fmt.Printf("f: %+v\n", f.A)
	// a = c
	// fmt.Printf("a: %+v\n", a)
	// fmt.Printf("a.Sname(): %+v\n", a.Sname())
	// fmt.Printf("a.Data(): %+v\n", a.Data())
	// tmp := a.(*B)
	// fmt.Printf("tmp: %+v\n", tmp.A)
	// d := a.Light()
	// println(d)
}
