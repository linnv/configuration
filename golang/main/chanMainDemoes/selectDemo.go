// Package main provides ...
package main

import (
	"fmt"
	"time"
)

var i = 0

func F() {
	i++
}

func A(done chan struct{}) {
	count := 0
	for {
		count++
		println("after ----", count)
		println(count)
		F()
		select {
		case <-done:
			println("exit goroutine")
			return
		default:
			println("default")
		}
	}
}

func main() {
	done := make(chan struct{}, 1)
	go A(done)
	time.Sleep(time.Second)
	close(done)
	time.Sleep(time.Second)
	fmt.Printf("i: %+v\n", i)
	println("done")
}
