// Package main provides ...
package main

import (
	"fmt"
	"os"
	"time"
)

func CloseChanDemo() {
	println("//<<-------------------------CloseChanDemo start-----------")
	start := time.Now()
	ic := make(chan int, 2)
	ic <- 1
	close(ic)
	time.Sleep(1 * time.Second)
	r := <-ic
	fmt.Fprintf(os.Stdout, "%d\n", r)
	fmt.Printf("CloseChanDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------CloseChanDemo end----------->>")
}

func main() {
	CloseChanDemo()
	// exit := make(chan bool)
	// random := make(chan string)
	// go func() {
	// 	for {
	// 		select {
	// 		case <-exit:
	// 			fmt.Printf("  exiting: \n")
	// 			return
	// 		case r := <-random:
	// 			fmt.Printf("  r: %+v\n", r)
	// 		}
	// 	}
	// }()
	//
	// ticks := time.Tick(time.Second * 1)
	// for v := range ticks {
	// 	random <- v.String()
	// }
}
