// Package main provides ...
package main

import "fmt"

func valueBasicDemo(a *int) {
	println("<<<valueBasicDemo---------------------------")
	*a = 100
	println("-----------------------------valueBasicDemo>>>")
	return
}

func forDemo() {
	println("<<<forDemo---------------------------")
	n := 9
	m := 0
	// for i := 0; i < n; i >>= 1 {
	for i := n; i > 0; i >>= 1 {
		fmt.Printf("%v: works\n", i)
		m++
	}
	fmt.Printf("m: %+v\n", m)

	println("-----------------------------forDemo>>>")
	return
}

func main() {
	// forDemo()
	//@toDelete
	r := (-2 % 2)
	fmt.Printf("-1: %+v\n", r)
	// fmt.Println("good")
	// b := 9
	// valueBasicDemo(&b)
	// fmt.Printf("b: works %v\n", b)
	// io.Copy()

}
