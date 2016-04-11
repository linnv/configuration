// Package main provides ...
package main

import (
	"fmt"
)

type A func(text string) []byte

func Wrap(a A) func(string) (s string) {
	return func(plaintText string) string { //返回的函数    这个函数将对a函数的返回结果进行处理
		return fmt.Sprintf("++%s", a(plaintText)) //将自由变量与 a 与 返回的函数 进行闭合
	}
}

func B(s string) []byte {
	s = "demo" + s
	return []byte(s)
}

func SumDemo(in []int, xform func(int) int) (out []int) {
	println("<<<SumDemo---------------------------")
	out = make([]int, len(in))
	for idx, val := range in {
		out[idx] = xform(val)
	}
	println("-----------------------------SumDemo>>>")
	return
}

func main() {
	// fmt.Println()
	// r := Wrap(B)
	// cr := r("mmmm")
	// fmt.Printf("cr: %+v\n", cr)

	data := []int{8, 6, 7, 5, 3, 0, 9}
	total := 0
	fmt.Printf("Result is %v\n", my_transform(data, func(v int) int {
		total += v
		return total
	}))
}
