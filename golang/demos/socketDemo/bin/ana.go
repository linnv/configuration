// Package main provides ...
package main

import (
	"fmt"
	"log"
	"net"
)

type headerType byte

const (
	hdrNewConn headerType = 0xc4 + iota
	hdrNewSession
)

func main() {
	ip := "104.224.132.128"
	port := 6800
	host := fmt.Sprint(ip, ":", port)
	log.Printf("host: %+v\n", host)
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Println("dial tcp connection err:%s", err.Error())
		return
	}
	defer conn.Close()
	log.Println("good : works")
	return
	// log.Printf("hdrNewConn: %+v\n", hdrNewConn)
	// log.Printf("hdrNewConn: %+v\n", hdrNewConn/8)
	// log.Printf("hdrNewConn: %+v\n", hdrNewConn%8)
	// log.Printf("hdrNewSession: %+v\n", hdrNewSession)
	// log.Printf("0x1: %+v\n", 0x1)
	// log.Printf("0x10: %+v\n", 0x10)
}
