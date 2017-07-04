// Package main provides ...
package demo

import (
	"encoding/base64"
	"fmt"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func JustDemo() {
	println("<<<JustDemo---------------------------")
	a := 19
	fmt.Printf("a: %+v\n", a)
	println("-----------------------------JustDemo>>>")
	return
}
