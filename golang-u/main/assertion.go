// Package main provides ...
package main

import (
	"fmt"
	"time"
)

type A struct {
	Count int64 `json:"Count"`
}

func AssertDemo(i interface{}) {
	println("<<<AssertDemo---------------------------")
	switch i.(type) {
	case *A:
		println("struct A")
	}
	println("-----------------------------AssertDemo>>>")
	return
}
func defaultValueDemo() (b bool) {
	println("<<<defaultValueDemo---------------------------")
	return
	println("-----------------------------defaultValueDemo>>>")
	return
}

type Duration struct {
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
}

func (this *Duration) checkDuration() (code int, ok bool) {
	if this.StartTime < 0 || this.EndTime < 0 {
		ok = false
		return
	}

	if this.StartTime == 0 || this.EndTime == 0 {
		this.EndTime = time.Now().Unix()
	}
	if this.EndTime < this.StartTime {
		ok = false
		return
	}
	ok = true
	return
}
func main() {
	b := defaultValueDemo()
	fmt.Printf("b: %+v\n", b)
	// a := &A{Count: 100}
	// AssertDemo(a)

}

func check(i interface{}) (code int, ok bool) {
	ok = true
	switch i.(type) { //bench mark
	case *Duration:
		request := i.(*Duration)
		return request.checkDuration()
	}
	return
}
