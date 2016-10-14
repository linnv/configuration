// Package main provides ...
package newDir

import (
	"fmt"
	"time"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	DeferFuntion()
	println("-----------------------------JustDemo end>>>")
	return
}

func DeferFuntion() (i int) {
	i = 9
	defer fmt.Printf("inner first i: %+v\n", i)

	defer func() {
		fmt.Printf("returning a constant interger: %+v\n", i)
		fmt.Printf("inner time.Now(): %+v\n", time.Now())
		fmt.Printf("inner i: %+v\n", i)
		i++
		fmt.Printf("plus one inner time.Now(): %+v\n", time.Now())
		fmt.Printf("plus one inner i: %+v\n\n", i)
	}()
	i++
	fmt.Printf("outer time.Now(): %+v\n", time.Now())
	fmt.Printf("out i: %+v\n", i)
	//equals i=3
	return 3
}
