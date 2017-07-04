package main

import (
	"fmt"
	"log"
	"net"
)

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	ip := "127.0.0.1"
	port := 6797
	host := fmt.Sprint(ip, ":", port)
	hubListener, err := net.Listen("tcp", host)
	checkErr(err)
	exitChan := make(chan struct{})
	log.Printf("listen on : %+v\n", port)
	for {
		oneClient, err := hubListener.Accept()
		checkErr(err)
		select {
		case <-exitChan:
			hubListener.Close()
			return
		default:
		}
		go HandlerRead(oneClient, exitChan)
		go HandlerWrite(oneClient, exitChan)
	}
}

func HandlerRead(conn net.Conn, exit chan struct{}) {
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
			close(exit)
		}
		log.Printf("server receives: %v\n", string(bs))
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
	}
}
