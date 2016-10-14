// Package main provides ...
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// r := mux.NewRouter()
	// r.HandleFunc("/a", basic)
	// r.HandleFunc("/b*", ServeHTTP)
	// http.Handle("/", r)

	http.HandleFunc("/", basic)
	log.Fatal(http.ListenAndServe(":9099", nil))
}

func basic(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Printf("%+v: %+v\n", k, v)
	}

	w.Header().Set("User-Agent", "manunally edit UA")
	// log.Println(r.Header)
	// w.Header().Set("Date", "manunally edit date")
	// w.WriteHeader(http.StatusAccepted)
	// log.Printf("http.StatusAccepted: %+v\n", http.StatusAccepted)
	w.Write([]byte("basic handler"))
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ServeHTTP"))
}
