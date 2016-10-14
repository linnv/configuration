// Package main provides ...
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var httpClient *http.Client
var key = "65b4780597f3c3d23e629f8b1d30002f"

const httpRoot = "/home/jialin/golang/src/demo"

var MyTransport http.RoundTripper

func init() {

	MyTransport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 1 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 1 * time.Second,
	}
}

func main() {
	// go GetHttpClient()

	// httpClient = &http.Client{}
	// err := GetDemo()
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }

	// pd := &PostBody{Key: "xxx",
	// 	Data: &[]PostData{PostData{Title: "demotitle", Author: "xxxx", CreateDate: time.Now(), Hit: 1999, Source: "xxxxxsource", Url: "fewwwurl "}, PostData{Title: "demotitle", Author: "xxxx", CreateDate: time.Now(), Hit: 1999, Source: "xxxxxsource", Url: "fewwwurl "}}}
	// fmt.Printf("before pd: %+v\n", pd.Data)
	// b, err := FillPostData(pd)
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }
	// fmt.Printf("b: %+v\n", b)

	// http.HandleFunc("/", rootHandler)
	// http.HandleFunc("/post", getPostHandler)
	// http.HandleFunc("/view/", viewHandler)
	// http.HandleFunc("/html/", htmlHandler)
	// http.HandleFunc("/page/", pageHandler)
	// http.ListenAndServe(":8049", nil)

	// PostDemo("fefe", "feefe")
	// PostDataDemo()
	// ParamtersOfPostDemo()
	// HttpTimeConsumptionDemo()
	// simplePost()
}

func Article(w http.ResponseWriter, r *http.Request) {

	// var b byteBuffer

	w.Write([]byte("xxx"))
}

func simplePost() {
	println("//<<-------------------------simplePost start-----------")
	start := time.Now()

	var a = struct {
		A      string   `json:"open"`
		Aa     []string `json:"a a"`
		Aaint  []int    `json:"b"`
		Aauint []uint   `json:"c"`
	}{A: "xxx33rj33j",
		Aa:     []string{"dd", "fefe"},
		Aaint:  []int{2, 5},
		Aauint: []uint{23, 445},
	}

	form, err := Json2HttpForm(a)
	if err != nil {
		panic(err.Error())
	}

	data := form.Encode()
	fmt.Printf("data for preparation simplePost costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	http.DefaultClient.Transport = MyTransport
	start = time.Now()
	fmt.Printf("start.String(): %+v========\n", start.String())

	req, _ := http.NewRequest("POST", "http://127.0.0.1:8099/debug", bytes.NewBuffer([]byte(data)))
	//this header must be set when posting data in body
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// fmt.Printf("data for sending simplePost costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))

	req.Close = false
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}

	ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Printf("post costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------simplePost end----------->>")
}

func HttpTimeConsumptionDemo() {
	println("//<<-------------------------HttpTimeConsumptionDemo start-----------")
	start := time.Now()

	var a = struct {
		A      string   `json:"open"`
		Aa     []string `json:"a a"`
		Aaint  []int    `json:"b"`
		Aauint []uint   `json:"c"`
	}{A: "xxx",
		Aa:     []string{"dd", "fefe"},
		Aaint:  []int{2, 5},
		Aauint: []uint{23, 445},
	}
	form, err := Json2HttpForm(a)
	if err != nil {
		panic(err.Error())
	}

	data := form.Encode()
	fmt.Printf("data for preparation HttpTimeConsumptionDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	http.DefaultClient.Transport = MyTransport
	count := 0
	for {
		start = time.Now()
		fmt.Printf("start.String(): %+v========\n", start.String())

		req, _ := http.NewRequest("POST", "http://127.0.0.1:8099/debug", bytes.NewBuffer([]byte(data)))
		// req, _ := http.NewRequest("POST", "http://cent.local", bytes.NewBuffer([]byte(data)))
		//this header must be set when posting data in body
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		// fmt.Printf("data for sending HttpTimeConsumptionDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))

		req.Close = true
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err.Error())
		}

		ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		fmt.Printf("post costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
		fmt.Printf("end.String(): %+v-------", time.Now().String())
		// println("sleeping 2 second")
		// time.Sleep(time.Second * 2)
		// println("wake up\n\n")
		count++
		time.Sleep(time.Second * 5)
		fmt.Printf("count: %+v\n", count)
	}
	println("//---------------------------HttpTimeConsumptionDemo end----------->>")
}

