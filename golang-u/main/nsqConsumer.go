package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/bitly/go-nsq"
)

type T struct {
	Ti time.Time
	Th int
}

type Int64Slice []int64

var timesConsumption Int64Slice

func (in Int64Slice) Less(i, j int) bool {
	return in[i] < in[j]
}

func (in Int64Slice) Len() int {
	return len(in)
}

func (in Int64Slice) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// timesConsumption := make([]int64, 0, 100)
	timesConsumption = make(Int64Slice, 0, 100)
	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("test", "test", config)
	// q, _ := nsq.NewConsumer("ssp_notify", "ssp_rocket_bid-192.168.10.187-7088", config)
	// q, _ := nsq.NewConsumer("ssp_notify", "ssp_rocket_bid-192.168.10.187-7088", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		// log.Printf("Got a message: %v", message)
		ti := T{}
		tn := time.Now()
		err := json.Unmarshal(message.Body, &ti)
		if err != nil {
			panic(err.Error())
		}
		tn2 := time.Now()
		//@toDelete
		fmt.Printf("  Unmarshal consumes %+v Nanoseconds()\n", tn2.Sub(tn).Nanoseconds())
		sub := tn.Sub(ti.Ti).Nanoseconds()
		timesConsumption = append(timesConsumption, sub)
		//@toDelete
		fmt.Printf("  receive: %+v\n", ti)
		//@toDelete
		fmt.Printf(" local time: %+v\n", tn)
		//@toDelete
		fmt.Printf("nsq consumes %+v nanoseconds\n", sub)
		// log.Printf("Got a message: %s\n", string(message.Body))
		// wg.Done()
		if ti.Th >= 99 {
			SortAndPrintTimeConsumed()
		}
		return nil

	}))
	// err := q.ConnectToNSQD("127.0.0.1:4150")
	err := q.ConnectToNSQD("192.168.10.41:4150")
	if err != nil {
		log.Panic("Could not connect")

	}
	wg.Wait()

}

func SortAndPrintTimeConsumed() {
	sort.Sort(timesConsumption)
	//@toDelete
	fmt.Printf("  timesConsumption: %+v\n", timesConsumption)
}
