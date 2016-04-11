package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	p := &Proxy{}
	http.Handle("/", p)
	http.HandleFunc("/root", p.root)
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func Foo() {
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("hello")
		default:
			fmt.Println("default")
			return
		}
	}

	for {
		ch := make(chan string, 1)
		select {
		case ch <- "xx":
			println("channel sent")
		default:
			println("don't block channel")
		}
	}

}

type Proxy struct {
	mu    sync.Mutex
	name  string
	proxy http.Handler
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("default handler")
}

func (p *Proxy) root(w http.ResponseWriter, r *http.Request) {
	log.Println("root handler")
}
