package main

import "demo/models/demos"

func main() {
	a := demos.A{N: 10}
	// 	demos.A(a).Demo() // these two are equal
	// 	a.Demo()

	// demos.A(&a).DemoP() illegal DemoP() must not a pointer function
}
