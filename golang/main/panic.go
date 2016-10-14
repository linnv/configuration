package main

import (
	"fmt"
	"harbor/utils/log"
	"runtime"
	"runtime/debug"
	"time"
)

func PP() {
	fmt.Println("recovered:", recover())
}

func p1inner() {
	if err := recover(); err != nil {
		debug.PrintStack()
	}
	println("inner recover")

}

func pani1Demo() {
	println("//<<-------------------------pani1Demo start-----------")
	start := time.Now()
	defer p1inner()
	panic(" 1 not good")
	fmt.Printf("pani1Demo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------pani1Demo end----------->>")
}

func main() {
	// defer PP()

	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 1<<10, 1<<10)
			// runtime.Stack(buf, false)
			runtime.Stack(buf, false)
			println("outer recover debug")
			debug.PrintStack()
			println("debug end")
			// log.Error("!!!!!!! catch uploadProcess panic:", err)
			log.Error("catch uploadProcess panic:\n", string(buf))

		}
	}()
	pani1Demo()
	//not work
	// defer func() {
	// 	PP()
	// }()

	// panic("not good")
}
