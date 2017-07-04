package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

	"github.com/linnv/logx"
)

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type handler struct {
	net.Conn
	//@TODO close only once
	exit    chan struct{}
	stopper sync.Once
}

func newHandler(conn net.Conn) *handler {
	return &handler{Conn: conn, exit: make(chan struct{}, 1)}
}

func (h *handler) biteMe() {
	go h.write()
	go h.read()
}

func (h *handler) gracefully() {
	h.stopper.Do(func() {
		//you may want to do something more for exiting
		close(h.exit)
	})
}

func (h *handler) write() {
	var s string
	for {
		select {
		case <-h.exit:
			return
		default:
		}
		fmt.Scanln(&s)
		if s == "exit" {
			h.gracefully()
			return
		}
		_, err := h.Write([]byte(s))
		if err != nil {
			log.Printf("connection %s lose: works", h.RemoteAddr().String())
			h.gracefully()
			return
		}
		logx.Debugf("server write to: %+v\n", h.RemoteAddr().String())
	}
}

func (h *handler) read() {
	bs := make([]byte, 1024)
	for {
		select {
		case <-h.exit:
			return
		default:
		}
		_, err := h.Read(bs)
		if err != nil {
			log.Printf("connection %s lose: works", h.RemoteAddr().String())
			h.gracefully()
			return
		}
		if strings.Index(string(bs), "rtt") > -1 {
			logx.Debugln("what?")
			h.Write([]byte("rtt"))
		}
		log.Printf("server receives: %v\n from %s\n", string(bs), h.RemoteAddr().String())
	}
}

var port = flag.Int("port", 6798, "listening port")

func main() {
	flag.Parse()
	ip := "127.0.0.1"
	host := fmt.Sprint(ip, ":", *port)
	hubListener, err := net.Listen("tcp", host)
	checkErr(err)
	log.Printf("listen on : %+v\n", *port)
	for {
		oneClient, err := hubListener.Accept()
		checkErr(err)
		newHandler(oneClient).biteMe()
	}
}
