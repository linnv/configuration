package main

import "net/http"

func main() {
	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/file-status", FileHandler)
	http.ListenAndServe(":8099", nil)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to file picker zone"))
	return
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
}
