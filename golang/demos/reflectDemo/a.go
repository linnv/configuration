// Package main provides ...
package newDir

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

const NAME = "Name"

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	a := &A{"jialin", 211, 33.2}
	// fmt.Printf("  a: %+v\n", a)
	// updateStructFieldFromInterface(a)
	// fmt.Printf("end  a: %+v\n", a)
	Interf(a)
	println("-----------------------------JustDemo end>>>")
	return
}

type A struct {
	Name string  `json:"Name"`
	Age  int     `json:"Age"`
	FJ   float64 `json:"FJ"`
}

func (a A) Afunc() {
	return
}

func updateStructFieldFromInterface(itf interface{}) {
	typ := reflect.ValueOf(itf).Elem()
	r := typ.FieldByName(NAME)
	fmt.Printf("  r: %+v\n", r)
	r.SetString("xxxx")
	fmt.Printf("r.Kind().String()  : %+v\n", r.Kind().String())
}

func Field2fieldDemo(itf interface{}) {
	println("//<<-------------------------field2fieldDemo start-----------")
	start := time.Now()
	vo := reflect.ValueOf(itf)
	if vo.Kind() == reflect.Struct {
		panic("only supports ptr")
	}
	// // fmt.Printf("vo.: %+v\n", vo.Type())
	// typ := vo.Type()
	// // tye := reflect.ValueOf(vo.Interface()).Elem()
	// var s string
	// for i := 0; i < typ.NumField(); i++ {
	// 	v := typ.Field(i)
	// 	log.Printf("typ.: %+v\n", v)
	// 	s = v.Name
	//
	// 	log.Printf("typ.: %+v\n", vo.FieldByName(s).CanSet())
	//
	// 	// switch v.Type.Kind() {
	// 	// case reflect.String:
	// 	// 	tye.FieldByName(s).SetString(s)
	// 	// case reflect.Int, reflect.Int64:
	// 	// 	tye.FieldByName(s).SetInt(1000)
	// 	// case reflect.Float32, reflect.Float64:
	// 	// 	tye.FieldByName(s).SetFloat(1000.1)
	// 	// default:
	// 	// 	log.Println("reflect others: works")
	// 	// }
	// }
	// return
	// only interface or ptr works
	// tye := reflect.ValueOf(vo).Elem()
	tye := vo.Elem()
	typ := tye.Type()
	// fmt.Printf("typ: %+v\n", typ)
	var s string
	for i := 0; i < typ.NumField(); i++ {
		v := typ.Field(i)
		log.Printf("typ.: %+v\n", v)
		s = v.Name
		switch v.Type.Kind() {
		case reflect.String:
			tye.FieldByName(s).SetString(s)
		case reflect.Int, reflect.Int64:
			tye.FieldByName(s).SetInt(1000)
		case reflect.Float32, reflect.Float64:
			tye.FieldByName(s).SetFloat(1000.1)
		default:
			log.Println("reflect others:not support yet")
		}
	}
	fmt.Printf("itf: %+v\n", itf)

	fmt.Printf("field2fieldDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------field2fieldDemo end----------->>")
}

func Interf(Value interface{}) {
	// s := reflect.ValueOf(Value).Elem()
	// typ := reflect.ValueOf(Value).Elem()
	// fmt.Printf("typ.Kind().String()  : %+v\n", typ.Kind())

	// for i := 0; i < typ.NumField(); i++ {
	// 	r := typ.FieldByName(NAME)
	// 	fmt.Printf("  r: %+v\n", r)
	// 	r.SetString("xxxx")
	// 	fmt.Printf("r.Kind().String()  : %+v\n", r.Kind().String())

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
	// }

	// Type(),for struct member variable details and functions details
	typ := reflect.ValueOf(Value).Elem().Type()

	// Type(),for struct functions details only
	// typ := reflect.ValueOf(Value).Type()

	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		fmt.Printf(" p.Name : %+v\n", p.Name)
		fmt.Printf(" p.Type: %+v\n", p.Type)

		// for struct functions details
		// for i := 0; i < typ.NumMethod(); i++ {
		// 	p := typ.Method(i)
		// 	fmt.Printf("p: %+v\n", p)

	}

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

// type tfHandler struct{}
//
// func (t tfHandler) CityName() interface{} {
// 	return t.cityName()
// }
//
// func (tfHandler) cityName() string {
// 	return "xx"
// }
//
// func InittfHandler() {
// 	v := reflect.ValueOf(tfHandler{})
// 	t := v.Type()
// 	numMethod := v.NumMethod()
// 	for i := 0; i < numMethod; i++ {
// 		name := t.Method(i).Name
// 		if strings.ToLower(name[:1]) == name[:1] {
// 			continue
// 		}
// 		tmpName := strings.ToLower(name[:1]) + name[1:]
// 		realMethod := v.MethodByName(tmpName)
// 		if realMethod.Kind() == reflect.Invalid {
// 			panic(fmt.Sprintf(
// 				"transform method '%v' do not have a shadow method: '%v'",
// 				name, strings.ToLower(name[:1])+name[1:],
// 			))
// 		}
// 		// fmt.Printf("realMethod: %+v\n", realMethod)
// 		// fmt.Printf("tmpName: %+v\n", tmpName)
// 	}
// }
