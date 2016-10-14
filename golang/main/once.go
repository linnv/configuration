package main

import (
	"fmt"
	"math"
	"sync"

	"gopkg.in/mgo.v2"
)

var instanceInit sync.Once
var a int

func Once() {
	instanceInit.Do(func() {
		a++
		fmt.Println(a)
	})
}

func main() {
	// Once()
	// fmt.Printf("a: %+v\n", a)
        //
	// Once()
	// fmt.Printf("2a: %+v\n", a)
	// }
