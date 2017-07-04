// Go's _select_ lets you wait on multiple channel
// operations. Combining goroutines and channels with
// select is a powerful feature of Go.

package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"time"

	"demo/utility"
)

type Item struct {
	ID   int
	Name string
}

func tmpFromNetDemo() {
	println("<<<tmpFromNetDemo start---------------------------")

	cases := make([]reflect.SelectCase, len(chans))
	for i, ch := range chans {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}
	chosen, value, ok := reflect.Select(cases)
	// # ok will be true if the channel has not been closed.
	ch := chans[chosen]
	msg := value.String()
	println("-----------------------------tmpFromNetDemo end>>>")
	return
}
func main() {
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(2)

	itemChan := make(chan Item, 3)
	sigChan := make(chan bool, 1)

	wg.Add(2) //count of goroutines
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		itemChan <- Item{ID: 10, Name: "jialin  wu"}
		time.Sleep(time.Second * 2)
		itemChan <- Item{ID: 100, Name: "jialin  wu"}
		return
	}()

	go func() {
		defer wg.Done()
		isRunning := true
		for isRunning {
			println("waiting for chan receiving")
			select { //select will effect each channel only one time, thus use for loop to receive one chan n times
			case msg1 := <-itemChan:
				fmt.Println("received", msg1)
			case sig := <-sigChan:
				println("exit  safe ", sig)
				isRunning = false
			}

		}
	}()
	sig := utility.SystemSignalWaiter()
	fmt.Printf("sig: [%+v]  invoked\n", sig)
	Exit(&sigChan)
	wg.Wait()
}

func Exit(sigChan *chan bool) {
	*sigChan <- true
	println("------ exit sig pause \n========================\n")
}

func ChanHandler() {
	//@TODO
	return
}
