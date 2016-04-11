// Package main provides ...
package main

import (
	"fmt"
	"os"
	"time"
)

func init() {
	ca, cb = make(chan int), make(chan int)
	POLL_DURATION = 250 * time.Millisecond
}

var POLL_DURATION time.Duration

var (
	ca, cb chan int
)

func Rec() {
	for {
		// if _, err := os.Stat(fw.Filename); err == nil {
		// 	return nil
		// } else if !os.IsNotExist(err) {
		// 	return err
		// }
		fmt.Printf("  for time.Now(): %+v\n", time.Now())
		select {
		case <-time.After(POLL_DURATION):
			fmt.Printf("  coninue time.Now(): %+v\n", time.Now())
			continue
		case d := <-ca:
			fmt.Printf("a  d: %+v\n", d)
		case d := <-cb:
			fmt.Printf("b  d: %+v\n", d)
		}
	}
	panic("unreachable")
}

func send() {

}

func main() {
	go Rec()
	ca <- 111
	time.Sleep(time.Second)
	cb <- 111
	os.Stdout.Write(append([]byte("end in 10 seconeds"), '\n'))
	time.Sleep(time.Second * 10)

}
