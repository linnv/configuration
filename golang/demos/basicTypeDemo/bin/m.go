package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	println("//<<-------------------------ScanDemo start-----------")
	start := time.Now()
	var s string
	for {
		fmt.Scanln(&s)
		if s == "exit" {
			return
		}
		log.Printf("s: %+v\n", s)
	}
	fmt.Printf("ScanDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------ScanDemo end----------->>")
}
