package demo

import "testing"

func TestJustDemo(t *testing.T) {
	JustDemo()
	// var m Moniter
	// m = Inheritan{}
	// r := m.GetConsumption()
	// //@toDelete
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

func Test_updateInInterfaceDemo(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{1}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateInInterfaceDemo(tt.args.i); got != tt.want {
				t.Errorf("updateInInterfaceDemo() = %v, want %v", got, tt.want)
			}
		})
	}
}
