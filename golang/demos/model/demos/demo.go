package demos

import (
	"demo/model"
)

type A struct {
	N int
}

func (this A) Demo() {
	println("in model!!!! a's demo'", this.N)
	fmt.Printf("model.VariableA: %+v\n", model.VariableA)
}

func (this *A) DemoP() {
	// func (this *A) Demo() {
	println("a's pointer demo'", this.N)
}
