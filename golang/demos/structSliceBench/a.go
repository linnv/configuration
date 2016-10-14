// Package main provides ...
package newDir

import (
	"fmt"
	"strings"
)

type A struct {
	Name string `json:"Name"`
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	i := strings.Index("abdc", "c")
	fmt.Printf(" i: %+v\n", i)
	b := strings.Contains("abdc", "dc")
	fmt.Printf("b: %+v\n", b)
	println("-----------------------------JustDemo end>>>")
	return
}
