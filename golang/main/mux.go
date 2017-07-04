// Package main provides ...
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/linnv/logx"
)

func One(w http.ResponseWriter, r *http.Request) {
	logx.Debugf(": %+v\n", r.URL.Path)
	if r.URL.Path == "/" {
		http.Redirect(w, r, r.URL.Path+"main", http.StatusFound)
	} else {
		w.Write([]byte("good"))
		// w.WriteHeader(http.StatusNotFound)
	}

}

func main() {
	r := mux.NewRouter()

	// This will serve files under http://localhost:8000/static/<filename>
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/Users/Jialin/web/jialinwu_local"))))
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/Users/Jialin/golang/src/ssp_mbv_web/public"))))
	// r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/Users/Jialin/golang/src/ssp_mbv_web/public"))))
	r.HandleFunc("/one", One)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/Users/Jialin/golang/src/ssp_mbv_web/public/main"))))
	// r.PathPrefix("/two").Handler(
	// 	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		logx.Debugf(": %+v\n", r.URL.Path)
	// 		if r.URL.Path == "/" {
	// 			http.Redirect(w, r, r.URL.Path+"main", http.StatusFound)
	// 		} else {
	// 			w.WriteHeader(http.StatusNotFound)
	// 		}
	// 	}))
	// r.PathPrefix("/main").Handler(http.FileServer(http.Dir("/Users/Jialin/golang/src/ssp_mbv_web/public/main")))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8092",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
