// Package main provides ...
package demo

import (
	"fmt"
	"time"

	"github.com/linnv/logx"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

func OkChanelDemo() bool {
	println("//<<-------------------------OkChanelDemo start-----------")
	start := time.Now()
	a := make(chan int)
	e := make(chan struct{})
	go func() {
		for {
			select {
			case i, ok := <-a:
				logx.Debugln("selecting")
				if !ok {
					logx.Debugln("receive chan wrong")
					return
				}
				logx.Debugf("i: %+v\n", i)
				close(e)
			}
		}
	}()
	time.Sleep(time.Second)
	a <- 1
	logx.Debugln("exiting")
	<-e

	fmt.Printf("OkChanelDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------OkChanelDemo end----------->>")
	return false
}

func closeChanDemo() {
	println("//<<-------------------------closeChanDemo start-----------")
	start := time.Now()
	a := make(chan bool)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logx.Debugf("err: %+v\n", err)
			}
		}()
		time.Sleep(time.Millisecond * 50)
		select {
		case a <- true:
			logx.Debugln("wrote")
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			logx.Debugf("err: %+v\n", err)
		}
	}()
	close(a)
	logx.Debugln("closed a")
	time.Sleep(time.Millisecond * 60)
	close(a)
	logx.Debugln("close a again")
	fmt.Printf("closeChanDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------closeChanDemo end----------->>")
}
