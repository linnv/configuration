// Package main provides ...
package demo

import "testing"

func TestJustDemo(t *testing.T) {
	InnerValidDemo()
	// JustDemo()
	// r := ReturnNil()
	// if r != nil {
	// 	fmt.Printf("  all nil\n")
	// 	return
	// }
	// fmt.Printf("  r: %+v\n", r)

	// n := error(nil)
	// m := 5
	// var n interface{}
	// n = m
	// // n := error(errors.New("good"))
	// fmt.Printf("  reflect.TypeOf(n): %+v\n", reflect.TypeOf(n).Kind())
	// // fmt.Printf("  reflect.TypeOf(n): %+v\n", reflect.ValueOf(n).Kind())
	// fmt.Printf("  reflect.TypeOf(n): %+v\n", reflect.ValueOf(n).Type())
	// fmt.Printf("  reflect.TypeOf(n): %+v\n", reflect.ValueOf(n).Kind())

	// if n != nil {
	// 	fmt.Println("n is not nil")
	// 	return
	// }
	// fmt.Printf("  n: %+v\n", n)
	// var x float64 = 3.4
	// v := reflect.ValueOf(x)
	// fmt.Println("type:", v.Type())
	// fmt.Printf("  v.kind(): %+v\n", v.Kind())
	// // fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	// fmt.Println("value:", v.Float())

}
