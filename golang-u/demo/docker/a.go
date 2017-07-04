// Package main provides ...
package demo

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	ConstDemo()
	sp := SelfPath()
	//@toDelete
	fmt.Printf("  sp: %+v\n", sp)
	println("-----------------------------JustDemo end>>>")
	return
}

func ConstDemo() {
	println("//<<-------------------------ConstDemo start-----------")

	const hex = "0123456789abcdef"
	bs := "a1aaaaaaaa"
	b := bs[1]
	// buf.WriteByte(hex[b&0xF])
	//@toDelete
	fmt.Printf("  hex[b&0xF]: %+v\n", hex[b&0xF])
	//@toDelete
	fmt.Printf("  b: %+v\n", b)
	fmt.Printf("  hex[b>>4]: %+v\n", hex[b>>4])
	//@toDelete
	fmt.Printf("  16>>4: %+v\n", 16>>4)
	fmt.Printf("  32>>4: %+v\n", 32>>4)
	println("//---------------------------ConstDemo end----------->>")
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
