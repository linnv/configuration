package main

import (
	"bytes"
	"fmt"
)

func main() {
	buff := bytes.NewBufferString("")
	campaign := "just a demo"
	buff.WriteString(fmt.Sprintf("%+v\n\n", campaign))
	fmt.Printf("buff.String(): %+v\n", buff.String())
}
