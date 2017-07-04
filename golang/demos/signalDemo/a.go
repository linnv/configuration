// Package main provides ...
package demo

import (
	"log"
	"os"
	"os/signal"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	log.Print("use c-c to exit: \n")
	<-sigChan
	println("-----------------------------JustDemo end>>>")
	return
}
