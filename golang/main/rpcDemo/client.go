// Package main provides ...
package main

import (
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9981")
	if err != nil {
		panic(err.Error())
	}
	msg := "hi server"
	var reply string
	err = client.Call("Echo.Do", msg, &reply)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("reply: %+v\n", reply)
}
