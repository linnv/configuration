package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	ip := "127.0.0.1"
	port := 6797
	host := fmt.Sprint(ip, ":", port)
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Println("dial tcp connection err:%s", err.Error())
		return
	}
	defer conn.Close()

	nowstr := time.Now().String()
	_, err = conn.Write([]byte("msg from client " + nowstr))
	if err != nil {
		log.Println("client occurs err whil writing to server, err: ", err.Error())
		return
	}

	time.Sleep(time.Second)

	bs := make([]byte, 1024)
	_, err = conn.Read(bs)
	if err != nil {
		log.Println("client occurs err whil reading to server, err: ", err.Error())
		return
	}

	log.Printf("%s \n", string(bs))
}
