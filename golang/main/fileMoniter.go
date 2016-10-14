package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var POLL_DURATION time.Duration

func init() {
	POLL_DURATION = 1 * time.Second
}

const (
	UPDATE = iota + 1
	EXIT
)

type FileMonitor struct {
	FilePath string
	action   chan int
	exit     chan bool
	wg       sync.WaitGroup
}

func NewFileMonitor(filePath string) *FileMonitor {
	return &FileMonitor{FilePath: filePath, action: make(chan int), exit: make(chan bool)}
}

func (fm *FileMonitor) Start() {
	var prevModTime time.Time
	for {
		select {
		case <-time.After(POLL_DURATION):
			finfo, err := os.Stat(fm.FilePath)
			if err != nil {
				panic(err.Error())
			}
			log.Printf("fm.FilePath: %+v\n", fm.FilePath)
			if finfo.IsDir() {
				os.Stdout.Write(append([]byte("dir"), '\n'))
			} else {
				modTime := finfo.ModTime()
				if modTime != prevModTime {
					prevModTime = modTime
					fm.action <- UPDATE
				}
			}
		case <-fm.exit:
			fm.action <- EXIT
			return
		}
	}
	panic("unreachable")
}

func (fm *FileMonitor) Stop() {
	fm.wg.Add(1)
	fm.exit <- true
	println("waiting exit")
	fm.wg.Wait()
}

func (fm *FileMonitor) Exit() {
	fm.exit <- true
}

func (fm *FileMonitor) ActionHandler() (err error) {
	for {
		select {
		case a := <-fm.action:
			switch a {
			case UPDATE:
				//@TODO  do something
				log.Println("update: works\n")
			case EXIT:
				println("permit exiting")
				fm.wg.Done()
				return nil
			default:
			}
		}
	}
}

func Wait(sig ...os.Signal) os.Signal {
	c := make(chan os.Signal, 1)
	if len(sig) == 0 {
		log.Printf("[]int{}: %+v\n", []syscall.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGCHLD, syscall.SIGSTOP})
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGCHLD, syscall.SIGSTOP)
	} else {
		signal.Notify(c, sig...)
	}
	// return <-c
	a := <-c
	log.Printf("issue signal: %+v\n", a)
	return a
}

func main() {
	fp := flag.String("file", "/Users/Jialin/golang/src/demos/main/a.go", "file to be monitored")
	flag.Parse()
	fm := NewFileMonitor(*fp)
	go fm.Start()
	go fm.ActionHandler()
	Wait()
	fm.Stop()
}
