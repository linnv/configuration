// Package main provides ...
package newDir

import (
	"log"
	"sync"
	"time"
)

//playground:`https://play.golang.org/p/ddKJNeLQwN`
func RWMutexDemo() {
	println("<<<JustDemo start---------------------------")
	var rwMutex sync.RWMutex
	i := 0
	c := make(chan struct{})
	go func() {
		defer rwMutex.RUnlock()
		rwMutex.RLock() //if you want to update i, you must invoke lock(but not rLock) wich will write after all reading is finished
		log.Printf("i: %+v read\n", i)
		time.Sleep(time.Second * 3)
		log.Printf("i: %+v read end\n", i)
		close(c)
	}()
	time.Sleep(time.Second)
	rwMutex.Lock() // it will be blocked,if read lock is working
	log.Println("write: works")
	i = 1
	rwMutex.Unlock()
	log.Println("write: works end")
	log.Printf("i: %+v last \n", i)
	<-c

	println("-----------------------------JustDemo end>>>")
	return
}
