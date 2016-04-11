// Package main provides ...
package main

import (
	"fmt"
	"time"
)

func main() {
	exit := make(chan bool)
	random := make(chan string)
	go func() {
		for {
			select {
			case <-exit:
				fmt.Printf("  exiting: \n")
				return
			case r := <-random:
				fmt.Printf("  r: %+v\n", r)
			}
		}
	}()

	ticks := time.Tick(time.Second * 1)
	for v := range ticks {
		random <- v.String()
	}
}
