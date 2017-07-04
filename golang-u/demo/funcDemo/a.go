// Package main provides ...
package demo

import (
	"fmt"
	"log"
	"time"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

type A struct {
	A string `json:"A"`
}

// func ReturnNil() error {
// func ReturnNil() *int {
func ReturnNil() *A {
	// return error(nil)
	// var a A = nil
	var a *A = nil
	return a
}

// func retDemo()( r ...int){
// println("//<<-------------------------retDemo start-----------")
// r=[]int{1,2,}
// println("//---------------------------retDemo end----------->>")
// 	return
// }

//ValidFunc implements ...
func ValidFunc(a string) bool {
	if a == "xx" {
		return true
	}
	return false
}

func InnerValidDemo() {
	println("//<<-------------------------InnerValidDemo start-----------")
	start := time.Now()
	ret := 0
	valid := func(f func(string) bool, p string, handler int) {
		if !f(p) {
			ret = handler
		}
	}
	valid(ValidFunc, "xx", 4)
	valid(ValidFunc, "3xx", 11)
	log.Printf("ret: %+v\n", ret)
	fmt.Printf("InnerValidDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------InnerValidDemo end----------->>")
}
