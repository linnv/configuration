// Package main provides ...
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("math.MinInt16: %+v\n", math.MinInt16)
	fmt.Printf("math.MaxInt16: %+v\n", math.MaxInt16)
	fmt.Printf("math.MinInt32: %+v\n", math.MinInt32)
	fmt.Printf("math.MaxInt32: %+v\n", math.MaxInt32)
	fmt.Printf("math.MinInt64: %+v\n", math.MinInt64)
	fmt.Printf("math.MaxInt64: %+v\n", math.MaxInt64)
	fmt.Printf("math.MaxUint32: %+v\n", math.MaxUint32)
	// fmt.Printf("math.MaxUint64: %s\n", string(math.MaxUint64>>10))
	println("math.MaxUint64:", math.MaxUint64)

	fmt.Printf("math.MaxInt32/1024: %+v\n", math.MaxInt32>>20)
	fmt.Printf("1024*1024>>20: %+v\n", (1024*1024)>>20)
}