func Json2HttpForm(s interface{}) (form *url.Values, err error) {
	if reflect.TypeOf(s).Kind() != reflect.Struct {
		return nil, fmt.Errorf("usopport convert type %s", reflect.TypeOf(s).Kind())
	}

	types := reflect.TypeOf(s)
	values := reflect.ValueOf(s)
	form = &url.Values{}
	for i := 0; i < types.NumField(); i++ {
		a := types.Field(i)
		av := values.Field(i)
		switch av.Kind() {
		case reflect.String:
			form.Add(a.Tag.Get("json"), av.String())
		case reflect.Float32, reflect.Float64:
			form.Add(a.Tag.Get("json"), fmt.Sprintf("%f", av.Float()))
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			form.Add(a.Tag.Get("json"), fmt.Sprintf("%d", av.Int()))
		case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			form.Add(a.Tag.Get("json"), fmt.Sprintf("%d", av.Uint()))
		case reflect.Slice:
			vs := reflect.ValueOf(av.Interface())
			vsLen := vs.Len()
			ss := make([]string, vsLen)
			switch t := vs.Interface().(type) {
			case []int, []int16, []int32, []int64:
				for j := 0; j < vsLen; j++ {
					ss[j] = fmt.Sprintf("%d", vs.Index(j).Int())
				}
			case []uint, []uint16, []uint32, []uint64:
				for j := 0; j < vsLen; j++ {
					ss[j] = fmt.Sprintf("%d", vs.Index(j).Uint())
				}
			case []float32, []float64:
				for j := 0; j < vsLen; j++ {
					ss[j] = fmt.Sprintf("%f", vs.Index(j).Float())
				}
			case []string:
				for j := 0; j < vsLen; j++ {
					ss[j] = fmt.Sprintf("%s", vs.Index(j))
				}
			default:
				return nil, fmt.Errorf("%s, element type is unsupported", t)

			}
			form.Add(a.Tag.Get("json"), strings.Join(ss, ","))

		default:
			return nil, fmt.Errorf("%s, element type is unsupported", av.Kind())

		}
	}

	return
}

func GetHttpClient() {
	httpClient = &http.Client{
		Timeout: time.Duration(time.Second * 15),
	}
}

type Page struct {
	Index int `json:"Index"`
	Pack  interface{}
}

var gp []*Page

