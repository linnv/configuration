// Package main provides ...
package newDir

import "testing"

func TestJustDemo(t *testing.T) {
	JustDemo()
	// var m Moniter
	// m = Inheritan{}
	// r := m.GetConsumption()
	//@toDelete
	// fmt.Printf("r: %+v\n", r)

	// var ters Timer
	// ters = &TimeInstanceA{Count: 99}
	// ters.GetDuration()

	// Count of TimeInstanceA in TimeInstanceInheritanceA will be initialise with default value 0 if TimeInstanceA in TimeInstanceInheritanceA is a value member,
	// but if TimeInstanceA in TimeInstanceInheritanceA is a pointer member, you must initialise it manually, or `invalid memory address or nil pointer dereference` will occurs
	// ters = &TimeInstanceInheritanceA{}
	// ters.GetDuration()
	// ters.GetDuration()
}

// func TestJustDemo(t *testing.T) {
// 	fmt.Printf("aSlic: %+v\n", ReturnConst())
// 	// r := ReturnInterface()
// 	r := ReturnSlice()
// 	fmt.Printf("reflect.TypeOf(r): %+v\n", reflect.TypeOf(r))
// 	fmt.Printf("reflect.ValueOf(r): %+v\n", reflect.ValueOf(r))
// 	fmt.Printf("r: %+v\n", r)
// 	fmt.Printf("r: %+v\n", reflect.TypeOf(r))
// 	fmt.Printf("r: %+v\n", reflect.TypeOf(r).Kind())
// 	// fmt.Printf("r: %d\n", reflect.TypeOf(r).Kind())
// 	fmt.Printf("r: %+v\n", reflect.ValueOf(r).Kind())
// 	fmt.Printf("r: %+v\n", reflect.ValueOf(r).Type())
// 	ir := reflect.ValueOf(r).Interface()
// 	// ir, ok := ir.(*MyTime)
// 	// if !ok {
// 	// 	fmt.Printf("error: %+v\n")
// 	// }
// 	// cr := reflect.ValueOf(r).Interface()
// 	// fmt.Printf("ir: %+v\n", ir)
// 	switch t := ir.(type) {
// 	case interface{}:
// 		fmt.Printf("type interface %v\n", t)
// 	}
// }
