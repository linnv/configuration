package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/linnv/logx"
)

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

var start time.Time

func main() {
	ip := "127.0.0.1"
	// ip := "192.168.11.16"
	// port := 60001
	port := 6798
	host := fmt.Sprint(ip, ":", port)
	conn, err := net.Dial("tcp", host)
	checkErr(err)
	start = time.Now()
	exit := make(chan struct{})
	log.Printf("connect to : %+v\n", host)
	go HandlerRead(conn, exit)
	go HandlerWrite(conn, exit)

	<-exit
}

func HandlerRead(conn net.Conn, exit chan struct{}) {
	//@TODO make slice base on length of msg server sent
	bs := make([]byte, 1024)
	for {
		select {
		case <-exit:
			return
		default:
		}
		_, err := conn.Read(bs)
		if err != nil {
			log.Println("connection lose: works")
			select {
			case <-exit:
				return
			default:
				close(exit)
				return
			}
		}
		if strings.Index(string(bs), "rtt") > -1 {
			fmt.Printf("pingServerDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
		}
		logx.Debugf("client receives: %v\n from %s", string(bs), conn.RemoteAddr().String())
	}
}

func HandlerWrite(conn net.Conn, exit chan struct{}) {
	var s string
	for {
		select {
		case <-exit:
			return
		default:
		}
		fmt.Scanln(&s)
		if s == "exit" {
			select {
			case <-exit:
				return
			default:
				close(exit)
				return
			}
		}
		// 	else if s == "ping" {
		// start := time.Now()
		// net.
		// fmt.Printf("pingServerDemo costs  %d millisecons actually %v\n",time.Since(start).Nanoseconds()/1000000,time.Since(start))
		// 	}

		start = time.Now()
		_, err := conn.Write([]byte(s))
		if err != nil {
			log.Println("connection lose: works")
			select {
			case <-exit:
				return
			default:
				close(exit)
				return
			}
		}
		logx.Debugf("write to: %+v\n", conn.RemoteAddr().String())
	}
}
