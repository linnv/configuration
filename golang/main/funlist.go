// Package main provides ...
package main

import "fmt"

func afDemo(n int) error {
	println("<<<afDemo---------------------------")

	println(n, "-----------------------------afDemo>>>")
	return nil
}
func bfDemo(n int) error {
	println("<<<bfDemo---------------------------")

	println(n, "-----------------------------bfDemo>>>")
	return nil
}

var Flist []func(int) error

func main() {
	var n int = 100
	Flist = []func(int) error{
		afDemo, //func name only
		bfDemo,
	}
	for k, v := range Flist {
		fmt.Printf("%+v: %+v\n", k, v(n)) //put parameters to fun
	}

	c := make([]int, 10, 20)
	for k, v := range c {
		if k == 2 {
			continue
		}
		fmt.Printf("%+v: %+v\n", k, v)
	}
}
