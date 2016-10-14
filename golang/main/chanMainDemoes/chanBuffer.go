// Package main provides ...
package main

import (
	"fmt"
	"time"
)

func main() {
	ss := make(chan string, 2)
	b := make(chan bool)
	go func() {
		for {
			select {
			case msg := <-ss:

				fmt.Printf("msg: %+v\n", msg)
				if msg == "two" {
					b <- true
					return
				}
			}
		}
	}()
	ss <- "one"
	time.Sleep(time.Second * 5)
	ss <- "two"
	<-b
	return
}
