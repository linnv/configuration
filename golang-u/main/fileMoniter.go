package main

import (
	"log"
	"os"
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
}

func NewFileMonitor(filePath string) *FileMonitor {
	return &FileMonitor{FilePath: filePath, action: make(chan int), exit: make(chan bool)}
}

func (fm *FileMonitor) Start() error {
	go fm.ActionHandler()
	var prevModTime time.Time
	for {
		select {
		case <-time.After(POLL_DURATION):
			finfo, err := os.Stat(fm.FilePath)
			if err != nil {
				return err
			}
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
			return nil
		}
	}
	panic("unreachable")
	return nil
}

func (fm *FileMonitor) Stop() {
	fm.exit <- true
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
				return nil
			default:
			}
		}
	}
}

func main() {
	fp := "/Users/Jialin/golang/src/demo/main/a.t"
	fm := NewFileMonitor(fp)
	go fm.Start()
	time.Sleep(time.Second * 2)
	fm.Stop()
}
