// Package main provides ...
package demo

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")

	println("-----------------------------JustDemo end>>>")
	return
}

func ExampleGet() {
	// res, err := http.Get("http://www.google.com/robots.txt")
	res, err := http.Get("http://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  res.ContentLength: %+v bytes\n", res.ContentLength)
	// _, err = ioutil.ReadAll(res.Body)
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  len(rebots): %+v\n", len(robots))
	fmt.Printf("  cap(rebots): %+v\n", cap(robots))
	// fmt.Printf("%s", robots)
}

func ExampleGetFixLength() {
	// res, err := http.Get("http://www.google.com/robots.txt")
	res, err := http.Get("http://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  res.ContentLength: %+v bytes\n", res.ContentLength)
	// _, err = ioutil.ReadAll(res.Body)
	// robots := make([]byte, 0, res.ContentLength)
	robots := make([]byte, res.ContentLength)
	n, err := io.ReadFull(res.Body, robots)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  readfull n bytes: %+v\n", n)
	fmt.Printf("  len(rebots): %+v\n", len(robots))
	fmt.Printf("  cap(rebots): %+v\n", cap(robots))
	// fmt.Printf("%s", robots)
}

func ExampleGetFixLengthUsingBuffer() {
	// res, err := http.Get("http://www.google.com/robots.txt")
	res, err := http.Get("http://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  res.ContentLength: %+v bytes\n", res.ContentLength)
	// _, err = ioutil.ReadAll(res.Body)
	// robots := make([]byte, 0, res.ContentLength)
	// buffer := bytes.NewBuffer(make([]byte, 0, 65536))
	// buffer := bytes.NewBuffer(make([]byte, 0, 19639))
	buffer := bytes.NewBuffer(make([]byte, 0, 19640))
	io.Copy(buffer, res.Body)
	robots := buffer.Bytes()
	fmt.Printf("  len(rebots): %+v\n", len(robots))
	fmt.Printf("  cap(rebots): %+v\n", cap(robots))
	length := len(robots)
	var body []byte
	//are we wasting more than 10% space?
	if cap(robots) > (length + length/10) {
		body = make([]byte, length)
		copy(body, robots)
	} else {
		body = robots
	}
	fmt.Printf("  len(body): %+v\n", len(body))
	fmt.Printf("  cap(body): %+v\n", cap(body))
}

func TemplateDemo() {
	println("//<<-------------------------TemplateDemo start-----------")
	start := time.Now()

	hc := `
	hello {{.}}
	`
	var t = template.Must(template.New("name").Parse(hc))
	err := t.Execute(os.Stdout, "jialin")
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("TemplateDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------TemplateDemo end----------->>")
}

func toUnicode(s string) (r string) {
	p := []rune(s)
	for i := 0; i < len(p); i++ {
		r += "%u" + fmt.Sprintf("%U", p[i])[2:]
	}
	return r
}

func UrlDemo() {
	println("//<<-------------------------UrlDemo start-----------")
	start := time.Now()
	// a := make(url.Values)
	// v := []rune("安庆")
	v := "安庆"
	r := toUnicode("安庆")
	// r := "%u" + fmt.Sprintf("%U", v[0])[2:]
	// r := strconv.QuoteRuneToASCII(v[0])
	// r = url.QueryEscape(r)
	log.Printf("v: %+v\n", v)
	log.Printf("r: %s\n", r)
	fmt.Printf("UrlDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------UrlDemo end----------->>")
}

func SecureHeadersDemo() {
	println("//<<-------------------------SecureHeadersDemo start-----------")
	start := time.Now()

	const (
		stsHeader           = "Strict-Transport-Security"
		stsSubdomainString  = "; includeSubdomains"
		stsPreloadString    = "; preload"
		frameOptionsHeader  = "X-Frame-Options"
		frameOptionsValue   = "DENY"
		contentTypeHeader   = "X-Content-Type-Options"
		contentTypeValue    = "nosniff"
		xssProtectionHeader = "X-XSS-Protection"
		xssProtectionValue  = "1; mode=block"
		cspHeader           = "Content-Security-Policy"
		hpkpHeader          = "Public-Key-Pins"
	)
	println(xssProtectionHeader, ":", xssProtectionValue)
	println(stsHeader, ":", stsSubdomainString, stsPreloadString) //@Todo
	println(frameOptionsHeader, ":", frameOptionsValue)
	println(contentTypeHeader, ":", contentTypeValue)

	fmt.Printf("SecureHeadersDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------SecureHeadersDemo end----------->>")
}
