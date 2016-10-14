// Package main provides ...
//@TODO  _, ok := spotInfoSetMap[adk]   vs v, ok := spotInfoSetMap[adk]    when updating v
package main

import "fmt"

var aMap map[int]string

type Mmap map[int]string
type Mbool map[int]bool
type WM struct {
	mm  Mmap
	mid Mbool
}

func wmDemo() {
	println("<<<wmDemo---------------------------")

	var tm *WM
	tm = new(WM)
	tm.mm, tm.mid = make(Mmap, 100), make(Mbool, 100)
	// tm.mm, tm.mid = make(Mmap,0, 1), make(Mbool, 10)
	for i := 0; i < 14; i++ {
		tm.mm[i] = "i"
	}
	fmt.Printf("tm.mm: %+v\n", tm.mm)
	println("-----------------------------wmDemo>>>")
	return
}

// func copyMap(v []map[int]string) {
//
// 	tmp := make([]map[int]string)
// 	tmp = v
// 	fmt.Printf("tmp: %+v\n", tmp)
//
// }
func rangeDemo() {
	println("<<<rangeDemo---------------------------")
	n := map[int]string{10: "foo", 2: "bar"}
	for k, v := range n {
		if k == 2 {
			continue
		}
		fmt.Printf("map k v :%+v: %+v\n", k, v)
	}
	println("-----------------------------rangeDemo>>>")
	return
}

func makeMapDemo() {
	println("<<<makeMapDemo---------------------------")
	m := make(map[int]int, 2)
	for i := 0; i < 10; i++ {
		fmt.Printf("len(m): %+v\n", len(m))
		m[i] = i * 10
	}

	for k, v := range m {
		fmt.Printf("%+v: %+v\n", k, v)
	}
	println("-----------------------------makeMapDemo>>>")
	return
}

func copyMapDemo() {
	println("<<<copyMapDemo---------------------------")

	m := make(map[int]int, 2)
	for i := 0; i < 10; i++ {
		fmt.Printf("len(m): %+v\n", len(m))
		m[i] = i * 10
	}
	n := make(map[int]int, 10)
	n = m
	fmt.Printf("n: %+v\n", n)
	println("-----------------------------copyMapDemo>>>")
	return
}

type data struct {
	name string
}

func main() {
	m := map[string]*data{"x": {"one"}}
	fmt.Printf("m: %+v\n", m)
	m["z"] = new(data)
	//vs no new() and get new key directly
	m["z"].name = "what?" //???
	fmt.Printf("m: %+v\n", m)

	// aInt := make([]string, 1024*102)
	// aInt := make([]string, 1024*1024*1024)
	// fmt.Printf("aInt: %+v\n", aInt)
	// aMap = make(map[int]string, 1024*1024*1024)
	// time.Sleep(time.Second * 10)
	// for i := 0; i < 1024*1024*1024; i++ {
	// 	aMap[i] = "a"
	// }
	// a := make(map[string]int, 9)
	// // a := make(map[string]int, 0)
	// a["dd"] = 1
	// a["de"] = 1
	// fmt.Printf("a: %+v\n", a)
	// fmt.Printf("len(a): %+v\n", len(a))

	// if _, ok := a["dd"]; ok {
	// 	fmt.Printf(": works\n")
	// }
	// if _, ok := a["dxxd"]; ok {
	// 	fmt.Printf(": works\n")
	// }

	// copyMapDemo()
	// makeMapDemo()
	// rangeDemo()
	// wmDemo()
	// if 2 > 1 {
	// 	var m []int
	// 	m = make([]int, 10)
	// }
	// fmt.Printf("cap(m): %+v\n", cap(m))
}
