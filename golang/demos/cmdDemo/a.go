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

	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	if len(out) == 0 {
	// 		return err
	// 	}
	// 	return fmt.Errorf("%s\n%v", out, err)
	// }
	// return nil

	println("-----------------------------JustDemo end>>>")
	return
}
