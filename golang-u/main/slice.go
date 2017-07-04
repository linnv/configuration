package main

import (
	"fmt"
	"sort"
	"time"
)

func makeStructSliceDemo() {
	println("<<<makeDemo start---------------------------")
	m := make([]A, 0, 1)
	fmt.Printf("cap(m): %+v\n", cap(m))
	fmt.Printf("len(m): %+v\n", len(m))
	m = append(m, A{})
	fmt.Printf("cap(m): %+v\n", cap(m))
	fmt.Printf("len(m): %+v\n", len(m))
	m = append(m, A{})
	fmt.Printf("cap(m): %+v\n", cap(m))
	fmt.Printf("len(m): %+v\n", len(m))
	println("-----------------------------makeDemo end>>>")
	return
}

func makeBasicDemo() {
	println("<<<makeBasicDemo start---------------------------")
	m := make([]int, 1) //equals to make([]int ,1,1)  /danger
	// m := make([]int, 0, 1)
	fmt.Printf("cap(m): %+v\n", cap(m))
	fmt.Printf("len(m): %+v\n", len(m))
	for i := 0; i < 5; i++ {
		println("append in", i)
		m = append(m, i)
		fmt.Printf("len(m): %+v\n", len(m))
		fmt.Printf("cap(m): %+v\n", cap(m))
	}

	println("-----------------------------makeBasicDemo end>>>")
	return
}

func SlicePointerDemo() {
	println("<<<SlicePointerDemo---------------------------")

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	p := &data
	p[5] = 50
	println(p[5])
	println(data[5])
	println("-----------------------------SlicePointerDemo>>>")

}

func mergeSliceDemo() {
	println("<<<merSliceDemo---------------------------")

	//illegal
	// dataOne := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// dataTwo := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	//legal
	dataOne := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	dataTwo := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var tmpData []int
	tmpData = append(tmpData, dataOne...)
	tmpData = append(tmpData, dataTwo...)
	fmt.Printf("tmpData: %+v\n", tmpData)
	println("-----------------------------merSliceDemo>>>")
}
func parameterDemo(p *[]int) {
	println("<<<parameterDemo---------------------------")
	fmt.Printf("address: %p\n", p)
	for k, v := range *p {
		fmt.Printf("%+v: %+v\n", k, v)
	}
	println("-----------------------------parameterDemo>>>")
	return
}

func pointerParameterDemo(p []int) {
	println("<<<parameterDemo-pointer--------------------------")
	fmt.Printf("address: %p\n", p)
	for k, v := range p {
		fmt.Printf("%+v: %+v\n", k, v)
	}
	p[0] = 2222 // midification will effect p parameter
	println("-----------------------------parameterDemo>>>")
	return
}

func rangeDemo() {
	println("<<<rangeDemo---------------------------")

	dataOne := []interface{}{0, 'a', "b", 1, 3, 2, 4545, 335, 6, 7, 8, 9}
	for k, v := range dataOne {
		fmt.Printf("%+v: %+v\n", k, v)
	}
	println("-----------------------------rangeDemo>>>")
	return
}

type A struct {
	N int
}

func getNill() *A {
	return nil
}

func AnnounceDemo() {
	println("<<<AnnounceDemo---------------------------")
	// all results  are same
	// var a []int
	// a := make([]int, 0, 1)   // but it maybe different between(0,1) and (0)  depent on make(type)
	// a := make([]int, 0)
	a := make([]int, 0, 10) //index out of range
	// a := make([]int, 0, 0)
	for i := 0; i < 10; i++ {
		if i == 4 {
			continue
		}
		// a[i] = i * 10
		a = append(a, i*10)
	}
	for k, v := range a {
		fmt.Printf("cap(a): %+v\n", cap(a))
		fmt.Printf("len(a): %+v\n", len(a))
		fmt.Printf("%+v: %+v\n", k, v)
	}
	println("-----------------------------AnnounceDemo>>>")
	v := 2
	m := getNill()
	if v == 1 || m.N > 0 || m == nil {
		println("nil here")
	}
}

func byteInitDemo() {
	// var BL_DIV_ELEM = byte('=') //词分隔符
	println("<<<byteDemo---------------------------")
	// var BL_ELEM_FAID = append([]byte(`faid`), BL_DIV_ELEM)
	// fmt.Printf("BL_ELEM_FAID: %+v\n", string(BL_ELEM_FAID))
	if []byte("aaa")[0] == []byte("ccc")[0] {
		fmt.Println([]byte("aaa")[0])
		fmt.Println([]byte("ccc")[0])
		fmt.Println("ok")
	}
	println("-----------------------------byteDemo>>>")
	return
}

