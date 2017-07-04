// Package main provides ...
package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err.Error())
	}
	// port := l.Addr().(*net.TCPAddr).Port
	en := l.Addr().(*net.TCPAddr)
	// log.Printf("port: %+v\n", port)
	log.Printf("en: %+v\n", en)

	l.Close()
}
