// Package main provides ...
package newDir

import (
	"fmt"
	"log"
	"time"
)

func TickDemo() {
	println("<<<TickDemo start---------------------------")
	// tick := time.Tick(time.Duration(2) * time.Second)

	// tick := time.Tick(int64(1) * time.Second)
	tick := time.Tick(time.Duration(1) * time.Second)
	for t := range tick {
		fmt.Printf("t: %+v\n", t)
	}
	println("-----------------------------TickDemo end>>>")
	return
}

type A struct {
	T    *time.Timer
	Name string `json:"Name"`
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	// fmt.Printf("time.Now(): %+v\n", time.Now().Unix())
	// fmt.Printf("time.Now() unixnano: %+v\n", time.Now().UnixNano())
	// str := strconv.FormatInt(time.Now().UnixNano(), 10)
	// fmt.Printf("  str: %+v\n", str)
	// filekey := str[0:13] + ".zip"
	// fmt.Printf("  filekey: %+v\n", filekey)
	//
	// t := time.Now().Local().Format("20060102")
	// fmt.Printf("  t: %+v\n", t)
	// //@toDelete
	// // fmt.Printf("now.b: %+v\n", now.BeginningOfDay())
	// // fmt.Printf("now.b: %+v\n", utils.ConvertTimeint2str(now.BeginningOfDay().Unix()-60*24*60))
	// nn := time.Now()
	//
	// du := time.Duration(-nn.Hour()) // count hours
	// fmt.Printf("du: %+v\n", du)
	// d := du * time.Hour //count durations(second?)
	// fmt.Printf("d: %+v\n", d)
	// tc := nn.Truncate(time.Hour) //hours left only
	// fmt.Printf("tc: %+v\n", tc)
	// tr := tc.Add(d)
	// fmt.Printf("tr: %+v\n", tr)
	//
	// n := rand.Intn(12)
	// fmt.Printf("n: %+v\n", n)
	// rr := time.Duration(-99)
	// fmt.Printf("rr: %+v\n", rr)

	a := new(A)
	a.T = time.AfterFunc(time.Second, func() {
		log.Println("timer: works")
	})
	println("-----------------------------JustDemo end>>>")
	return
}