func byteLoopDemo() {
	println("<<<byteLoopDemo---------------------------")
	s := "a1b2c3"
	b := []byte(s)
	for k, v := range b {
		fmt.Printf("%+v: %+v\n", k, string(v))
		// fmt.Printf("%+v: %+v\n", k, v)
	}
	println("-----------------------------byteLoopDemo>>>")
	return
}

func arrayCopyDemo() {
	println("<<<arrayCopyDemo---------------------------")

	arint := make([]int, 4, 8)
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	copy(arint, data)        //only 4 memory space is available
	arint = append(arint, 5) //5th memory space will be create
	fmt.Printf("arint: %+v\n", arint)
	println("-----------------------------arrayCopyDemo>>>")
	return
}

func UniqueArrayDemo() {
	println("<<<UniqueArrayDemo---------------------------")

	// data := []int{0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var data []int
	data = []int{}
	for i := 0; i < 10000; i++ {
		data = append(data, i)
	}
	for i := 5000; i < 10000; i++ {
		data = append(data, i)
	}

	tmp := make([]int, 0, len(data))
	duplicate := make([]int, 0, len(data))
	start := time.Now()

	for i := 0; i < len(data); i++ {
	loopAgain: //better than put it out ot of for loopAgain
		for j := 0; j < len(tmp); j++ {
			if tmp[j] == data[i] {
				duplicate = append(duplicate, data[i])
				if i < len(data)-1 {
					i++
				}
				continue loopAgain
			}
		}
		tmp = append(tmp, data[i])
	}
	fmt.Printf("duplicate array: %+v\n", len(duplicate))
	fmt.Printf("tmp: %+v\n", len(tmp))
	println("-----------------------------UniqueArrayDemo>>>")
	done := time.Now()
	fmt.Printf("time consumption: %v\n", done.Sub(start).Seconds())

	return
}

func fileterArrayDemo(s, t []int) (duplicate []int, newElement []int) {
	println("<<<fileterDuplicateArrayDemo---------------------------")
	duplicate = make([]int, 0, len(t))
	newElement = make([]int, 0, len(t))
	for i := 0; i < len(t); i++ {
	loopA:
		for j := 0; j < len(s); j++ {
			if s[j] == t[i] {
				duplicate = append(duplicate, t[i])
				if i < len(t)-1 {
					i++
				}
				continue loopA
			}
		}
		newElement = append(newElement, t[i])
	}

	fmt.Printf("duplicate: %+v\n", duplicate)
	fmt.Printf("newElement: %+v\n", newElement)

	println("-----------------------------fileterDuplicateArrayDemo>>>")
	return
}

func arrayMapDemo() {
	println("<<<arrayMapDemo---------------------------")
	arint := make(map[int]bool, 9)
	// data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// copy(arint, data)        //only 4 memory space is available
	// arint = append(arint, 5) //5th memory space will be create
	for i := 0; i < 3; i++ {
		// if i/2 == 0 {
		// 	arint[i] = true
		// }
		arint[i] = true
	}
	fmt.Printf("arint: %+v\n", arint)
	println("-----------------------------arrayMapDemo>>>")
	return
}

func arrayLenDemo() {
	println("<<<arrayDemo---------------------------")
	arint := make([]int, 10)
	for i := 0; i < len(arint); i++ {
		fmt.Printf("i%d: %+v\n", i, arint[i])
	}
	fmt.Printf("len(arint)%d,cap(arint): %+v\n", len(arint), cap(arint))
	ar := arint[5:5]
	fmt.Printf("len(ar)%d,cap(ar): %+v\n", len(ar), cap(ar))
	println("-----------------------------arrayDemo>>>")

	return
}

func makeArrayDemo() {
	println("<<<makeArrayDemo---------------------------")
	const length = 20
	ar := new([length]int)[0:10] //these two are equal
	// ar := make([]int, 10, 20)

	for i := 0; i < 2*length; i++ {
		// if i > 9 {
		if i == 9 {
			fmt.Printf("len(ar)%d,cap(ar): %+v\n", len(ar), cap(ar))
			println("extenting underlying array")
			ar = ar[:cap(ar)]
			fmt.Printf("len(ar)%d,cap(ar): %+v\n", len(ar), cap(ar))
		}
		if i == length {
			fmt.Printf("aa: works\n")
			// tmpl := cap(ar) * 2
			_ar := new([40]int)[:cap(ar)]
			// println(_ar)
			for k, v := range ar {
				_ar[k] = v
			}
			// _ar = _ar[:cap(_ar)]
			fmt.Printf("_ar: %+v\n", _ar)
			fmt.Printf("len(ar)%d,cap(ar): %+v\n", len(_ar), cap(_ar))
			ar = _ar
			fmt.Printf("len(ar)%d,cap(ar): %+v\n", len(ar), cap(ar))
			//@TODO
			// break
			// ar = _ar[:cap(_ar)]
		}
		ar[i] = 9 * i
	}
	for k, v := range ar {
		fmt.Printf("%+v: %+v\n", k, v)
	}
	fmt.Printf("len(ar)%d,cap(ar): %+v\n", len(ar), cap(ar))
	println("-----------------------------makeArrayDemo>>>")
	return
}

