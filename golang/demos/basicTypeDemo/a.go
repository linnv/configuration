// Package main provides ...
package newDir

import (
	"fmt"
	"log"
	"os"
	"time"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")

	UnescapeByteDemo()
	// log.Printf("uint(0): %b\n", uint(0))
	// log.Printf("u)nt(0): %v\n", (^uint(0))>>1)
	// a := 1e8
	// b := 1e3 //float64
	// fmt.Printf("  a: %f\n", a)
	// // b: %!d(float64=1000)
	// fmt.Printf("  b: %d\n", b)
	// c, err := utility.CoerceFloat64(b)
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }
	// //@toDelete
	// fmt.Printf("  c: %+v\n", c)
	//
	// fmt.Printf("getClicksRatio(300,10): %v\n", getClicksRatio(100, 33))
	// simpleUnitDemo()
	println("-----------------------------JustDemo end>>>")
	return
}
func UnescapeByteDemo() {
	println("//<<-------------------------UnescapeByteDemo start-----------")
	start := time.Now()
	s := ".snug \u0026 raw was I ere I saw war \u0026 gunS"
	// bs:=[]byte(s)
	fmt.Fprint(os.Stdout, "%s", s)
	// byteS = bytes.Replace(byteS, []byte("\\u0026"), []byte("&"), -1)

	fmt.Printf("UnescapeByteDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------UnescapeByteDemo end----------->>")
}

func simpleUnitDemo() {
	println("//<<-------------------------simpleUnitDemo start-----------")
	start := time.Now()
	log.Printf("0x1: %b\n", 0x1)
	log.Printf("0x2: %b\n", 0x2)
	log.Printf("0x3: %b\n", 0x3)
	fmt.Printf("simpleUnitDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------simpleUnitDemo end----------->>")
}

func algorithmOperation() {

}

func getClicksRatio(impressions, clicks int) float32 {
	return float32(int64(float32(clicks)/float32(impressions)*100)) / 100
	// return float32(clicks) / float32(impressions)
}
