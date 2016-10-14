package main

import (
	"fmt"
	"strconv"
	"time"
)

func Receive(data <-chan string, done chan struct{}) {
	var t string
	for {
		select {
		case t = <-data:
			println("Receive ", t)
			time.Sleep(time.Second)
		case <-done:
			println("exit ")
			return
		}
	}
}

func main() {
	item := make(chan string, 1)
	done := make(chan struct{})

	go Receive(item, done)

	for i := 0; i < 10; i++ {
		item <- strconv.Itoa(i)
	}

	//issue:  item may lose being received,when done sends data
	time.Sleep(time.Second)
	done <- struct{}{}
	// close(done) //don't it, if you can't ensure that main() goroutine will exit after all of other goroutines exiting
	fmt.Println("done")
}
