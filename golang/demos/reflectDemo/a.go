// Package main provides ...
package demo

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/linnv/logx"
)

const NAME = "Name"

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("type of Type:", t)
	if t == v.Type() {
		logx.Debugln("type from Value equales to type from Type")
	}
	vk := v.Kind()
	tk := t.Kind()
	logx.Debugf("vk: %+v\n", vk)
	logx.Debugf("tk: %+v\n", tk)

	var data []int
	dataShort := []int{}
	fmt.Printf("data: %+v,cap(data):%d,len(data):%d arrd:%v \n", data, cap(data), len(data), &data)
	fmt.Printf("dataShort: %+v,cap(dataShort):%d,len(dataShort):%d arrd:%v \n", dataShort, cap(dataShort), len(dataShort), &dataShort)
	if reflect.DeepEqual(data, dataShort) {
		println("e")
	} else {
		println("ne")
	}
	a := &A{"jialin", 211}
	//@toDelete
	fmt.Printf("  a: %+v\n", a)
	// Interf(a)
	updateStructFieldFromInterface(a)
	fmt.Printf("end  a: %+v\n", a)
	println("-----------------------------JustDemo end>>>")
	return
}

type GS interface {
	Get() int
	Set(i int)
}

type A struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func (a *A) XxxHandlerGet() int {
	return 10 * a.Age
}

func (a *A) Get() int {
	return a.Age
}

func (a *A) Set(i int) {
	a.Age = i
}

type AutoRouter struct {
	MethodName string
	T          reflect.Type
}

func NewByReflectDemo() {
	println("//<<-------------------------NewByReflectDemo start-----------")
	start := time.Now()
	var s reflect.Type
	a := &A{}
	a.Set(9)
	log.Printf("a.Get(): %d\n", a.Get())
	log.Printf("XxxHandlerGet: %+v\n", a.XxxHandlerGet())

	reflectVal := reflect.ValueOf(a)
	s = reflect.Indirect(reflectVal).Type() //reflect.Type

	rv := reflect.New(s)

	av, ok := rv.Interface().(GS)
	if !ok {
		logx.Debugln("not ok")
	}

	log.Printf("av.Get(): %+v\n", av.Get())
	av.Set(111)

	var in []reflect.Value
	const methodName = "XxxHandlerGet"
	method := rv.MethodByName(methodName)
	xxr := method.Call(in)

	fmt.Printf("xxr: %+v,cap(xxr):%d,len(xxr):%d arrd:%v \n", xxr, cap(xxr), len(xxr), &xxr[0])
	xi := xxr[0].Interface()
	log.Printf("xi: %+v\n", xi)
	log.Printf("updated by interface av.Get(): %+v\n", av.Get())
	log.Printf("updated by a.Get(): %d\n", a.Get())
	log.Printf("XxxHandlerGet: %+v\n", a.XxxHandlerGet())

	fmt.Printf("NewByReflectDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------NewByReflectDemo end----------->>")
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

type SA struct {
	Name string `json:"Name"`
}

func (a *SA) One(n int) {

	logx.Debugf("invoking method n: %+v\n", n)
}

func (a SA) Two(n int) {

}

func structInfoDemo() {
	println("//<<-------------------------structINfoDemo start-----------")
	start := time.Now()
	a := &SA{}
	// a := SA{}
	v := reflect.ValueOf(a)
	logx.Debugf("v.Kind(): %+v\n", v.Kind())
	n := v.NumMethod() //method number of kind of ptr includes pointer and value receiver
	t := v.Type()      //the defined struct type: you can get attibutes like mothed name,filed name here
	logx.Debugf("t: %+v\n", t)
	logx.Debugf("t.Name(): %+v\n", t.String())
	in := strings.LastIndex(t.String(), ".") + 1
	logx.Debugf("struct name t.String(): %+v\n", t.String()[in:])

	// n := v.NumField()
	logx.Debugf("n: %+v\n", n)
	for i := 0; i < n; i++ {
		m := t.Method(i).Name
		mtt := t.Method(i).Type
		logx.Debugf("method name m: %+v\n", m)
		logx.Debugf("mtt: %+v\n", mtt)

		mk := v.Method(i).Kind()
		logx.Debugf("mk: %+v\n", mk)
		// aslo work for getting function type
		switch handler := v.Method(i).Interface().(type) {
		case func(int):
			handler(22)
		}
		// mt := v.Method(i).Type()
		// logx.Debugf("mt: %+v\n", mt)
		// logx.Debugf("func type mt.String(): %+v\n", mt.String())
		// switch mt.String() {
		// case "func(int)":
		// 	v.Method(i).Interface().(func(int))(3)
		// 	logx.Debugln("good func")
		// default:
		// 	logx.Debugln("bad func")
		// }
	}
	logx.Debugf("v: %+v\n", v)
	fmt.Printf("structINfoDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------structINfoDemo end----------->>")
}

func RuneCountDemo(str string) int {
	println("//<<-------------------------RuneCountDemo start-----------")
	start := time.Now()
	fmt.Printf("RuneCountDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------RuneCountDemo end----------->>")
	return utf8.RuneCountInString(str)
}
