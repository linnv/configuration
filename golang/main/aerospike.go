package main

import (
	"fmt"

	as "github.com/aerospike/aerospike-client-go"
	// "time"
	// "io"
	"log"
	// "os"
)

func logDemo() {
	println("<<<logDemo---------------------------")
	// var loger *log.Logger

	// func init() {
	// 	logFile, err := os.OpenFile("./log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
	// 	if err != nil {
	// 		log.Fatalln("log conf error:", err.Error())
	// 	}

	// 	loger = log.New(io.MultiWriter(logFile, os.Stdout), "", log.LstdFlags|log.Lshortfile)
	// }
	println("-----------------------------logDemo>>>")
	return
}

func addbimapDemo() {
	println("<<<addbimapDemo---------------------------")
	// binMap := as.BinMap{
	// 	"key": keys,
	// }
	// fmt.Printf("binMap: %+v\n", binMap)
	//
	// t := time.Now()
	//
	// var date string
	//
	// for i := 0; i < 10; i++ {
	//
	// 	client.Put(nil, key, binMap)
	//
	// 	date := t.AddDate(0, 0, i).Format(FORMAT)
	// 	binMap := as.BinMap{
	// 		"pv": 110,
	//		"c":  100,
	// 	}
	//
	// 	// Write a record
	// 	err = client.Add(nil, key, binMap)
	// 	if err != nil {
	// 		log.Printf("%+v", err)
	// 		return
	// 	}
	// 	fmt.Printf("binMap: %+v\n", binMap)
	//
	// }
	println("-----------------------------addbimapDemo>>>")
	return
}

const (
	FORMAT = "20060102"
)

func main() {
	client, err := as.NewClient("192.168.10.60", 3000)
	if err != nil {
		log.Printf("%+v", err)
		return
	}

	// log.Printf("node name: %v", client.GetNodeNames())

	//add data

	keys := "185"
	key, err := as.NewKey("ssp", "ordersum", keys)
	// binMap := as.BinMap{
	// 	"pv": 1803,
	// 	"c":  1800,
	// }
	// err = client.Add(nil, key, binMap)
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }

	// key, err := as.NewKey("ssp", "set", keys)

	fmt.Printf("key: %+v\n---end\n", key)
	fmt.Printf("v key: %v\n---end\n", key)

	// Read a record
	record, err := client.Get(nil, key)
	if err != nil {
		log.Printf("%+v", err)
		return
	}

	fmt.Printf("record: %+v\n", record)

	// log.Printf("all record %s\n", record.String())
	// log.Printf("record key %s\n", record.Key.String())

	// if value, ok := record.Bins["key"]; ok {
	// 	log.Println("bin key ", value)
	// }
	// if value, ok := record.Bins["all"]; ok {
	// 	log.Println(value)
	// 	log.Println("bin all ", value)
	// }

	// if value, ok := record.Bins[date]; ok {
	// 	log.Println(value)
	// }

	// date = t.AddDate(0, 0, 1).Format(FORMAT)
	// if value, ok := record.Bins[date]; ok {
	// 	log.Println(value)
	// }
	client.Close()
}
