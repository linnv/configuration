package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

type T struct {
	Ti time.Time
	Th int
}

func main() {
	config := nsq.NewConfig()
	topic := "ssp_mbv_notify"

	count := flag.Int("count", 10, "count")
	flag.Parse()
	// w, _ := nsq.NewProducer("192.168.10.41:4150", config)
	// w, err := nsq.NewProducer("cent.local:4150", config)
	w, err := nsq.NewProducer("192.168.100.27:4150", config)
	if err != nil {
		panic(err.Error())
	}

	t := T{}
	// for i := 0; i < 100; i++ {
	for i := 0; i < *count; i++ {
		tn := time.Now()
		t.Ti = tn
		t.Th = i
		bs, err := json.Marshal(t)
		if err != nil {
			panic(err.Error())
		}
		tn2 := time.Now()
		//@toDelete
		fmt.Printf("  Marshal consumes %+v Nanosecond()\n", tn2.Sub(tn).Nanoseconds())
		err = w.Publish(topic, bs)
		// err = w.Publish("ssp_notify", bs)
		// err := w.Publish("test", []byte("msg from jialin's mbp"))
		if err != nil {
			log.Panic("Could not connect %s", err.Error())

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
