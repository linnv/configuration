package model

import (
	"demo/model/demos"
	"fmt"
)

var VariableA = demos.A{}

func init() {
	VariableA.N = 12
	println("inital model")
}

func Do() {
	fmt.Printf("model: works\n")
	fmt.Printf("VariableA: %+v\n", VariableA)

}
