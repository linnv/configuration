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

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b")) //parameter is calculated when invoke defer, e.g. when parameter is a function, this funciton will be invoking and the deferFunc will be invoke when this deferFunc returns
	fmt.Println("in b")
	a()
}

func JustDemo() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	b()
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}
