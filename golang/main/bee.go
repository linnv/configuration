package main

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	panic("do it")
	this.Ctx.WriteString("hello world")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}

// Package main provides ...
// package main
//
// import "net/http"
//
// func Root(w http.ResponseWriter, r *http.Request) {
// 	panic("do it")
// 	w.Write([]byte("good"))
// }
//
// func main() {
// 	http.HandleFunc("/", Root)
// 	http.ListenAndServe(":8080", nil)
// }
