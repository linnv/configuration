// Package main provides ...
package demo

import (
	"encoding/json"
	"fmt"
)

type A struct {
	N string `json:"name"`
	B
}

type AP struct {
	N string `json:"name"`
	*B
}

func (this *A) Afun() {
	println("A func")
}

type B struct {
	BN string `json:"bName"`
}

func (this *B) BFun() {
	println("b func")
}

func MarshalDemoA() (y error) {
	println("<<<MarshalDemo A---------------------------")
	a := &AP{}
	a.N = "a"
	a.B = new(B)
	a.BN = "bn"
	bs, y := json.Marshal(a)
	if y != nil {
		return y
	}
	fmt.Printf("string(bs): %+v\n", string(bs))

	println("-----------------------------MarshalDemo A>>>")
	return
}

func MarshalDemo() (y error) {
	println("<<<MarshalDemo---------------------------")
	a := &A{}
	a.N = "a"
	a.BN = "bn"
	bs, y := json.Marshal(a)
	if y != nil {
		return y
	}

	fmt.Printf("string(bs): %+v\n", string(bs))

	println("-----------------------------MarshalDemo>>>")
	return
}

func MarshalDemoAP() (y error) {
	println("<<<MarshalDemo AP---------------------------")
	a := &AP{}
	a.N = "a"
	//@TODO B is nil error occurs
	a.B = new(B)
	a.BN = "bn"
	bs, y := json.Marshal(a)
	if y != nil {
		return y
	}
	fmt.Printf("string(bs): %+v\n", string(bs))

	println("-----------------------------MarshalDemo AP>>>")
	return
}

func MarshalDemoP() (y error) {
	println("<<<MarshalDemo p---------------------------")
	a := A{}
	a.N = "a"
	a.BN = "bn"
	bs, y := json.Marshal(a)
	if y != nil {
		return y
	}
	fmt.Printf("string(bs): %+v\n", string(bs))

	println("-----------------------------MarshalDemo P>>>")
	return
}

func UnMarshalDemoAP(bs []byte) (y error) {
	a := &AP{}
	y = json.Unmarshal(bs, &a)
	fmt.Printf("unmarshal AP: %+v\n", a)
	fmt.Printf("a.B: %+v\n", a.B)
	return
}

func UnMarshalDemoA(bs []byte) (y error) {
	a := AP{}
	y = json.Unmarshal(bs, &a)
	fmt.Printf("unmarshal A: %+v\n", a)
	fmt.Printf("a.B: %+v\n", a.B)
	return
}
func UnMarshalDemoP(bs []byte) (y error) {
	a := &A{}
	y = json.Unmarshal(bs, &a)
	fmt.Printf("unmarshal P: %+v\n", a)
	fmt.Printf("a.B: %+v\n", a.B)
	return
}

func UnMarshalDemo(bs []byte) (y error) {
	a := A{}
	y = json.Unmarshal(bs, &a)
	fmt.Printf("unmarshal: %+v\n", a)
	return
}

func JustDemo() {
	println("<<<JustDemo---------------------------")
	a := 19
	fmt.Printf("a: %+v\n", a)
	println("-----------------------------JustDemo>>>")
	return
}
