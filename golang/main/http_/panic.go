package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	m := mux.NewRouter()
	m.Handle("/", RecoverWrap(http.HandlerFunc(handler))).Methods("GET")

	http.Handle("/", m)
	log.Println("Listening...")

	http.ListenAndServe(":3001", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	panic(errors.New("panicing from error"))
}

func RecoverWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// var err error
		// defer func() {
		// 	r := recover()
		// 	if r != nil {
		// 		switch t := r.(type) {
		// 		case string:
		// 			err = errors.New(t)
		// 		case error:
		// 			err = t
		// 		default:
		// 			err = errors.New("Unknown error")
		// 		}
		// 		sendMeMail(err)
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	}
		// }()
		h.ServeHTTP(w, r)
	})
}

func sendMeMail(err error) {
	println("do mail sending")
	fmt.Printf("err.: %+v\n", err.Error())
	// send mail
}
