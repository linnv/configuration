// Package main provides ...
package newDir

import (
	"fmt"
	"time"
)

func JustDemo() {
	println("//<<-------------------------JustDemo start-----------")
	start := time.Now()

	fmt.Printf("JustDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------JustDemo end----------->>")
}