// func main() {
//
// 	t := make([]int, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
// 	// for i := range s {
// 	// 	t[i] = s[i]
// 	// }
// 	// s = t
// 	// fmt.Printf("cap(s): %+v\n", cap(s))
// 	// fmt.Printf("len(s): %+v\n", len(s))
// 	ret := copy(t, s)
// 	s = t
// 	fmt.Printf("cap(s): %+v\n", cap(s))
// 	fmt.Printf("len(s): %+v\n", len(s))
// 	fmt.Printf("ret: %+v\n", ret)
//
// 	// d := data[4:] //[from(include):end(not_include)]
// 	// fmt.Printf("s: %+v\n", s)
// 	// fmt.Printf("d: %+v\n", d)
// 	// copy(s, d) //no more than what s can contain
// 	// s[2] = 100
// 	// fmt.Printf("s: %+v\n", s)
// 	// copys := make([]int, len(s))
// 	// fmt.Printf("len(copys): %+v\n", len(copys))
// 	// fmt.Printf("cap(copys): %+v\n", cap(copys))
// 	// copy(copys, s) //copys's length must not be zero or data will not be copied
// 	//
// 	// fmt.Printf("copy s: %+v\n", s)
// 	// fmt.Printf("copys: %+v\n", copys)
// 	ArrayDemo()
// 	n, m := RetDemo()
// 	fmt.Printf("n,m: %di, %d\n", n, m)
// }

func Ret2Demo() (a, b int) {
	println("<<<-----------------Ret2Demo----------")
	a = 2
	b = 3
	println("-----------------Ret2Demo------------>>>")
	return
}

func RetDemo() (i, j int) {
	i = 0
	j = 1
	println("<<<-----------------RetDemo----------")
	println("-----------------RetDemo------------>>>")
	return Ret2Demo()
}

func ArrayDemo() {
	println("<<<-----------------ArrayDemo----------")
	// var br [][]int
	// ar := []int{1, 2, 32, 4, 3}
	// tmp, ok := ar[8]
	// if ok {
	// 	fmt.Printf("tmp: %+v\n", tmp)
	// }
	// br = [][]int{ar}
	// fmt.Printf("len(br): %+v\n", len(br))
	// var ar [][2]int = [][2]int{{2, 23}, {323, 434}, {32}}
	// fmt.Printf("len(ar): %+v\n", len(ar))
	println("-----------------ArrayDemo------------>>>")
}

func initialSliceDemo() {
	println("<<<initialSliceDemo---------------------------")
	NetWorkType := []int{2, 333, 4, 44, 5, 9}
	nwt := sort.IntSlice(NetWorkType)
	nwt.Sort()
	fmt.Printf("nwt: %+v\n", nwt)
	println("-----------------------------initialSliceDemo>>>")
	return
}

func main() {
	makeBasicDemo()
	// makeStructSliceDemo()
	// initialSliceDemo()
	// makeArrayDemo()
	// arrayLenDemo()

	// s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// t := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 3342, 22222, 39}
	// var d, n []int
	// d, n = fileterArrayDemo(s, t)
	// fmt.Printf("d %+v,n: %+v\n", d, n)

	// d := []int{1, 23}
	// pointerParameterDemo(d)
	// fmt.Printf("d: %+v\n", d)
	// UniqueArrayDemo()
	// ArrayDemo()
	// AnnounceDemo()
	// rangeDemo()
	// arrayMapDemo()
	// arrayDemo()
	// arrayCopyDemo()
	// byteInitDemo()
	// byteLoopDemo()

	// data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// parameterDemo(&data)
	// pointerParameterDemo(data)
	// mergeSliceDemo()
	// // SlicePointerDemo()
	//
	// data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// for k, v := range data {
	// 	fmt.Printf("xx%+v: %+v\n", k, v)
	// }
	// s := data[:3] //s'cap equal to data, but len is different
	// fmt.Printf("s: %+v\n", s)
	// fmt.Printf("data: %+v\n", data)
	// s[1] = 111
	// fmt.Printf("s: %+v\n", s)
	// fmt.Printf("data: %+v\n", data)
	// fmt.Printf("cap(data): %+v\n", cap(data))
	// fmt.Printf("len(data): %+v\n", len(data))
	// // s = s[:cap(s)]
	// fmt.Printf("cap(s): %+v\n", cap(s))
	// fmt.Printf("len(s): %+v\n", len(s))
}
