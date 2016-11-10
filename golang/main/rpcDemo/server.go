package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// type Echo struct{}
type Echo int

func (e *Echo) Do(args string, replay *string) error {
	log.Println("invoking do(): works " + time.Now().String())
	*replay = fmt.Sprintf("reply to [%s]msg from server", args)
	return nil
}

func main() {
	serve := rpc.NewServer()
	serve.Register(new(Echo))
	// serve.HandleHTTP("rpcPat", "rpcRouter")
	//@TODO it use the default http.defaultServeMux, try to build from the root manunally
	serve.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	// rpc.Register()
	// rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":9981")
	if err != nil {
		panic(err.Error())
	}
	http.Serve(l, nil)
}
