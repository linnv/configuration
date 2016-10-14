// Package main provides ...
package newDir

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"
)

func bitOperation() {
	var bits int
	fmt.Printf("bits: %b\n", bits)
	var count uint = 10
	var i uint = 0
	//initial all bit of int to 1
	for i = 0; i < count; i++ {
		bits |= (1 << i)
	}
	fmt.Printf("bits: %b\n", bits)

	// set the second bit to 0
	bits &= ^(1 << 2)
	bits &= ^(1 << 4)
	fmt.Printf("bits: %b\n", bits)
}

func R() (a int) {
	a = 109
	//defer: first in first out
	defer func(a *int) { *a = 9898 }(&a)
	defer func() { a = 12244 }()
	return 19
}

func ScopeDemo() {
	println("//<<-------------------------ScopeDemo start-----------")
	start := time.Now()
	e := 33
	a := []int{1,
		2,
	}
	fmt.Printf("a: %+v\n", a)
	fmt.Printf("e: %+v %p\n", e, &e)
	strconv.Itoa(3)
	func() {
		// e := 11
		//vs
		// e = 11
		e, f := 11, 33
		fmt.Printf("f: %+v\n", f)
		fmt.Printf("inner function e: %+v %p\n", e, &e)
	}()
	fmt.Printf("e: %+v %p\n", e, &e)
	fmt.Printf("ScopeDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------ScopeDemo end----------->>")
}

type MyData struct {
	One int

	// two string `json:"xxx"`
	//vs
	Two string `json:"xxx"`
}

func CopyPointerDemo() {
	println("//<<-------------------------CopyPointerDemo start-----------")
	start := time.Now()
	a := new(MyData)
	a.One = 333
	b := a
	b.One = 21
	fmt.Printf("a: %+v %p\n", a, a)
	fmt.Printf("b: %+v %p\n", b, b)
	fmt.Printf("CopyPointerDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------CopyPointerDemo end----------->>")
}

func ErrShadowDemo() (err error) {
	println("//<<-------------------------ErrShadowDemo start-----------")
	start := time.Now()
	a, err := 1, fmt.Errorf("feefe")
	if a > 1 {
		var b int
		b, err = 1, fmt.Errorf("feefe")
		// b, err := 1, fmt.Errorf("feefe")
		fmt.Printf("b: %+v\n", b)
	}
	fmt.Printf("a: %+v\n", a)
	fmt.Printf("ErrShadowDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------ErrShadowDemo end----------->>")
	err = fmt.Errorf("xxx")
	return nil
}

func caseDemo() {
	println("//<<-------------------------caseDemo start-----------")
	start := time.Now()
	for i := 0; i < 10; i++ {
		switch i {
		case 3:
			continue
		case 2:
			println("good")
			goto share
		}
		// println(i)
	}
share:
	println("share")

	println("ggo")
	fmt.Printf("caseDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------caseDemo end----------->>")
}

func ReplaceDemo() {
	println("//<<-------------------------ReplaceDemo start-----------")
	start := time.Now()
	js := `
		{PRODUCT_NAME} is  on the stage
		PRODUCT_NAME is  on the stage
		`
	replacer := strings.NewReplacer(
		"{PRODUCT_NAME}", "xxxjialn's product",
		"PRODUCT_NAME", "jialn's product",
		"{PRODUCT_ICON_NAME}", "iconName",
	)
	js = replacer.Replace(js)
	fmt.Printf("js: %+v\n", js)
	fmt.Printf("ReplaceDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------ReplaceDemo end----------->>")
}
func loopDemo() {
	println("//<<-------------------------loopDemo start-----------")
	start := time.Now()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 {
				break
			}
			println(i)
		}
	}
	fmt.Printf("loopDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------loopDemo end----------->>")
}

