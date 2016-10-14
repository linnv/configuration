// Package main provides ...
package main

import (
	"errors"
	"os"
	"os/exec"
	"time"
)

func JustDemo() error {
	println("<<<JustDemo start---------------------------")
	n := time.Now().String()
	cmd := exec.Command("bash", "./pause.sh", n)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return errors.New("bad")
	}
	println("-----------------------------JustDemo end>>>")
	return errors.New("bad")
}

func main() {
	go JustDemo()
	time.Sleep(time.Second)
	go JustDemo()
	time.Sleep(time.Second * 10)
	println("should not run to herer")
}
