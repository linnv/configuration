package main

import (
	"fmt"
	"time"
)

func readline(ch chan string) {
	var line string
	fmt.Scanf("%s", &line)
	ch <- line
}

func timeout(t chan bool) {
	time.Sleep(5 * time.Second)
	t <- true
}

func main() {
	t := make(chan bool)
	go timeout(t)

	ch := make(chan string)
	go readline(ch)

	select {
	case line := <-ch:
		fmt.Println("Received", line)
	case <-t:
		fmt.Println("Timeout.")
	}
}
