// Package main provides ...
package newDir

import (
	"os"
	"os/exec"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	cmd := exec.Command("ls", "-all")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return
	}
	println("-----------------------------JustDemo end>>>")
	return
}
