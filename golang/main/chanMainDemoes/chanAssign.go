// Package main provides ...
package main

import (
	"fmt"
	"runtime"
	"time"
)

func RangeChanT(c chan int) {
	for v := range c {
		// if v >= 2 {
		// 	fmt.Printf("  close c2 \n")
		// 	close(c)
		// }
		fmt.Printf("t  v: %+v\n", v)
	}
}
func RangeChan(c chan int) {
	for v := range c {
		if v >= 2 {
			fmt.Printf("  close c1 \n")
			close(c)
		}
		fmt.Printf("  v: %+v\n", v)
	}
}

func GenerateChanT(c chan<- int) {
	for i := 11; i < 30; i++ {
		c <- i
	}
}
func GenerateChan(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	c1 := make(chan int, 10)
	c2 := make(chan int)

	c2 = c1
	// go GenerateChan(c1)
	go GenerateChanT(c2)

	// fmt.Printf("  <-c1: %+v\n", <-c1)
	// fmt.Printf("  <-c2: %+v\n", <-c2)
	// fmt.Printf("  <-c2: %+v\n", <-c2)
	go RangeChanT(c2)
	go RangeChan(c1)
	time.Sleep(time.Second * 1)
}
