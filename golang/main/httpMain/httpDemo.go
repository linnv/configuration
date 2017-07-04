// Package main provides ...
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/linnv/logx"
)

var httpClient *http.Client

func main() {
	// go GetHttpClient()
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/get", WriteHandler)
	port := ":8089"
	fmt.Printf("listening : %+v\n", port)
	server := &http.Server{Addr: port, Handler: nil}
	go http.ListenAndServe(port, nil)
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	logx.Debug("use c-c to exit: \n")
	<-sigChan
	server.Shutdown(nil)

}

func GetHttpClient() {
	httpClient = &http.Client{
		Timeout: time.Duration(time.Second * 15),
	}
}

type A struct {
	Name string `json:"name"`
}

func WriteHandler(w http.ResponseWriter, r *http.Request) {
	todo := A{Name: "xx"}
	bs, err := json.Marshal(todo)
	if err != nil {
		panic(err.Error())
		return
	}

	w.Write(bs)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	println("call twice?")
	//only post works
	// body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	// if err != nil {
	// 	panic(err)
	// }
	// if err := r.Body.Close(); err != nil {
	// 	panic(err)
	// }
	// todo := A{Name: "xx"}
	// fmt.Printf("string(body): %+v\n", string(body))
	// if err := json.Unmarshal(body, &todo); err != nil {
	// }
	// fmt.Printf("todo: %+v\n", todo)
	// query := r.URL.Query()
	// tmpReqLenStr := query.Get("c")
	// fmt.Printf("tmpReqLenStr: %+v\n", tmpReqLenStr)
	// fmt.Fprintf(w, "todo: %+v\n", todo)
	// fmt.Fprintf(w, "root: %s\n", r.URL.Path)
	// fmt.Fprintf(w, "url: %s\n", r.URL)
	// fmt.Fprintf(w, "method: %s\n", r.Method)
	// fmt.Fprintf(w, "ReadRequestURI: %s\n", r.RequestURI)
	// fmt.Fprintf(w, "Proto: %s\n", r.Proto)
	// fmt.Fprintf(w, "host: %s\n", r.Host)
	// fmt.Println("handler demo")
	// fmt.Fprintln(w, "handler demoee")
	// w.Write([]byte(homePage))
	// time.Sleep(time.Second * 10)
	// logx.EnableDevMode(true)
	// logx.LogConfigure()

	// logx.Debug(": %+v\n", r.URL.Query())
	// logx.LogConfigure()
	logx.Debug("keyword : %+v\n", r.FormValue("keyword"))

	// for k, v := range r.Header {
	// 	fmt.Printf("%+v: %+v\n", k, v)
	// }
	w.Header().Set("Access-Control-Allow-Headers", "whatw,what")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// '':'Access-Control-Allow-Origin,X-Requested-With, X-HTTP-Method-Override, Content-Type, Accept',
	w.Write([]byte("data power by golang " + time.Now().String()))
}

const (
	homePage = `
<!DOCTYPE html>
<html>
	<head>
		<title></title>
		<meta charset="utf-8" />
		<script src="./demo.js" type="text/javascript" charset="utf-8"></script>
	</head>
	<body onload="init()">
		<header>
			<p>json demo</p>	
		</header> 

		<center>
			<p id="send">send</p>
		</center>

		<button id="myBtn">onclike demo</button>
		<!-- <button id="myBtn" onclick="displayDate()">onclike demo</button> -->
		<p id="demo"></p>

<input type="text" id="fname" onchange="myFunction()">
<p>当您离开输入字段时，会触发将输入文本转换为大写的函数。</p>

			<p id="demojson"></p>
			<p id="ajaxDemo">ajax demo</p>
		<button id="ajax" onclick="getData()">onclike ajax demo</button>
		<footer>
		</footer>

		<script type="text/javascript" charset="utf-8">
function myFunction()
{
	var x=document.getElementById("fname");
	x.value=x.value.toUpperCase();
}

var xmlr=new XMLHttpRequest();
var i=0
function getData() {

var xmlrfalse=new XMLHttpRequest();
	xmlrfalse.open("GET","http://127.0.0.1:8089/get",false);
	xmlrfalse.send();
	document.getElementById("ajaxDemo").innerHTML=xmlrfalse.responseText+i;
	i++;
	console.log(xmlrfalse.responseText);
}

function init() {
	document.getElementById("myBtn").onclick=function(){displayDate()};
}

function trash(){
	xmlr.onreadystatechange=function()
	{
		if (xmlr.readyState==4 && xmlr.status==200)
		{
			document.getElementById("ajaxDemo").innerHTML=xmlr.responseText;
			console.log(xmlr.responseText);
		}
	}
	xmlr.open("GET","http://127.0.0.1:8089/get",true);
	xmlr.send();
}

function displayDate()
{
	document.getElementById("demo").innerHTML=Date();
}

		</script>

</body>
`
)

// func jsonPost() {
// 	m := map[string]interface{}{
// 		"name":    "backy",
// 		"species": "dog",
// 	}
// 	mJson, _ := json.Marshal(m)
// 	contentReader := bytes.NewReader(mJson)
// 	req, _ := http.NewRequest("POST", "http://example.com", contentReader)
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Notes", "GoRequest is coming!")
// 	client := &http.Client{}
// 	resp, _ := client.Do(req)
// }
