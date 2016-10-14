// Package main provides ...
package newDir

import (
	"fmt"
	"time"

	"gopkg.in/redis.v2"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	// c := ExampleNewClient()
	// ExampleClient(c)
	redisv4Pub()
	// time.Sleep(time.Second)
	// redisv4Sub()
	// redisv2Demo()
	println("-----------------------------JustDemo end>>>")
	return
}

// func ExampleNewClient() *redis.Client {
// 	client := redis.NewClient(&redis.Options{
// 		// Addr:     "192.168.8.129:6379",
// 		Addr:     "192.168.9.127:6379",
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})
//
// 	pong, err := client.Ping().Result()
// 	fmt.Println(pong, err)
// 	return client
// 	// Output: PONG <nil>
// }
//
// func ExampleClient(client *redis.Client) {
// 	err := client.Set("key", "value", 0).Err()
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	val, err := client.Get("key").Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("key", val)
//
// 	val2, err := client.Get("key2").Result()
// 	if err == redis.Nil {
// 		fmt.Println("key2 does not exists")
// 	} else if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("key2", val2)
// 	}
// 	// Output: key value
// 	// key2 does not exists
// }

func redisv2Demo() {
	println("//<<-------------------------redisv2Demo start-----------")
	start := time.Now()
	redisClient := redis.NewTCPClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", "192.168.8.136", 6379),
		Password: "",
		DB:       0,
	})
	errcmd := redisClient.Ping()
	if errcmd != nil {
	}

	// val, err := redisClient.Get("a").Result()
	// fmt.Println("key a ", val)
	// if err != nil {
	// 	panic(err)
	// }

	keyid := "id1-expire"
	expireTime := time.Second * 3
	err := redisClient.SetEx(keyid, expireTime, " 333333").Err()
	if err != nil {
		panic(err)
	}

	val, err := redisClient.Get(keyid).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keyid, "\t", val)

	// time.Sleep(time.Second)
	time.Sleep(expireTime)

	val, err = redisClient.Get(keyid).Result()
	if err != nil {
		fmt.Printf("val: %+v\n", val)
		fmt.Printf("err.Error(): %+v\n", err.Error())
		fmt.Printf("err: %T\n", err)
		// panic(err)
	}
	fmt.Println(keyid, "\t", val)

	fmt.Printf("redisv2Demo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------redisv2Demo end----------->>")
}

func redisv4Pub() {
	println("//<<-------------------------redisv2Demo start-----------")
	start := time.Now()
	redisClient := redis.NewTCPClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", "192.168.8.136", 6379),
		Password: "",
		DB:       0,
	})
	defer redisClient.Close()
	errcmd := redisClient.Ping()
	if errcmd != nil {
		panic("network error")
	}

	channel := "pubc1"
	msg := "content from channel "
	for {
		cmd := redisClient.Publish(channel, msg)
		if cmd.Err() != nil {
			panic("pub error")
		}

		fmt.Printf("pub v %+v,s %+v,e %+v,\n",
			cmd.Val(),
			cmd.String(),
			cmd.Err())
		time.Sleep(time.Second)
	}

	fmt.Printf("redisv2Demo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------redisv2Demo end----------->>")
}

func redisv4Sub() {
	println("//<<-------------------------redisv2Demo start-----------")
	start := time.Now()
	redisClient := redis.NewTCPClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", "192.168.8.136", 6379),
		Password: "",
		DB:       0,
	})
	errcmd := redisClient.Ping()
	if errcmd != nil {
		panic("network error")
	}

	channel := "pubc1"
	pubsub := redisClient.PubSub()
	defer pubsub.Close()
	pubsub.Subscribe(channel)

	var msg interface{}
	var err error
	for {
		msg, err = pubsub.ReceiveTimeout(time.Second)
		if err != nil {
			panic(err.Error())
		}
		switch v := msg.(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Payload)
			fmt.Printf("v.String(): %+v\n", v.String())
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			println("error")
			return
		}
	}
	fmt.Printf("redisv2Demo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------redisv2Demo end----------->>")
}
