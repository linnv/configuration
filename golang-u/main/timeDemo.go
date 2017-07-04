// Package main provides ...
package main

import (
	"fmt"
	"reflect"
	"time"
)

//@TODO synca Wait
func tickDemo() {
	println("<<<tickDemo---------------------------")
	tic := time.NewTicker(time.Millisecond * 100)
	// time.Friday
	go func() {
		for t := range tic.C {
			fmt.Printf("Tick at : %+v\n", t)
		}
	}()
	time.Sleep(time.Millisecond * 1500)
	tic.Stop()
	// println("-----------------------------tickDemo>>>")
	// return
}

func durationDemo() {
	println("<<<durationDemo---------------------------")

	now := time.Now()
	t := reflect.ValueOf(now).Kind()
	println("-----------------------------durationDemo>>>")
	time.Sleep(time.Second)
	done := time.Now()
	fmt.Printf("time consumpe: %.2f s\n", done.Sub(now).Seconds())
	return
}

func timeStampDemo() {
	println("<<<timeStampDemo---------------------------")
	// ts := time.Unix()
	ts := time.Now().Unix()
	fmt.Printf("ts: %d\n", int64(ts))
	// tsr := ts.Format("20060102")
	// i, _ := strconv.ParseInt(tsr, 10, 64)
	// fmt.Printf("tsr: %+v\n", tsr)
	// fmt.Printf("i: %+v\n", i)
	println("-----------------------------timeStampDemo>>>")
	return
}

func main() {
	timeStampDemo()
	// durationDemo()
	// go tickDemo()
	// tickDemo()

	// ts := int32(time.Now().Unix())
	// fmt.Printf("ts: %+v\n", ts)

	// tic := time.NewTicker(time.Second * 2)
	// go func() {
	// 	for t := range tic.C {
	// 		fmt.Printf("Tick at : %+v\n", t)
	// 	}
	// }()

}