func init() {
	count := 10000
	gp = make([]*Page, count)
	for k, _ := range gp {
		gp[k] = &Page{k, "row" + strconv.Itoa(k)}
	}

}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseForm()
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }

	// requestbody, _ := ioutil.ReadAll(r.Body)
	// fmt.Printf("requestbody: %+v\n", requestbody)
	// query := r.URL.Query()
	query := r.URL.Query()
	val1 := query.Get("key")
	fmt.Printf("key1: %+v\n", val1)
	skips, sort, limits := r.FormValue("offset"), r.FormValue("sort"), r.FormValue("limit")
	// fmt.Printf("query.Get(limit): %+v\n", query.Get("limit"))
	// skips, sort, limits := query.Get("offset"), query.Get("sort"), query.Get("limit")
	fmt.Printf("limits: %+v\n", limits)
	fmt.Printf("skips: %+v\n", skips)
	skip, _ := strconv.Atoi(skips)
	limit, _ := strconv.Atoi(limits)

	if skip > 1 {
		skip = skip - 1
	}
	fmt.Printf("skip: %+v\n", skip)
	fmt.Printf("sort: %+v\n", sort)
	fmt.Printf("limit: %+v\n", limit)
	limitp := skip + limit
	bs, err := json.Marshal(gp[skip:limitp])
	if err != nil {
		return
	}

	w.Write(bs)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "root: %s\n", r.URL.Path)
	fmt.Fprintf(w, "url: %s\n", r.URL)
	fmt.Fprintf(w, "method: %s\n", r.Method)
	fmt.Fprintf(w, "ReadRequestURI: %s\n", r.RequestURI)
	fmt.Fprintf(w, "Proto: %s\n", r.Proto)
	fmt.Fprintf(w, "host: %s\n", r.Host)
	fmt.Println("handler demo")
	query := r.URL.Query()
	val1 := query.Get("key1")
	fmt.Printf("key1: %+v\n", val1)
	fmt.Fprintln(w, "handler demoee")
}

type ClientPostData struct {
	Ids   []int
	All   bool
	Token string
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	htmlbody, err := ioutil.ReadAll(r.Body)
	// rb := reflect.ValueOf(&b).Elem().Interface()
	// fmt.Printf("string(b): %+v\n", rb.(string))
	fmt.Printf("html body: %+v\n", string(htmlbody))
	r.Body.Close()
	if err != nil {
		return
	}
	c := ClientPostData{}
	err = json.Unmarshal(htmlbody, &c)
	if err != nil {
		return
	}
	//@toDelete
	fmt.Printf("c: %+v\n", c)
	w.Write(htmlbody)
	return
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "view handler")
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("handler : %s\n", r.URL.Path)
	filename := httpRoot + r.URL.Path
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

func GetDemo() error {
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", req.Host)
	req.Header.Set("Accept-Encoding", "identity;q=1, *;q=0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.153 Safari/537.36")
	req.Header.Set("RA-Sid", "DA14F3FC-20140730-093442-bad426-396f08")
	req.Header.Set("RA-Ver", "2.7.1")
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("获取远程文件失败[%d]", resp.StatusCode)
	}
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	// rb := reflect.ValueOf(&b).Elem().Interface()
	fmt.Printf("string(b): %+v\n", string(b))
	// fmt.Printf("string(b): %+v\n", rb.(string))
	resp.Body.Close()
	if err != nil {
		return err
	}
	return nil
}

type PostBody struct {
	Key  string      `json:"key"`
	Data *[]PostData `json:"data"`
}

type PostData struct {
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	CreateDate time.Time `json:"create_date"`
	Hit        int       `json:"hit"`
	Content    string    `json:"content"`
	Source     string    `json:"source"`
	Url        string    `json:"url"`
}

func FillPostData(postBody *PostBody) (postStr string, err error) {
	postBody.Key = key
	cd := time.Now()
	tmpd := postBody.Data
	for k, _ := range *postBody.Data { //addre not support range or index
		(*postBody.Data)[k].CreateDate = cd
	}
	fmt.Printf("afer tmpd: %+v\n", tmpd)
	bs, err := json.Marshal(postBody)
	postStr = string(bs)
	return

}

func PostDataDemo() {
	println("<<<-----------------PostDataDemo----------")
	d := PostBody{Key: key,
		Data: &[]PostData{PostData{Title: "demotitle", Author: "xxxx", CreateDate: time.Now(), Hit: 1999, Source: "xxxxxsource", Url: "fewwwurl "}}}
	// fmt.Printf("d: %+v\n", d)
	bs, err := json.Marshal(d)
	if err != nil {
		panic(err.Error())
		return
	}
	fmt.Printf("bd: %+v\n", string(bs))
	println("-----------------PostDataDemo------------>>>")
}

