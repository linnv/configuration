// Package main provides ...
package main

import (
// "fmt"
)

type demo2 struct {
	a int
}
type demo struct {
	a    int
	sfun func()
}

type istruct interface {
	funa()
}

func funa() {
	println("damn")
}

func demof() {
	println("testing")
}

var demoArr []istruct = make([]istruct, 0)

func (t *demo2) funa() {
	println("aaaa", t.a)
}

func (t *demo) funa() {
	println("bbbbbb", t.a)
}

func do() *demo {
	return &demo{a: 19, sfun: demof}
}

//
// func m(i istruct) {
// 	i.funa()
// }
func main() {
	// //interface
	// n := &demo{a: 1}
	// n2 := &demo2{a: 2}
	// demoArr = append(demoArr, n)
	// demoArr = append(demoArr, n2)
	// // demoArr = append(demoArr, istruct)
	// for _, singleDemo := range demoArr {
	// 	singleDemo.funa()
	// }
	//

	// // structure and function pointer
	// v := do()
	// v.sfun()
}
