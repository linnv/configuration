package demo

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/linnv/logx"
)

func ConstDemo() {
	println("//<<-------------------------ConstDemo start-----------")
	start := time.Now()

	const (
		ClickIndex  = iota*2 + 1 // 0
		UrlIndex                 // 2
		TrackIndex               // 4
		WidthIndex               // 6
		HeightIndex              // 8
	)

	logx.Debug("UrlIndex: %+v\n", UrlIndex)
	logx.Debug("WidthIndex: %+v\n", WidthIndex)
	logx.Debug("HeightIndex: %+v\n", HeightIndex)
	logx.Debug("TrackIndex: %+v\n", TrackIndex)
	logx.Debug("ClickIndex: %+v\n", ClickIndex)

	var ls = []string{
		"one",
	}
	logx.Debug("ls[UrlIndex]: %+v\n", ls[UrlIndex])

	fmt.Printf("ConstDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------ConstDemo end----------->>")
}

func FloatConvertionDemo() {
	println("//<<-------------------------FloatConvertionDemo start-----------")
	start := time.Now()
	// a := 1.22222222
	a := 0.0006
	r := fmt.Sprintf("%.3f", a)
	log.Printf("r: %s\n", r)
	fmt.Printf("FloatConvertionDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------FloatConvertionDemo end----------->>")
}

func ReturnDemo() (r int) {
	println("//<<-------------------------ReturnDemo start-----------")
	start := time.Now()
	r = 1
	fmt.Printf("ReturnDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------ReturnDemo end----------->>")
	return 4
}

func pointerErroroDemo(err *error) {
	println("//<<-------------------------pointerErroroDemo start-----------")
	start := time.Now()
	*err = errors.New("inner")
	fmt.Printf("pointerErroroDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------pointerErroroDemo end----------->>")
}

//doesn't work
func GotoDemo() {
	println("//<<-------------------------GotoDemo start-----------")
	start := time.Now()
	valid := func(s string) {
		println("valid", s)
		// goto End
	}
	valid("debug")
	goto End
	println("should not reach here")
	fmt.Printf("GotoDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))

End:
	allowDeclearOrNot := false
	log.Printf("allowDeclearOrNot: %+v\n", allowDeclearOrNot)
	println("return")

	println("//---------------------------GotoDemo end----------->>")
	return
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	// a := []byte("a")
	// b := []byte("b")
	// c := a + b
	// log.Printf("c: %s\n", c)
	var err error
	// pointerErroroDemo(&err)
	log.Printf("err: %+v\n", err)
	// GotoDemo()
	// r := ReturnDemo()
	// log.Printf("r: %+v\n", r)

	// a := []int{1, 2}
	// var b []int
	// b = a
	// fmt.Printf("a: %+v,cap(a):%d,len(a):%d arrd:%v \n", a, cap(a), len(a), &a[0])
	// fmt.Printf("b: %+v,cap(b):%d,len(b):%d arrd:%v \n", b, cap(b), len(b), &b[0])
	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		if j == 3 {
	// 			break
	// 		}
	// 		fmt.Printf("j: %+v\n", j)
	// 	}
	// 	println()

	// for i := 0; i < 10; i++ {
	// fmt.Printf("{\"normal\",args{%d},%d},\n", i, i%9)
	// print(i, "\t")
	// if i < 10 {
	// 	print("\t")
	// }
	// if i%9 == 0 {
	// 	println()
	// }
	// }
	// UnescapeByteDemo()
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

func strSliceDemo() {
	println("//<<-------------------------strSliceDemo start-----------")
	start := time.Now()
	const s = `false
true
true
true
true
true
true
false
true
true
true
false
true
false
false
true
true
false
true
false
true
false
false
false
true
true
false
true
false
false
true
false
true
false
false
false
true
false
false
false
	`
	b, e := 0, 0
	j := 0
	// ss := make([]string, 0, 32)
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			e = i
			fmt.Printf("{\"normal\", args{%d},%s},\n", j, string(s[b:e]))

			b = e + 1
			j++
			// ss = append(ss, s[b:e])

		}
	}
	// for i := 0; i < len(ss); i++ {
	// 	fmt.Printf("%d: %v\n", i, ss[i])
	// }
	fmt.Printf("strSliceDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------strSliceDemo end----------->>")
}

func endTypeDemo() {
	println("//<<-------------------------endTypeDemo start-----------")
	start := time.Now()
	var sessionId uint64
	// sessionId = 257
	// sessionId = 256
	sessionId = 256*2 + 7
	b := make([]byte, 8)
	// highe bit is locate at low(little) end
	binary.LittleEndian.PutUint64(b, sessionId)
	log.Printf("little b: %+v\n", b)
	binary.BigEndian.PutUint64(b, sessionId)
	log.Printf("big b: %+v\n", b)
	// log.Printf("[]byte(sessionId): %+v\n", []byte(sessionId))
	fmt.Printf("endTypeDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------endTypeDemo end----------->>")
}

func ScanDemo() {
	println("//<<-------------------------ScanDemo start-----------")
	start := time.Now()
	var s string
	fmt.Scanln(&s)
	log.Printf("s: %+v\n", s)
	fmt.Printf("ScanDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------ScanDemo end----------->>")
}

func forFuncDemo() {
	println("//<<-------------------------forFuncDemo start-----------")
	start := time.Now()
	get := func() int {
		println("invoking me?")
		return 10
	}

	// for i := 0; i < get(); i++ { don't use this loop, because get() will be invoked every loop
	length := get()
	for i := 0; i < length; i++ {
		println(i)
	}
	fmt.Printf("forFuncDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------forFuncDemo end----------->>")
}

func PathSeperatorDemo() {
	println("//<<-------------------------PathSeperatorDemo start-----------")
	start := time.Now()
	s := "/"
	// s := "\\"
	is := os.IsPathSeparator(s[0])
	log.Printf("is: %+v\n", is)
	fmt.Printf("PathSeperatorDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------PathSeperatorDemo end----------->>")
}

func CaseDemo() {
	println("//<<-------------------------CaseDemo start-----------")
	start := time.Now()
	i := 8
	switch {
	case i > 10:
		println("great than 10")
	case i < 9:
		println("do nothing")
	case i > 1:
		println("1-9")
	}
	fmt.Printf("CaseDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------CaseDemo end----------->>")
}
