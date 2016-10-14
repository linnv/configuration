// Package main provides ...
package newDir

import (
	"fmt"
	"time"

	as "github.com/aerospike/aerospike-client-go"

	"demos/demos/utility"
)

func aerospikeTestDemo() {
	println("//<<-------------------------aerospikeTestDemo start-----------")
	start := time.Now()
	client, err := as.NewClient("192.168.100.27", 3000)

	utility.CheckError(err)

	key, err := as.NewKey("ssp", "setOne",
		"key value goes here and can be any supported primitive")
	utility.CheckError(err)

	bin1 := as.NewBin("bin1", "value1")
	bin2 := as.NewBin("bin2", "value2")

	// Write a record
	err = client.PutBins(nil, key, bin1, bin2)
	utility.CheckError(err)

	// Read a record
	record, err := client.Get(nil, key)
	fmt.Printf("record: %+v\n", record)

	client.Close()
	fmt.Printf("aerospikeTestDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------aerospikeTestDemo end----------->>")
}

func JustDemo() {
	println("//<<-------------------------JustDemo start-----------")
	start := time.Now()

	fmt.Printf("JustDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------JustDemo end----------->>")
}
