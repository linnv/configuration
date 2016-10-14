// Package main provides ...
package newDir

import (
	"fmt"
	"io"
	"io/ioutil"
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
	// buf := make([]byte, 8)
	// if _, err := io.ReadFull(r, buf); err == io.EOF {
	// 	// return io.ErrUnexpectedEOF
	// } else if err != nil {
	// 	// return err
	// }
	ioutil.ReadFile("./ddd")
	return &LimitedReaderDemo{R: r, N: n}
}

func bufioRead() {
	// reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	// line, err := reader.ReadSlice('\n')
	// utils.CheckErr(err)
	// fmt.Printf("line: %+v\n", string(line))
	// line2, err := reader.ReadSlice('\n')
	// // utils.CheckErr(err)
	// fmt.Printf("line2: %+v\n", string(line2))

	// line, err := reader.ReadBytes('\n')
	// utils.CheckErr(err)
	// fmt.Printf("line: %+v\n", string(line))
	// line2, err := reader.ReadBytes('\nSlice')
	// // utils.CheckErr(err)
	// fmt.Printf("line2: %+v\n", string(line2))
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	bufioRead()
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
	}
	fmt.Printf("r: %+v\n", lr)
	println("-----------------------------JustDemo end>>>")
	return
}
