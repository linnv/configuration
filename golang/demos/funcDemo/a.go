// Package main provides ...
package newDir

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

type A struct {
	A string `json:"A"`
}

// func ReturnNil() error {
// func ReturnNil() *int {
func ReturnNil() *A {
	// return error(nil)
	// var a A = nil
	var a *A = nil
	return a
}

func retDemo()( r ...int){
println("//<<-------------------------retDemo start-----------")
r=[]int{1,2,}
println("//---------------------------retDemo end----------->>")
	return 
}
