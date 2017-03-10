// Package main provides ...
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/linnv/logx"
)

type Handler struct{}

var handlerFunc map[string]http.HandlerFunc

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logx.Debug("r: %+v\n", r.RequestURI)
	if hf, ok := handlerFunc[r.RequestURI]; ok {
		hf(w, r)
		return
	}
	fmt.Fprintln(w, "Hello, client, no data were found")
}

func init() {
	handlerFunc = make(map[string]http.HandlerFunc)
	handlerFunc["/home"] = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client,welcome to home")
	}
	handlerFunc["/color/city"] = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client,color city")
	}
}

func getCheck(url string) (want string) {
	res, err := http.Get(url)
	if err != nil {
		logx.Fatalf(err.Error())
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		logx.Fatalf(err.Error())
	}

	fmt.Printf("%s", greeting)
	return string(greeting)
}

func main() {
	ts := httptest.NewServer(Handler{})
	defer ts.Close()
	listOfURI := []string{
		"/home",
		"/color/city",
	}
	for _, v := range listOfURI {
		want := getCheck(ts.URL + v)
		logx.Debug("want: %+v\n", want)
	}
}