const (
	C_READ = 1 << iota
	C_MODIFY
	C_MODDEL
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	i := 3
	fmt.Printf("i: %+v\n", i)
	a := make(map[int]string)
	b := a[3]
	fmt.Printf("b: %+v\n", b)
	fmt.Printf("C_READ: %+v\n", C_READ)
	fmt.Printf("C_MODIFY: %+v\n", C_MODIFY)
	fmt.Printf("C_MODDEL	: %+v\n", C_MODDEL)
	// t, err := time.Parse("2006-01-01", "2006-01-02")
	// fmt.Println(t)
	// fmt.Println(err)
	// ReplaceDemo()
	// gotoDemo()
	// a := make([]int, 4)
	// for i := 0; i < 4; i++ {
	// 	a[i] = i
	// }
	// b := make([]int, 3, 4)
	//
	// // dst src
	// copy(b, a)
	// fmt.Printf("b: %+v\n", b)
	// fmt.Printf("a: %+v\n", a)
	// fmt.Printf("2.3/1: %+v\n", 2.3/1)
	// err := os.Remove("/Users/Jialin/golang/src/demos/demos/basicTypeDemo/b")
	// err := ErrShadowDemo()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Printf("err.Error(): %+v\n", err.Error())
	// err := fmt.Errorf("文件格式不在允许的上传范围")
	// fmt.Printf("err.Error(): %+v\n", err.Error())
	// caseDemo()
	// println("good")

	// j := uint(1)
	// j := int(1)
	// fmt.Printf("j: %b\n", j)
	// fmt.Printf("^j: %b\n", ^j)

	// in := MyData{1, "two"}
	// fmt.Printf("%#v\n", in) //prints main.MyData{One:1, two:"two"}
	//
	// encoded, _ := json.Marshal(in)
	// fmt.Println(string(encoded)) //prints {"One":1}

	// bitOperation()
	// ch := make(chan int)
	// done := make(chan struct{})
	// for i := 0; i < 3; i++ {
	// 	go func(idx int) {
	// 		select {
	// 		case ch <- idx:
	// 			fmt.Println(idx, "sent result")
	// 		case <-done:
	// 			fmt.Println(idx, "exiting")
	// 		}
	// 	}(i)
	// }
	//
	// //get first result
	// fmt.Println("result:", <-ch)
	// close(done)
	// fmt.Println("result:", <-ch)
	// //do other work
	// time.Sleep(1 * time.Second)
	//
	// s := "a"
	// // s := "fejfejejeej"
	// fmt.Printf("len(s): %+v\n", len(s))
	// fmt.Printf("s[0]: %s\n", string(s[0]))

	// b := R()
	// fmt.Printf("b: %+v\n", b)

	// str := "，,"
	// str := "，"
	// str := ","

	// str := "．."
	// // str := "。."
	// // str := "，"
	// // str := "."
	// bs := []rune(str)
	// // . 46 ,44
	// for i := 0; i < len(bs); i++ {
	// 	fmt.Printf("bs[i]: %+v\n", bs[i])
	// }

	// Zh2En()
	// a := 1
	// switch a {
	// case 1, 2:
	// 	println("it is 1 or 2")
	// 	fallthrough
	// case 44:
	// 	println("it is 3331 or 233")
	// default:
	// 	println("unkown")
	// }

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
	println("-----------------------------JustDemo end>>>")
	return
}

func algorithmOperation() {
}

func getClicksRatio(impressions, clicks int) float32 {
	return float32(int64(float32(clicks)/float32(impressions)*100)) / 100
	// return float32(clicks) / float32(impressions)
}

// func Zh2En() {
// 	zhs := make(map[rune]string)
// 	zhs[[]rune("。")[0]] = "."
// 	// "０１２３ＡＢＣＤＦＷＳ＼＂，．？＜＞｛｝［］＊＆＾％＃＠！～（）＋－｜：；"
// 	problemStr := "3。5"
// 	slice := []rune(problemStr)
// 	// convert := make([]rune, len(slice))
// 	var convert string
// 	for i := 0; i < len(slice); i++ {
// 		if v, ok := zhs[slice[i]]; ok {
// 			// convert = append(convert, slice[i])
// 			convert += v
// 			continue
// 		}
// 		convert += string(slice[i])
// 		// convert = append(convert, slice[i])
// 		// fmt.Printf("zsh[slice[i]: %+v\n", zhs[slice[i]])
// 	}
// 	fmt.Printf("convert: %+v\n", convert)
// }

