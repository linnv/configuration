// Package main provides ...
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func init() {
	ca, cb = make(chan int), make(chan int)
	// POLL_DURATION = 250 * time.Millisecond
	POLL_DURATION = 1 * time.Second
}

var POLL_DURATION time.Duration

var (
	ca, cb chan int
)

const filePath = "/Users/Jialin/golang/src/demo/main/a.t"

func Rec() {
	var prevModTime time.Time
	for {
		// if _, err := os.Stat(fw.Filename); err == nil {
		// 	return nil
		// } else if !os.IsNotExist(err) {
		// 	return err
		// }
		// fmt.Printf("  for time.Now(): %+v\n", time.Now())
		select {
		case <-time.After(POLL_DURATION):
			finfo, err := os.Stat(filePath)
			if err != nil {
				// no such file or dir
				//@toDelete
				fmt.Printf("  err.Error(): %+v\n", err.Error())
			}
			if finfo.IsDir() {
				// it's a directory
				os.Stdout.Write(append([]byte("dir"), '\n'))
			} else {
				modTime := finfo.ModTime()
				if modTime != prevModTime {
					prevModTime = modTime
					str, err := FileToStringDemo(filePath)
					if err != nil {
						fmt.Printf("  err.Error(): %+v\n", err.Error())
					}
					//@toDelete
					fmt.Printf("  str: %+v\n", str)
				}
			}
			//@toDelete
			fmt.Printf("  finfo: %+v\n", finfo)
			fmt.Printf("  coninue time.Now(): %+v\n", time.Now())
			continue
		case d := <-ca:
			fmt.Printf("a  d: %+v\n", d)
		case d := <-cb:
			fmt.Printf("b  d: %+v\n", d)
		}
	}
	panic("unreachable")
}

func send() {

}

func main() {
	go Rec()
	ca <- 111
	time.Sleep(time.Second)
	cb <- 111
	os.Stdout.Write(append([]byte("end in 10 seconeds"), '\n'))
	time.Sleep(time.Second * 30)

}

func FileToStringDemo(filePath string) (str string, err error) {
	os.Stdout.Write(append([]byte("open file"), '\n'))
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	thisb, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	str = string(thisb)
	return
}
