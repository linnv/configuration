// Package main provides ...
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type D struct {
	Num int
}

type floatf float64

func parseDemo() {
	println("<<<parseDemo---------------------------")
	var v string
	// var v int64
	v = "333"
	x, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		panic(err.Error())
		return
	}
	var s int
	s = x
	// x, err := strconv.ParseInt(v,10, 32)
	fmt.Printf("s: %+v\n", s)
	fmt.Printf("valueof type(): %+v\n", reflect.ValueOf(x).Type())
	fmt.Printf("valueof kind(): %+v\n", reflect.ValueOf(x).Kind())

	println("-----------------------------parseDemo>>>")
	return
}

func reflectDemo() {
	println("<<<reflectDemo---------------------------")
	var x floatf = 222.333
	// var x float64 = 222.333
	// fmt.Printf("typeof: %+v\n", reflect.TypeOf(x))
	// fmt.Printf("typeof kind: %+v\n", reflect.TypeOf(x).Kind())
	//
	// fmt.Printf("valueof: %+v\n", reflect.ValueOf(x))
	// fmt.Printf("valueof type(): %+v\n", reflect.ValueOf(x).Type())
	// fmt.Printf("valueof kind(): %+v\n", reflect.ValueOf(x).Kind())
	// fmt.Printf("valueof int(): %+v\n", reflect.ValueOf(x).Float())
	// fmt.Printf("typeof: %+v\n", reflect.TypeOf(reflect.ValueOf(x).Interface()))
	// fmt.Printf("value typeof kind(): %+v\n", reflect.ValueOf(reflect.ValueOf(x).Interface()).Type())
	// fmt.Printf("valueof kind(): %+v\n", reflect.ValueOf(reflect.ValueOf(x).Interface()).Kind())
	//
	// //convert to interface{}
	// a := reflect.ValueOf(x).Interface()
	// a = int(222) //legal type of a is interface{}
	// // x = int(222) //unlegal: assign value of type int to variabl of type float
	// fmt.Printf("a: %+v\n", a)
	//
	// s := reflect.ValueOf(&x) //if  ValueOf(tmp) tmp is pointer , to get interface type  must invoke Elem() or will get the address,  no pointer  invoke Interface() directly
	// fmt.Printf("s.Interface(): %+v\n", s.Interface())
	// fmt.Printf("s.Interface(): %+v\n", s.Interface().(float64))

	v := reflect.ValueOf(&x).Elem() //if  ValueOf(tmp) tmp is pointer , to get interface type  must invoke Elem(),  no pointer  invoke Interface() directly;   if tmp is not pointer, v is not settable
	fmt.Printf("settablibility of v: %+v\n", v.CanSet())
	v.SetFloat(99934.4)
	fmt.Printf("v.interface: %+v\n", v.Interface())
	fmt.Printf("x: %+v\n", x)

	ReflectStruct()

	println("-----------------------------reflectDemo>>>")
	return
}

func main() {
	parseDemo()
	// var i interface{}
	// i := &D{Num: 222}
	// va := reflect.ValueOf(i).Elem()
	// fmt.Printf("va: %+v\n", va)
}

func ReflectStruct() {
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Kind(), f.Interface())
	}
}
