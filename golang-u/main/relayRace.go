package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

var allOver chan bool

var TimeMonitor chan Durantion

var batonFlag chan BatonInfo
var batonStation []BatonInfo

type BatonInfo struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

func WaitingBaton() {
	for {
		select {
		case t := <-batonFlag:
			batonStation = append(batonStation, t)
			time.Sleep(time.Second)
			if t.Id == 4 {
				allOver <- true
			}
		}
	}
}

func main() {
	// Create an unbuffered channel
	baton := make(chan int)
	allOver = make(chan bool)
	timeSignal := make(chan Durantion)
	TimeMonitor = make(chan Durantion)
	batonFlag = make(chan BatonInfo) //if not initalize, deallock will occur
	batonStation = make([]BatonInfo, 0, 4)

	// First runner to his mark
	go Runner(baton, timeSignal)
	go TimeCaculator()
	go WaitingBaton()

	// Start the race
	baton <- 1
	timeSignal <- Durantion{}
	<-allOver
	<-allOver
	fmt.Println("baton info detail:\n")
	for k, v := range batonStation {
		fmt.Printf("%+v: %+v\n", k, v)
	}
	// Give the runners time to race
	// time.Sleep(500 * time.Millisecond)
}

type Durantion struct {
	Start int64
	End   int64
}

func TimeCaculator() {

	for {
		select {
		case g := <-TimeMonitor:
			log.Printf(": %v->nanoSecond:%v\n", g, (g.Start - g.End))
			//one billionth of second: 1/1 000 000 000
			log.Printf(": %v->nanoSecond:%v\n", g, (g.End - g.Start))
			allOver <- true
		}
	}
}

func Runner(baton chan int, timeSignalTransact chan Durantion) {
	var newRunner int
	// Wait to receive the baton
	runner := <-baton
	currentTime := <-timeSignalTransact
	if runner == 1 {
		currentTime.Start = time.Now().UnixNano()
		// time.Sleep(time.Second)
	}

	// Start running around the track
	fmt.Printf("Runner %d Running With Baton\n", runner)

	// New runner to the line
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton, timeSignalTransact)
	}

	// Running around the track
	// time.Sleep(100 * time.Millisecond)

	// Is the race over
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		currentTime.End = time.Now().UnixNano()
		TimeMonitor <- currentTime
		// ending time must get before send baton flag to channel,or time error margin will occur
		batonFlag <- BatonInfo{runner, "jialin" + strconv.Itoa(runner)}
		// currentTime.End = time.Now().UnixNano()
		// TimeMonitor <- currentTime
		return
	}

	// Exchange the baton for the next runner
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)
	baton <- newRunner
	batonFlag <- BatonInfo{runner, "jialin" + strconv.Itoa(runner)}
	timeSignalTransact <- currentTime
}