func PostDemo(json string, apiurl string) error {
	println("<<<-----------------PostDemo----------")
	// var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	v := url.Values{}
	v.Add("a", "33")
	// v.Set("b", "3b3")
	v.Set("b", "3b3")
	data := v.Encode()
	fmt.Printf("data: %+v\n", data)
	req, err := http.NewRequest("POST", apiurl, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err
	}
	fmt.Printf(": %+v\n", req.RequestURI)

	// req.Header.Set("Accept", "*/*")
	// req.Header.Set("Connection", "keep-alive")
	// req.Header.Set("Host", req.Host)
	// req.Header.Set("Accept-Encoding", "identity;q=1, *;q=0")
	// req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.153 Safari/537.36")
	// req.Header.Set("RA-Sid", "DA14F3FC-20140730-093442-bad426-396f08")
	// req.Header.Set("RA-Ver", "2.7.1")
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("获取远程文件失败[%d]", resp.StatusCode)
	}
	var htmlbody []byte
	htmlbody, err = ioutil.ReadAll(resp.Body)
	// rb := reflect.ValueOf(&b).Elem().Interface()
	// fmt.Printf("string(b): %+v\n", rb.(string))
	fmt.Printf("html body: %+v\n", string(htmlbody))
	resp.Body.Close()
	if err != nil {
		return err
	}
	println("-----------------PostDemo------------>>>")
	return nil
}

// func ParameterOfGetDemo() {
// 	println("<<<-----------------ParameterOfGetDemo----------")
//
// 	query := r.URL.Query()
// 	if err != nil {
// 		log.Error("url parse query error")
// 		err = errors.New("url parse query error")
// 		return
// 	}
// 	fs := query.Get("fs")
// 	if fs == "" {
// 		log.Errorf("files empty!")
// 		err = errors.New("files empty!")
// 		return
// 	}
//
// 	//offical example
// 	// req := &Request{Method: "GET"}
// 	// req.URL, _ = url.Parse("http://www.google.com/search?q=foo&q=bar")
// 	//// FormValue returns the first value for the named component of the query.
// 	//// POST and PUT body parameters take precedence over URL query string values.
// 	// if q := req.FormValue("q"); q != "foo" {
// 	// 	t.Errorf(`req.FormValue("q") = %q, want "foo"`, q)
// 	// }
//
// 	println("-----------------ParameterOfGetDemo------------>>>")
//
// }

func ParamtersOfPostDemo() {
	println("<<<ParamtersOfPostDemo---------------------------")

	req, _ := http.NewRequest("POST", "http://www.google.com/search?q=foo&q=bar&both=x&prio=1&empty=not",
		strings.NewReader("z=post&both=y&prio=2&empty="))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	// Form contains the parsed form data, including both the URL
	// field's query parameters and the POST or PUT form data.
	// This field is only available after ParseForm is called.

	if qs := req.Form["q"]; !reflect.DeepEqual(qs, []string{"foo", "bar"}) {
		fmt.Printf(`req.Form["q"] = %q, want ["foo", "bar"]`, qs)
	}

	// FormValue returns the first value for the named component of the query.
	// POST and PUT body parameters take precedence over URL query string values.
	if prio := req.FormValue("prio"); prio != "2" {
		fmt.Printf(`req.FormValue("prio") = %q, want "2" (from body)`, prio)
	}

	// PostForm contains the parsed form data from POST or PUT
	// body parameters.
	// !!!This field is only available after ParseForm is called.!!!
	if bq, found := req.PostForm["q"]; found {
		fmt.Printf(`req.PostForm["q"] = %q, want no entry in map`, bq)
	}

	// PostFormValue returns the first value for the named component of the POST
	// or PUT request body. URL query parameters are ignored.
	if bz := req.PostFormValue("z"); bz != "post" {
		fmt.Printf(`req.PostFormValue("z") = %q, want "post"`, bz)
	}
	println("-----------------------------ParamtersOfPostDemo>>>")
}
