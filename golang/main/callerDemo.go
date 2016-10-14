// Package main provides ...
package main

import (
	"fmt"
	"runtime"
	"time"
)

func callerDemo() {
	println("//<<-------------------------callerDemo start-----------")
	start := time.Now()

	// var initpc, fn, l, ok = runtime.Caller(0)
	var initpc, fn, l, ok = runtime.Caller(-1)
	fmt.Printf("initpc: %+v\n", initpc)
	fmt.Printf("fn: %+v\n", fn)
	fmt.Printf("l: %+v\n", l)
	fmt.Printf("ok: %+v\n", ok)
	fmt.Printf("callerDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------callerDemo end----------->>")
}

func main() {
	callerDemo()
}
