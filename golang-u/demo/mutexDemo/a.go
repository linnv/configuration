// Package main provides ...
package demo

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
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer rwMutex.RUnlock()
		rwMutex.RLock() //if you want to update i, you must invoke lock(but not rLock) wich will write after all reading is finished
		log.Printf("i: %+v read\n", i)
		time.Sleep(time.Second * 2)
		log.Printf("i: %+v read end\n", i)
	}()

	go func() {
		defer wg.Done()
		defer rwMutex.RUnlock()
		time.Sleep(time.Second * 1)
		rwMutex.RLock() //if you want to update i, you must invoke lock(but not rLock) wich will write after all reading is finished
		log.Printf("i: %+v read by second goroutine\n", i)
		time.Sleep(time.Second * 4)
		log.Printf("i: %+v read end second goroutine\n", i)
	}()
	time.Sleep(2 * time.Second)
	rwMutex.Lock() // it will be blocked,if read lock is working
	log.Println("write: works")
	i = 1
	rwMutex.Unlock()
	log.Println("write: works end")
	log.Printf("i: %+v last \n", i)
	wg.Wait()

	println("-----------------------------JustDemo end>>>")
	return
}
