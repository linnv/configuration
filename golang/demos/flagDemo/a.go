// Package main provides ...
package newDir

import (
	"flag"
	"fmt"
	"time"
)

func FlagDemo(name string, ss []string) int {
	println("//<<-------------------------FlagDemo start-----------")
	start := time.Now()
	var pid int
	flag.IntVar(&pid, name, -1, "")
	flag.CommandLine.Parse(ss)
	fmt.Printf("FlagDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------FlagDemo end----------->>")
	return pid
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}
