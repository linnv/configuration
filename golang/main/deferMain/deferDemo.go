// Package main provides ...
package main

import (
	"fmt"
)

func DeferNoParametersDemo() {
	println("<<<DeferNoParametersDemo---------------------------")

	for i := 0; i < 9; i++ {
		defer func() {
			fmt.Printf("i: %+v\n", i)
		}()
	}

	println("-----------------------------DeferNoParametersDemo>>>")
}

func DeferWithParametersDemo() {
	println("<<<DeferWithParametersDemo---------------------------")

	for i := 0; i < 9; i++ {
		defer func(o int) {
			fmt.Printf("i: %+v\n", o)
		}(i)
	}
	println("-----------------------------DeferWithParametersDemo>>>")
}

func main() {
	DeferNoParametersDemo()
	DeferWithParametersDemo()
}

func noRzturnDemo() {
	// noReturnDemo()
	println("<<<noReturnDemo---------------------------")
	println("works")
	defer func() {
		println("good")
	}()
	println("-----------------------------noReturnDemo>>>")
	return
}
