// Package main provides ...
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	sp := SelfPath()
	//@toDelete
	fmt.Printf("  sp: %+v\n", sp)
	as := []string{"-l", "-h"}
	bs, err := exec.Command("ll", as...).CombinedOutput()
	if err != nil {
		panic(err.Error())
		return
	}
	//@toDelete
	fmt.Printf("  string(bs): %+v\n", string(bs))
}

func SelfPath() string {
	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		if os.IsNotExist(err) {
			return ""
		}
		if execErr, ok := err.(*exec.Error); ok && os.IsNotExist(execErr.Err) {
			return ""
		}
		panic(err)
	}
	path, err = filepath.Abs(path)
	if err != nil {
		if os.IsNotExist(err) {
			return ""
		}
		panic(err)
	}
	return path
}
