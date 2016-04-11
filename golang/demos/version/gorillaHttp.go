// Package main provides ...
package main

import (
	"diy_package/buf"
	"diy_package/tmp"
	"fmt"
	// "github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

const httpRoot = "/home/sunteng/golang/src/demo"

func main() {

	tmp.Tmp()
	buf.Tmp()
	var a int
	defer func() {
		fmt.Println(a)
	}()
	a = 10
	// r := mux.NewRouter()
	// r.HandleFunc("/", rootHandler)
	// r.HandleFunc("/view/", viewHandler)
	// // r.HandleFunc("/html/", htmlHandler) ///url/html/ works,but /url/html doesn't work
	// r.HandleFunc("/html", htmlHandler) ///url/html/ works,but /url/html doesn't work     /url/html and /url/html/ is different
	// http.Handle("/", r)
	// http.ListenAndServe(":8089", nil)

	// http.HandleFunc("/", rootHandler)
	// http.HandleFunc("/view/", viewHandler)
	// http.HandleFunc("/html/", htmlHandler)
	// http.ListenAndServe(":8089", nil)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "root: %s\n", r.URL.Path)
	fmt.Fprintf(w, "url: %s\n", r.URL)
	fmt.Fprintf(w, "method: %s\n", r.Method)
	fmt.Fprintf(w, "ReadRequestURI: %s\n", r.RequestURI)
	fmt.Fprintf(w, "Proto: %s\n", r.Proto)
	fmt.Fprintf(w, "host: %s\n", r.Host)
	fmt.Println("handler demo")
	fmt.Fprintln(w, "handler demoee")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "view handler")
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("handler : %s\n", r.URL.Path)
	filename := httpRoot + r.URL.Path + "/blog.html"
	// filename := httpRoot + r.URL.Path
	// filename := httpRoot + "/view/"
	fmt.Println("path", r.URL.Path)
	fmt.Println("filename", filename)
	fileext := filepath.Ext(filename)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		// fmt.Fprintln(w, "404 Not Found!\n")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var contype string
	switch fileext {
	case ".html", "htm":
		contype = "text/html"
	case ".js":
		contype = "application/javascript"
	case ".css":
		contype = "text/css"
	case ".png":
		contype = "image/png"
	case ".jpg":
		contype = "image/jpg"
	case ".git":
		contype = "image/git"
	default:
		contype = "text/plain"

	}
	fmt.Printf("ext %s,contype = %s\n", fileext, contype)
	w.Header().Set("Content-Type", contype)
	fmt.Fprintf(w, "%s", content)
}
