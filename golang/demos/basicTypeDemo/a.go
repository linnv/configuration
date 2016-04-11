// Package main provides ...
package newDir

import (
	"demo/demos/utility"
	"fmt"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	a := 1e8
	b := 1e3 //float64
	fmt.Printf("  a: %f\n", a)
	// b: %!d(float64=1000)
	fmt.Printf("  b: %d\n", b)
	c, err := utility.CoerceFloat64(b)
	if err != nil {
		panic(err.Error())
		return
	}
	//@toDelete
	fmt.Printf("  c: %+v\n", c)

	fmt.Printf("getClicksRatio(300,10): %v\n", getClicksRatio(100, 33))
	println("-----------------------------JustDemo end>>>")
	return
}

func algorithmOperation() {
}

func getClicksRatio(impressions, clicks int) float32 {
	return float32(int64(float32(clicks)/float32(impressions)*100)) / 100
	// return float32(clicks) / float32(impressions)
}
