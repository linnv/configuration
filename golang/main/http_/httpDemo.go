// Package main provides ...
package main

import (
	"fmt"
	"net/http"
	"path"
	"time"
)

var httpClient *http.Client

func main() {
	// fmt.Printf("path.Dir(): %+v\n", path.Join("/ddd", string(os.PathSeparator)))
	fmt.Printf("path.Dir(): %+v\n", path.Join("/ddd", "addd"))
	// go GetHttpClient()
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/debug", GetParameter)
	port := ":8099"
	fmt.Printf("listening : %+v\n", port)
	http.ListenAndServe(port, nil)

}

func GetHttpClient() {
	httpClient = &http.Client{
		Timeout: time.Duration(time.Second * 15),
	}
}

type A struct {
	Name string `json:"name"`
}

func GetParameter(w http.ResponseWriter, r *http.Request) {
	debugstr := r.FormValue("open")
	fmt.Printf("debugstr: %+v\n", debugstr)
	w.Write(nil)
	return
	// debug, err := strconv.ParseBool(debugstr)
	// utils.CheckErr(err)
	// fmt.Printf("debug: %+v\n", debug)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// r.Header
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
	w.Write([]byte(homePage))
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

//
// 	js := `[{
//   "stock": {
//     "name": 300111,
//     "id": 10
//   }},
//   {"astock": {
//     "name": 300222,
//     "id": 20
//   }},
//   {"bstock": {
//     "name": 300333,
//     "id": 30
//   }}
// ]`
