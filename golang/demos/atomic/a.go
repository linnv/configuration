// Package main provides ...
package demo

import (
	"sync/atomic"

	"github.com/linnv/logx"
)

type A struct {
	N string `json:"N"`
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	var a int64 = 1

	r := atomic.LoadInt64(&a)
	logx.Debugf("a: %+v\n", a)
	logx.Debugf("r: %+v\n", r)
	b := atomic.CompareAndSwapInt64(&a, 1, 9)
	logx.Debugf("b: %+v\n", b)
	r = atomic.LoadInt64(&a)
	logx.Debugf("a: %+v\n", a)
	logx.Debugf("r: %+v\n", r)
	var e chan *A
	if e != nil {
		println("no nil")
	}

	aslice := []int{1, 2}
	var o atomic.Value
	o.Store(aslice)
	ir := o.Load().([]int)
	logx.Debugf("aslice: %+v\n", aslice)
	ir[0] = 444
	logx.Debugf("ir: %+v\n", ir)
	logx.Debugf("look aslice: %+v\n", aslice)
	println("-----------------------------JustDemo end>>>")
	return
}
