package main

import (
	"demo/model"
	d1 "demo/model/demos"
	d2 "demo/models/demos"
	"fmt"

	// "demo/model/demos"
	// "demo/models/demos"
)

func main() {
	// a := demos.A{N: 10}
	// b := demos.A{N: 12}
	// a.Demo()
	fmt.Printf("  fej: %+v\n", fej)
	// b.Demo()

	a := d1.A{N: 10}
	b := d2.A{N: 10}
	a.Demo()
	b.Demo()
	// fmt.Println(a, b)
	// model.VariableA.N = 10
	model.Do()

}
