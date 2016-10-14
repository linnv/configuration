// Package main provides ...
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", RootHandler)
	fmt.Println("listening 9091")
	http.ListenAndServe(":9091", nil)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("r: %+v\n", r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}
