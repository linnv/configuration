// Package main provides ...
package main

import (
	"fmt"

	"github.com/hpcloud/tail"
)

func main() {
	c := make(chan string)
	// t, err := tail.TailFile("/Users/Jialin/golang/src/demo/main/a.t", tail.Config{Follow: true, ReOpen: true})
	// t, err := tail.TailFile("/Users/Jialin/golang/src/demo/main/a.t", tail.Config{Follow: true, MustExist: false, Poll: true})
	// t, err := tail.TailFile("/Users/Jialin/golang/src/demo/main/a.t", tail.Config{Follow: true, MustExist: false})

	//this works when edit and save file many times by editor
	go ReceivePort(c)
	t, err := tail.TailFile("/Users/Jialin/golang/src/demo/main/a.t", tail.Config{Follow: true, MustExist: false, Poll: true, ReOpen: true})
	if err != nil {
		return
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
		c <- line.Text
	}
}

func ReceivePort(ch chan string) {
	for {
		select {
		case t := <-ch:
			//@toDelete
			fmt.Printf("  t: %+v\n", t)
		}
	}
}
