// Package main provides ...
package main

import "time"

func voidA(a []byte) {

}

func main() {
	// a := make([]byte, 0, len(1024))
	// a := make([]byte, 0, 1024)
	// voidA(a)
	println("works")
	time.Sleep(time.Second * 10)
}
