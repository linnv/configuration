// Package main provides ...
package newDir

import (
	"fmt"
	"reflect"
)

const NAME = "Name"

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	a := &A{"jialin", 211}
	//@toDelete
	fmt.Printf("  a: %+v\n", a)
	// Interf(a)
	updateStructFieldFromInterface(a)
	fmt.Printf("end  a: %+v\n", a)
	println("-----------------------------JustDemo end>>>")
	return
}

type A struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func updateStructFieldFromInterface(itf interface{}) {
	typ := reflect.ValueOf(itf).Elem()
	r := typ.FieldByName(NAME)
	//@toDelete
	fmt.Printf("  r: %+v\n", r)
	r.SetString("xxxx")
	//@toDelete
	fmt.Printf("r.Kind().String()  : %+v\n", r.Kind().String())
}

func Interf(Value interface{}) {
	// s := reflect.ValueOf(Value).Elem()
	typ := reflect.ValueOf(Value).Elem()
	//@toDelete
	fmt.Printf("typ.Kind().String()  : %+v\n", typ.Kind())

	for i := 0; i < typ.NumField(); i++ {
		r := typ.FieldByName(NAME)
		//@toDelete
		fmt.Printf("  r: %+v\n", r)
		r.SetString("xxxx")
		//@toDelete
		fmt.Printf("r.Kind().String()  : %+v\n", r.Kind().String())

		// p := typ.Field(i)
		// //@toDelete
		// fmt.Printf("i%d  p: %+v\n", i, p)
		// if i == 0 {
		// 	p.SetString("xxx")
		// }
		//@toDelete
		// fmt.Printf("  p: %+v\n", p)
		// pn := p.FieldByName(NAME).String()
		// //@toDelete
		// fmt.Printf("  pn: %+v\n", pn)
		// fmt.Printf(" p.Name : %+v\n", p.Name)
		// fmt.Printf(" p.Type: %+v\n", p.Type)
	}
	// typ := reflect.TypeOf(Value).Elem()
	// for i := 0; i < typ.NumField(); i++ {
	// 	p := typ.Field(i)
	// 	//@toDelete
	// 	fmt.Printf(" p.Name : %+v\n", p.Name)
	// 	fmt.Printf(" p.Type: %+v\n", p.Type)
	// }
	//@toDelete
	// fmt.Printf("  s: %+v\n", s)
	// // println(s.String())
	// // println(s.Elem().String())
	// metric, ok := typ.FieldByName(NAME)
	// if !ok {
	// 	//@toDelete
	// 	fmt.Printf("  not: %+v\n", "not")
	// }
	// fmt.Println(metric)
}
