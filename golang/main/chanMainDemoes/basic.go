// Package main provides ...
package main

import (
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"
)

//when close done channel in main funtion, all upstream channel running a goroutine will know to exit
func CloseChanAfterAllSending(done chan struct{}, cs ...<-chan int) {
	var wg sync.WaitGroup
	out := make(chan int)

	funcOut := func(c <-chan int) {
		// out <- <-c
		for v := range c {
			select {
			case out <- v:
				//if done is closed, cancel blocking other channel
			case <-done:
			}
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, v := range cs {
		funcOut(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}

type Ai interface {
	Done() <-chan struct{}
}

type emptyCtx int

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (e *emptyCtx) String() string {
	switch e {
	// case background:
	// 	return "context.Background"
	// case todo:
	// 	return "context.TODO"
	case e:
		return "exxfjifejfej"
	case ai:
		return "ai"
	}
	return "unknown empty Context"
}

func interfaceDemo(a Ai) {
	println("//<<-------------------------interfaceDemo start-----------")
	start := time.Now()
	itv := reflect.ValueOf(a)
	it := itv.Type()
	fmt.Printf("i.UnmMethod(): %+v\n", it.NumMethod())
	for i := 0; i < it.NumMethod(); i++ {
		switch handler := itv.Method(i).Interface().(type) {
		case func() string:
			log.Println(": works\t", handler())
		}
	}
	fmt.Printf("interfaceDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------interfaceDemo end----------->>")
}

var e = new(emptyCtx)
var ai Ai

func main() {
	select {
	case x := <-e.Done():
		log.Printf("nil: works%v\n", x)
	default:
		log.Printf("default : works\n")
	}
	xx := fmt.Sprint(e)
	fmt.Printf("fmt.Sprint(e): %q\n", xx)

	ai = new(emptyCtx)
	interfaceDemo(ai)
	// fmt.Printf("fmt.Sprint(ai): %q\n", ai)

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
