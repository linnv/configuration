// Package main provides ...
package main

import (
	// "diy_package/buf"
	"diy_package/tmp"
	_ "diy_package/tmp2"
	_ "diy_package/tmp3"
	_ "diy_package/tmpa"
	"fmt"
)

var avar = 10

type d interface {
}

func init() {
	println("a")
}
func init() {
	println("b")
}

var by = []byte(`abc`)

func main() {
	// buf.Tmp()
	tmp.Tmp()
	println(avar)
	fmt.Printf("%s", by)
	fmt.Println()

}
