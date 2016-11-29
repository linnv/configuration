package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	var pid int
	// fmt.Fscanf(os.Stdin, "%d", &pid)
	fmt.Fscan(os.Stdin, &pid)
	unixFiles := fmt.Sprintf("/tmp/unix_%d.sock", pid)
	fd, err := net.Dial("unix", unixFiles)
	if err != nil {
		panic(err.Error())
	}

	// _, err = fd.Write([]byte(byte(1)))
	_, err = fd.Write([]byte{byte(0x1)})
	if err != nil {
		panic(err.Error())
	}
	all, err := ioutil.ReadAll(fd)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(os.Stdout, "%s\n", string(all))
	fd.Close()

	// fmt.Fprintf(os.Stdout, "%s\n", unixFiles)
	// fmt.Fprintf(os.Stdout, "%d\n", pid)
	return
}

// 	l, err := net.Listen("unix", unixFiles)
// 	if err != nil {
// 		panic(err.Error())
// 	}
//
// 	sig := make(chan os.Signal, 1)
// 	signal.Notify(sig, os.Interrupt)
// 	go func() {
// 		<-sig
// 		os.Remove(unixFiles)
// 		os.Exit(1)
// 	}()
//
// 	buf := make([]byte, 1)
// 	for {
// 		fd, err := l.Accept()
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		if _, err := fd.Read(buf); err != nil {
// 			fmt.Fprint(os.Stderr, err.Error())
// 			continue
// 		}
// 		go handlerConnection(fd, buf)
// 	}
//
// }
//
// func handlerConnection(conn net.Conn, buf []byte) error {
// 	switch buf[0] {
// 	case 1:
// 		fmt.Fprint(os.Stdout, "good")
// 	default:
// 		fmt.Fprint(os.Stdout, "default")
// 	}
// 	_, err := conn.Write([]byte("good msg"))
// 	return err
// }
