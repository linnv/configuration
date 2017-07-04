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
			for condition {
				for condition {

				}

			}
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

func noRzturnDemo() {
	println("<<<noReturnDemo---------------------------")
	println("works")
	defer func() {
		println("good")
	}()
	println("-----------------------------noReturnDemo>>>")
	return
}

func main() {
	noReturnDemo()
	// DeferNoParametersDemo()
	// DeferWithParametersDemo()
}
