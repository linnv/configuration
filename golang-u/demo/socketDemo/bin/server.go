package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

var closeAvailable chan bool

func main() {
	ip := "127.0.0.1"
	port := 6797
	host := fmt.Sprint(ip, ":", port)
	master, err := net.Listen("tcp", host)
	if err != nil {
		log.Println("listen tcp connection err:%s", err.Error())
		return
	}
	defer master.Close()
	log.Printf("master.Addr(): %+v\n", master.Addr())
	closeAvailable = make(chan bool)
	for {
		conn, err := master.Accept()
		if err != nil {
			log.Println("accept connection err:%s", err.Error())
			return
		}
		go HandlerClientRead(conn)
		go HandlerClient(conn)
	}

	fmt.Printf("host: %+v\n", host)
}

func HandlerClientRead(conn net.Conn) error {
	bs := make([]byte, 1024)
	_, err := conn.Read(bs)
	if err != nil {
		return err
	}
	log.Printf("read from client %s \n", string(bs))
	closeAvailable <- true
	return nil
}

func HandlerClient(conn net.Conn) error {
	for {
		select {
		case <-closeAvailable:
			nowstr := time.Now().String()
			_, err := conn.Write([]byte("haha from server " + nowstr))
			log.Println("closing client connection")
			conn.Close()
			return err
		}
	}
}
