// Package main provides ...
package demo

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

func f() int {
	buf := make([]byte, 1024)
	return len(buf)
}
func f2() int {
	var bufArray [1024]byte
	buf := bufArray[:]
	return len(buf)
}
func f3() {
	var bufArray [1024]byte
	_ = bufArray[:]
}
