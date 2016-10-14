// Package main provides ...
package newDir

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	UrlValicationDemo()
	println("-----------------------------JustDemo end>>>")
	return
}

func ExampleGet() {
	// res, err := http.Get("http://www.google.com/robots.txt")
	mateURL := "https://jialinwu.com"
	// mateURL := "http://127.0.0.1"
	res, err := http.Get(mateURL)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	defer io.Copy(ioutil.Discard, res.Body) // ensure EOF for keep-alive
	if res.StatusCode != 200 {
		return
	}

	fmt.Printf("  res.ContentLength: %+v bytes\n", res.ContentLength)
	// _, err = ioutil.ReadAll(res.Body)
	robots, err := ioutil.ReadAll(res.Body)
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
	fmt.Printf("res.ContentLength: %+v bytes\n", res.ContentLength)
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

	fmt.Printf("res.ContentLength: %+v bytes\n", res.ContentLength)
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

func UrlValicationDemo() {
	println("//<<-------------------------urlValicationDemo start-----------")
	start := time.Now()
	fmt.Println(url.Parse("a b"))
	fmt.Printf("urlValicationDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------urlValicationDemo end----------->>")
}

func DNSResolvTimeDemo() *net.TCPAddr {
	println("//<<-------------------------DNSResolvTimeDemo start-----------")
	start := time.Now()

	t0 := time.Now() // before dns resolution
	host, port := "jialinwu.com", "443"
	raddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("unable to resolve host: %v", err)
	}
	//@toDelete
	fmt.Printf("raddr: %+v\n", raddr)
	//@toDelete
	fmt.Printf("time consumes : %+v\n", time.Since(t0))
	fmt.Printf("DNSResolvTimeDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------DNSResolvTimeDemo end----------->>")
	return raddr
}

func TCPConnectionTimeDemo() {
	println("//<<-------------------------TCPConnectionTimeDemo start-----------")
	start := time.Now()
	raddr := DNSResolvTimeDemo()
	TCPConnectionStart := time.Now() // after dns resolution, before connect
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Fatalf("unable to connect to host %vv %v", raddr, err)
	}
	fmt.Printf("connected to %s\n", raddr.String())
	fmt.Printf("tcp connect consumes %+v\n", time.Since(TCPConnectionStart))
	scheme := "https"
	if scheme == "https" {
		// https use tls handshake
		handshakeStart := time.Now()
		tlsConn := tls.Client(conn, &tls.Config{
			ServerName:         raddr.IP.String(),
			InsecureSkipVerify: true,
		})
		if err := tlsConn.Handshake(); err != nil {
			panic(err.Error())
		}
		fmt.Printf("tls handshake consumes: %+v\n", time.Since(handshakeStart))
	}

	fmt.Printf("TCPConnectionTimeDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------TCPConnectionTimeDemo end----------->>")
}
