package demos

type A struct {
	N int
}

func (this A) Demo() {
	// func (this *A) Demo() {
	println("in models a's demo'", this.N)
}

func (this *A) DemoP() {
	// func (this *A) Demo() {
	println("a's pointer demo'", this.N)
}
