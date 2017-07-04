// Package main provides ...
package demo

import (
	"strings"
	"unsafe"
)

type A struct {
	Name string `json:"Name"`
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

var s = strings.Repeat("a", 1024)

func copyStr() {
	b := []byte(s)
	_ = string(b)
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func zeroCopyStr() {
	b := str2bytes(s)
	_ = bytes2str(b)
}
