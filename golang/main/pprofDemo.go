// Package main provides ...
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
)

var (
	cpuProfFileP = flag.String("cpuproffile",
		"cpu-bidlog.prof", "the file to save cpu performance")
)

type A struct {
	N string `json:"N"`
}

func NewA() {
	for i := 0; i < 100; i++ {
		a := A{N: "jjj"}
		fmt.Printf("  a: %+v\n", a)
	}
}

func main() {
	flag.Parse()

	cpuProfFile := *cpuProfFileP

	if cpuProfFile != "" {
		f, _ := os.Create(cpuProfFile)
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	NewA()
}
