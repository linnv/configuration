package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

type Handler struct{}

var handlerFunc map[string]http.HandlerFunc

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.RequestURI: %+v\n", r.RequestURI)
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

func getCheck(url string) (ouput string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err.Error())
	}
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
		output := getCheck(ts.URL + v)
		fmt.Printf("output: %+v\n", output)
	}
}
