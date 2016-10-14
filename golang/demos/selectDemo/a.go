// Package main provides ...
package newDir

import (
	"fmt"
	"time"
)

func JustDemo() {
	println("//<<-------------------------JustDemo start-----------")
	start := time.Now()

	fmt.Printf("JustDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------JustDemo end----------->>")
}

func SelectTimeoutDemo() {
	println("//<<-------------------------SelectTimeoutDemo start-----------")
	start := time.Now()
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()
	ch := make(chan int)
	// go func() {
	// 	ch <- 22
	// }()

	select {
	case <-ch:
		// println("xxx")
	case <-timeout:
		println("time out while waiting for channel")
	}

	fmt.Printf("SelectTimeoutDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------SelectTimeoutDemo end----------->>")
}

func SelectIsfullDemo() {
	println("//<<-------------------------SelectIsfullDemo start-----------")
	start := time.Now()
	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
	default:
		fmt.Println("channel is full !")
	}
	fmt.Printf("SelectIsfullDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------SelectIsfullDemo end----------->>")
}

//read from a closed channel is legal, but send to it is panic
func SelectCloseChan() {
	a := make(chan int, 1)
	timeout := time.NewTimer(1e9)
	close(a)
	for {
		select {
		case <-a:
			println("get from closed channel")
			time.Sleep(time.Millisecond * 500)
			timeout.Stop()
		case <-timeout.C:
			println("time out")
			return

		}
	}
}
