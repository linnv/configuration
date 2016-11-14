package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
)

func main() {
	unixFiles := fmt.Sprintf("/tmp/unix_%d.sock", os.Getpid())
	l, err := net.Listen("unix", unixFiles)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(os.Stdout, "file id %d\n", os.Getpid())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		os.Remove(unixFiles)
		l.Close()
		os.Exit(1)
	}()

	buf := make([]byte, 1)
	for {
		fd, err := l.Accept()
		if err != nil {
			panic(err.Error())
		}
		if _, err := fd.Read(buf); err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			continue
		}
		handlerConnection(fd, buf)
		fd.Close()
	}

}

func handlerConnection(conn net.Conn, buf []byte) error {
	switch buf[0] {
	case 1:
		fmt.Fprint(os.Stdout, "good")
	default:
		fmt.Fprint(os.Stdout, "default")
	}
	_, err := conn.Write([]byte("good msg"))
	return err
}
