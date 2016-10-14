// Go's _select_ lets you wait on multiple channel
// operations. Combining goroutines and channels with
// select is a powerful feature of Go.

package main

import "time"

type Item struct {
	ID int
	On bool
	// Name string
	Cap chan int
}

func Listening(ch *Item) {
	var i = 0
	for {
		i++
		// t := <-ch
		if ch.On {
			if i == 4 {
				println("end point ")
				break
			}
			ch.Cap <- i
			println("get chan on")
			time.Sleep(time.Second)
		}
	}
}

func Invoke(ch *Item) {
	// func Invoke(ch chan<- Item) {
	println("send it")
	// ch <- t
	time.Sleep(time.Second)
}

func moniter(ch *Item) {
	for {
		select {
		case n := <-ch.Cap:
			println("looking at ", n)
		}
	}
}

func main() {

	t := &Item{ID: 10, On: true, Cap: make(chan int)}
	// c := make(chan Item)
	go Listening(t)
	// go Invoke(t)
	go moniter(t)
	time.Sleep(time.Second * 10)
	println("good")
}

// func main() {
// 	var wg sync.WaitGroup
// 	runtime.GOMAXPROCS(2)
//
// 	itemChan := make(chan Item, 3)
// 	sigChan := make(chan bool, 1)
//
// 	wg.Add(2) //count of goroutines
// 	go func() {
// 		defer wg.Done()
// 		time.Sleep(time.Second * 1)
// 		itemChan <- Item{ID: 10, Name: "jialin  wu"}
// 		time.Sleep(time.Second * 2)
// 		itemChan <- Item{ID: 100, Name: "jialin  wu"}
// 		return
// 	}()
//
// 	go func() {
// 		defer wg.Done()
// 		isRunning := true
// 		for isRunning {
// 			println("waiting for chan receiving")
// 			select { //select will effect each channel only one time, thus use for loop to receive one chan n times
// 			case msg1 := <-itemChan:
// 				fmt.Println("received", msg1)
// 			case sig := <-sigChan:
// 				println("exit  safe ", sig)
// 				isRunning = false
// 			}
//
// 		}
// 	}()
// 	sig := utility.SystemSignalWaiter()
// 	fmt.Printf("sig: [%+v]  invoked\n", sig)
// 	Exit(&sigChan)
// 	wg.Wait()
// }

func Exit(sigChan *chan bool) {
	*sigChan <- true
	println("------ exit sig pause \n========================\n")
}

func ChanHandler() {
	//@TODO
	return
}
