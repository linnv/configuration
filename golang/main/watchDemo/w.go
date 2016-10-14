// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !plan9

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/fsnotify/fsnotify"
)

var dir = "/Users/Jialin/golang/src/demos/main/watchDemo/a/a"

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan struct{})
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			case <-done:
				return
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}
	go LogToFile(done)
	time.Sleep(time.Second * 10)
	close(done)
}

const (
	// Bits or'ed together to control what's printed. There is no control over the
	// order they appear (the order listed here) or the format they present (as
	// described in the comments).  A colon appears after these items:
	//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         = 1 << iota     // the date: 2009/01/23
	Ltime                         // the time: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

func LogToFile(done <-chan struct{}) {
	fileName := dir

	fmt.Printf("  fileName: %+v\n", fileName)
	os.MkdirAll(path.Dir(fileName), 0777)
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		return
	}
	defer f.Close()

	// os.Stderr.Write(append([]byte(""), "This is a test log entry"...))

	log.SetOutput(f)
	// log.SetFlags(log.Lshortfile | LstdFlags)
	log.SetFlags(log.Llongfile | LstdFlags)
	log.SetPrefix("[Error]")

	for {
		log.Println("This is a test log entry ", time.Now().String())
		time.Sleep(time.Second)
		select {
		case <-done:
			return
		}
	}
}
