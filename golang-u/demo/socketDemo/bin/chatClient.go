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
	conn, err := net.Dial("tcp", host)
	checkErr(err)
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
		log.Printf("client receives: %v\n", string(bs))
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
