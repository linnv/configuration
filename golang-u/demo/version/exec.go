// Package main provides ...
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/ls", "-l")
	buf, err := cmd.Output()
	// err := cmd.Run()
	if err != nil {
		return
	}
	fmt.Printf("%s", buf)
}
