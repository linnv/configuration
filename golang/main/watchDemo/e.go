// Package main provides ...
package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/hpcloud/tail"

	fsnotify "gopkg.in/fsnotify.v1"
)

var fn = "/Users/Jialin/golang/src/demos/main/watchDemo/a/a"

func main() {
	n := time.Now()
	// go tailFile()
	go updateFile()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Printf("%v,modified file:%s", n, event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(fn)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func updateFile() {
	file, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	i := 1
	for {
		i++
		_, err := file.Write([]byte(strconv.Itoa(i) + "a\n"))
		if err != nil {
			panic(err.Error())
		}
		time.Sleep(5 * time.Second)
	}
}

func tailFile() {
	t, err := tail.TailFile(fn, tail.Config{Follow: true})
	if err != nil {
		panic(err.Error())
	}
	for line := range t.Lines {
		log.Printf("line.Text: %+v\n", line.Text)
		// fmt.Println(line.Text)
	}
	// file, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer file.Close()
	// for {
	// 	_, err := file.Write([]byte(`a`))
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	time.Sleep(5 * time.Second)
	// }
}
