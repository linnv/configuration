// Package main provides ...
package main

import (
	"fmt"
)

type A struct {
	N int
}

func main() {
	// a := new(A) //they are equal
	a := &A{}
	a.N = 1000
	fmt.Printf("a: %+v\n", a)

}
