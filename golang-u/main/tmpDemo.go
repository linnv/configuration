// Package main provides ...
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func strReplaceDemo() {
	println("<<<strReplaceDemo start---------------------------")

	r := strings.NewReplacer(
		"{USER_ID}", strconv.Itoa(cu.User.UserId),
		"{PRODUCT_ID}", c.ProductId,
		"{DOMAIN}", this.Ctx.Input.Host(),
	)
	info.LogoutUrl := r.Replace(info.LogoutUrl)
	println("-----------------------------strReplaceDemo end>>>")
	return
}

func Abs(x float64) float64
func main() {
	a := 2.1
	r := Abs(a)
	fmt.Printf("r: %+v\n", r)
	fmt.Printf(": %+v\n", 2/1)
	fmt.Printf(": %+v\n", 2%1)
	// strconv.Itoa(adPId)
	// var a interface{}
	// a = 22
	// b := a.(int)
	//
	// fmt.Printf("b: %+v\n", reflect.ValueOf(b).Kind())
	// fmt.Printf("a: %+v\n", reflect.ValueOf(a).Kind())
	//
	// c := reflect.ValueOf(b).Interface()
	// fmt.Printf("b: %+v\n", reflect.ValueOf(b).Type())
	// fmt.Printf("a: %+v\n", reflect.ValueOf(a).Type())
	// fmt.Printf("c: %+v\n", c)
	// fmt.Printf("c: %+v\n", reflect.ValueOf(c).Kind())

	// b := 1
	// fmt.Printf("b<<10: %+v\n", b<<10) //1* 2^0<<10
	// // c := interface{}(b).(int)
	// var d interface{}
	// d = 3
	// t := d.(int)
	// fmt.Printf("t: %+v\n", t)
	// tt := d.(type)
	// fmt.Printf("d.(type): %+v\n", tt)
	// fmt.Printf("d.(int): %+v\n", d.(type))
}
