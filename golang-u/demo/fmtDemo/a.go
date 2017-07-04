// Package main provides ...
package demo

import "fmt"

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	a := []int{1, 2, 3}
	//@toDelete
	fmt.Printf("  a: %d\n", a)
	println("-----------------------------JustDemo end>>>")
	return
}
