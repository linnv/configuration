// Package main provides ...
package demo

import (
	"fmt"
	"io"
)

type LimitedReaderDemo struct {
	R io.Reader // underlying reader
	N int64     // max bytes remaining
}

type IOReaderDemo struct {
	R io.Reader // underlying reader
}

func (this IOReaderDemo) Read(bs []byte) (n int, err error) {
	println("undeylying reader")
	return 122, nil
}

func (this LimitedReaderDemo) Read(bs []byte) (n int, err error) {
	println("implement reader")
	//R must be instanced beforce using it's method, or nil pointer error will occurs
	if this.R == nil {
		println("error: reader is not an instance")
		//R must be instanced beforce using it's method, or nil pointer error will occurs
		return 0, nil

	}
	return this.R.Read(bs)
	// return 1, nil
}

//instance Reader interface
// The underlying implementation is a *LimitedReaderDemo.
func ReadDemo(r io.Reader, n int64) io.Reader {
	return &LimitedReaderDemo{R: r, N: n}
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	bs := []byte(`a`)
	// ior := IOReaderDemo{}
	// a := ReadDemo(ior, 10)

	//implement member R (io.Reader interface is not initalized,run fails)
	// panic: runtime error: invalid memory address or nil pointer dereference [recovered
	a := LimitedReaderDemo{}
	lr, err := a.Read(bs)
	// r, err := a.R.Read(bs)
	if err != nil {
		panic(err.Error())
		return
	}
	fmt.Printf("r: %+v\n", lr)
	println("-----------------------------JustDemo end>>>")
	return
}
