// Package main provides ...
package newDir

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b")) //parameter is calculated when invoke defer, e.g. when parameter is a function, this funciton will be invoking and the deferFunc will be invoke when this deferFunc returns
	fmt.Println("in b")
	a()
}

func JustDemo() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	b()
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}
