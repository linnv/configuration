package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	name := "win27v@gmail.com"
	h := md5.New()
	h.Write([]byte(name))
	fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil)))

	t := md5.New()
	io.WriteString(t, name)
	fmt.Printf("%s\n", fmt.Sprintf("%x", t.Sum(nil)))
}
