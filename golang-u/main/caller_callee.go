// Package main provides ...
package main

func CallerDemo() {
	defer func() {
		println("<<<CallerDemo defer ---------------------------")

	}()
	println("<<<CallerDemo---------------------------")

	CalleeDemo()
	println("-----------------------------CallerDemo>>>")
	return
}

func CalleeDemo() {
	defer func() {
		println("<<<CalleeDemo defer ---------------------------")

	}()
	println("<<<CalleeDemo---------------------------")

	println("-----------------------------CalleeDemo>>>")
	return
}
func main() {
	CallerDemo()

}