func Zh2En() {
	// zhs := make(map[rune]rune)
	// zhs[[]rune("。")[0]] = "."
	gbkStr := "０１２３＼＂，．？＜＞｛｝［］＊＆＾％＃＠！～（）＋－｜：；"
	utf8Str := "0123\\\",.?<>{}[]*&^%#@1~()+-|:;"
	gbkRunes := []rune(gbkStr)
	utf8Runes := []rune(utf8Str)
	// gbkRunesLen := len(gbkRunes)
	// fmt.Printf("len(utf8Str): %+v\n", len(utf8Runes))
	// fmt.Printf("gbkRunes): %+v\n", len(gbkRunes))
	// fmt.Printf("gbkRunes[9]: %s\n", string(gbkRunes[gbkRunesLen-1]))
	// fmt.Printf("utf8Runes[9]: %s\n", string(utf8Runes[gbkRunesLen-1]))
	problemStr := "3．5"
	slice := []rune(problemStr)
	convert := make([]rune, 0, len(slice))
	var tmpRune rune
	for i := 0; i < len(slice); i++ {
		tmpRune = slice[i]
		for j := 0; j < len(gbkRunes); j++ {
			if slice[i] == gbkRunes[j] {
				tmpRune = utf8Runes[j]
				break
			}
		}
		convert = append(convert, tmpRune)
	}
	fmt.Printf("convert: %+v\n", string(convert))
}

func utf8Demo() {
	println("//<<-------------------------utf8Demo start-----------")
	start := time.Now()
	data := "♥2吴［"
	fmt.Printf("utf8.ValidString(data): %+v\n", utf8.ValidString(data))
	fmt.Println(utf8.RuneCountInString(data))
	fmt.Printf("len(data): %+v\n", len(data))
	fmt.Printf("rune(data): %+v\n", []rune(data))
	fmt.Printf("len rune(data): %+v\n", len([]rune(data)))
	// println(len(data))
	fmt.Printf("utf8Demo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------utf8Demo end----------->>")
}

func gotoDemo() {
	println("//<<-------------------------gotoDemo start-----------")
	start := time.Now()
	if float64(1e3*2) > 1000.2 {
		println("less")
	}
	i := 2
	switch i {
	case 1:
		println(1)
		goto end
	case 2:
		println(2)
		fallthrough
	case 3:
		println(3)
		fallthrough
	case 4:
		println(4)
		goto here
	}

here:
	println("here reached")

end:
	println("end")

	fmt.Printf("gotoDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------gotoDemo end----------->>")
}
func bytesOfUnitDemo() {
	println("//<<-------------------------bytesOfUnitDemo start-----------")
	start := time.Now()
	var a [3]uint32
	var i int
	var b byte
	var u16 uint16
	fmt.Println(unsafe.Sizeof(a))    // prints 12
	fmt.Println(unsafe.Sizeof(a[0])) // prints 12
	fmt.Println(unsafe.Sizeof(i))    // prints 12
	fmt.Println(unsafe.Sizeof(b))    // prints 12
	fmt.Println(unsafe.Sizeof(u16))  // prints 12
	type S struct {
		a uint16
		b uint32
	}
	var s S
	// The alignment of a basic type is usually equal to its width, but the alignment of a struct is the maximum alignment of any field, and the alignment of an array is the alignment of the array element. The maximum alignment of any value is therefore the maximum alignment of any basic type. Even on 32-bit systems this is often 8 bytes, because atomic operations on 64-bit values typically require 64-bit alignment.
	fmt.Println(unsafe.Sizeof(s)) // prints 8, not 6
	fmt.Printf("bytesOfUnitDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------bytesOfUnitDemo end----------->>")
}
