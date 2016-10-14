// Package main provides ...
package main

/*
#include "a.h"
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"time"
)

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	start := time.Now()
	fmt.Printf("nunmber from cgo Random(): %+v\n", Random())
	C.cgoPrintDemo()

	fmt.Printf("JustDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
}

// func JustDemo() {
// 	println("//<<-------------------------JustDemo start-----------")
// 	println("//---------------------------JustDemo end----------->>")
// }
