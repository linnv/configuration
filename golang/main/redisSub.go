// Package main provides ...
package main

import (
	"fmt"
	"time"

	"gopkg.in/redis.v2"
)

func main() {
	done := make(chan struct{})
	// go redisv4Pub(done)
	go redisv4Sub(done)
	time.Sleep(time.Second * 20)
	close(done)
	println("exiting")
	time.Sleep(time.Second * 2)
}

func redisv4Sub(done chan struct{}) {
	println("//<<-------------------------redisv2Demo start-----------")
	start := time.Now()
	redisClient := redis.NewTCPClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", "192.168.8.136", 6379),
		Password: "",
		DB:       0,
	})
	errcmd := redisClient.Ping()
	if errcmd.Err() != nil {
		panic("network error")
	}

	channel := "pubc1"
	pubsub := redisClient.PubSub()
	defer pubsub.Close()
	pubsub.Subscribe(channel)

	var msg interface{}
	var err error
	for {
		// msg, err = pubsub.ReceiveTimeout(time.Second)
		msg, err = pubsub.Receive()
		if err != nil {
			panic(err.Error())
		}
		// fmt.Printf("msg: %+v\n", msg)
		switch v := msg.(type) {
		case *redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Payload)
			fmt.Printf("v.String(): %+v\n", v.String())
		case *redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			println("error")
			return
		}

		select {
		case <-done:
			return
		default:

		}
	}
	fmt.Printf("redisv2Demo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------redisv2Demo end----------->>")
}
