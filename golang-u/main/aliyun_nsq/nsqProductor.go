package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/bitly/go-nsq"
)

type T struct {
	Ti time.Time
	Th int
}

func main() {
	config := nsq.NewConfig()
	// w, _ := nsq.NewProducer("127.0.0.1:4150", config)
	// w, _ := nsq.NewProducer("192.168.10.165:4150", config)

	count := flag.Int("count", 10, "count")
	flag.Parse()
	// err := q.ConnectToNSQD("182.92.72.126:4150")
	w, _ := nsq.NewProducer("182.92.76.126:4150", config)
	t := T{}
	// for i := 0; i < 100; i++ {
	for i := 0; i < *count; i++ {
		tn := time.Now()
		t.Ti = tn
		t.Th = i
		bs, err := json.Marshal(t)
		if err != nil {
			return
		}
		tn2 := time.Now()
		//@toDelete
		fmt.Printf("  Marshal consumes %+v Nanosecond()\n", tn2.Sub(tn).Nanoseconds())
		err = w.Publish("test", bs)
		// err = w.Publish("ssp_notify", bs)
		// err := w.Publish("test", []byte("msg from jialin's mbp"))
		if err != nil {
			log.Panic("Could not connect")

		}
		//@toDelete
		fmt.Printf("  send: %+v\n", t)
		// fmt.Printf(" th:%d  tn: %+v\n", th, tn)
		time.Sleep(time.Second)
	}
	//@toDelete
	fmt.Printf("  time.Now(): %+v\n", time.Now())

	w.Stop()
}
