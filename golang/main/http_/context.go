// Package main provides ...
package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	port := ":9999"
	http.HandleFunc("/", Root)
	log.Printf("listening %s\n", port)
	http.ListenAndServe(port, nil)
}

func Root(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 1)
	w.Write([]byte("good response"))
	// e := r.FormValue("e")
	// e = strings.TrimSpace(e)
}
