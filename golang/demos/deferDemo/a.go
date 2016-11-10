// Package main provides ...
package newDir

import (
	"fmt"
	"log"
	"time"
)

func PanicDeferRecoverDemo() {
	println("//<<-------------------------PanicDeferRecoverDemo start-----------")
	start := time.Now()
	defer func() {
		log.Println("level 2: works")
		r := recover()
		log.Printf("r: %+v\n", r)
	}()
	panic("from level 1 function")
	fmt.Printf("PanicDeferRecoverDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------PanicDeferRecoverDemo end----------->>")
}

func DeferDemo() {
	println("//<<-------------------------DeferDemo start-----------")
	start := time.Now()
	defer log.Println("2: works")
	defer func() {
		log.Println("iner 4: works")
	}()
	defer log.Println("3: works")
	log.Println("1: works")

	fmt.Printf("DeferDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------DeferDemo end----------->>")
}
func JustDemo() {
	println("<<<JustDemo start---------------------------")
	DeferDemo()
	println("-----------------------------JustDemo end>>>")
	return
}
