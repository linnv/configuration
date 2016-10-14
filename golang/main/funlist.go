// Package main provides ...
package main

import (
	"fmt"
	"runtime"
)

func afDemo(n int) error {
	println("<<<afDemo---------------------------")
	const size = 4096
	buf := make([]byte, size)
	buf = buf[:runtime.Stack(buf, false)]
	// buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("string(buf): %+v\n", string(buf))
	println(n, "-----------------------------afDemo>>>")
	return nil
}
func bfDemo(n int) error {
	println("<<<bfDemo---------------------------")
	// debug.PrintStack()
	println(n, "-----------------------------bfDemo>>>")
	return nil
}

var Flist []func(int) error

func main() {
	var n int = 100
	Flist = []func(int) error{
		afDemo, //func name only
		bfDemo,
	}
	for k, v := range Flist {
		fmt.Printf("%+v: %+v\n", k, v(n)) //put parameters to fun
	}
	var ms = new(runtime.MemStats)
	runtime.ReadMemStats(ms)
	fmt.Printf("ms: %+v\n", ms)
	// time.Sleep(time.Second * 900)
}
