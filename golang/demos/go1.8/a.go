// Package main provides ...
package demo

import (
	"fmt"
	"sort"
	"time"

	"github.com/linnv/logx"
)

type A struct {
	B byte `json:"B"`
}

func SortDemo() {
	println("//<<-------------------------SortDemo start-----------")
	start := time.Now()
	const count = 10
	as := make([]A, 0, 10)
	// as := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		as = append(as, A{byte(i)})
		// as = append(as, i)
	}
	logx.Debug("first as: %+v\n", as)
	sort.Slice(as, func(i, j int) bool { return as[i].B > as[j].B })
	// sort.Slice(as, func(i, j int) bool { return as[i] < as[j] })
	logx.Debug("as: %+v\n", as)
	fmt.Printf("SortDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------SortDemo end----------->>")
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	SortDemo()
	println("-----------------------------JustDemo end>>>")
	return
}
