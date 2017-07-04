// Package main provides ...
package demo

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/linnv/logx"
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

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	fmt.Printf("time.Now(): %+v\n", time.Now().Unix())
	fmt.Printf("time.Now() unixnano: %+v\n", time.Now().UnixNano())
	str := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Printf("  str: %+v\n", str)
	filekey := str[0:13] + ".zip"
	fmt.Printf("  filekey: %+v\n", filekey)

	t := time.Now().Local().Format("20060102")
	fmt.Printf("  t: %+v\n", t)
	//@toDelete
	// fmt.Printf("now.b: %+v\n", now.BeginningOfDay())
	// fmt.Printf("now.b: %+v\n", utils.ConvertTimeint2str(now.BeginningOfDay().Unix()-60*24*60))
	nn := time.Now()

	du := time.Duration(-nn.Hour()) // count hours
	fmt.Printf("du: %+v\n", du)
	d := du * time.Hour //count durations(second?)
	fmt.Printf("d: %+v\n", d)
	tc := nn.Truncate(time.Hour) //hours left only
	fmt.Printf("tc: %+v\n", tc)
	tr := tc.Add(d)
	fmt.Printf("tr: %+v\n", tr)

	n := rand.Intn(12)
	fmt.Printf("n: %+v\n", n)
	// rr := time.Duration(-99)
	// fmt.Printf("rr: %+v\n", rr)
	println("-----------------------------JustDemo end>>>")
	return
}

var days = [...]string{
	"7Sunday",
	"1Monday",
	"2Tuesday",
	"3Wednesday",
	"4Thursday",
	"5Friday",
	"6Saturday",
}

func FormatDemo() {
	println("//<<-------------------------FormatDemo start-----------")
	start := time.Now()

	time_ := "20161227073000"
	// const format = "20061102"

	const le = len("20160102")
	const format = "20060102"
	d, err := time.ParseInLocation(format, time_[:le], time.Local)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("d: %s\n", d.Format("2006-01-02"))
	// log.Printf("d: %+v\n", d.Format("yyyy-MM-dd"))
	// s := d.Local().Format("2006-01-02")
	// log.Printf("s: %+v\n", s)

	n := time.Now()
	nn := n.Weekday()
	logx.Debugf("n.we: %+v\n", nn)
	logx.Debugf("days[nn]: %+v\n", days[nn])
	fmt.Printf("FormatDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------FormatDemo end----------->>")
}

// func AliyunProcessDemo(params)type {
// println("//<<-------------------------AliyunProcessDemo start-----------")
// 	start := time.Now()
//
// 	fmt.Printf("AliyunProcessDemo costs  %d millisecons actually %v\n",time.Since(start).Nanoseconds()/1000000,time.Since(start))
// println("//---------------------------AliyunProcessDemo end----------->>")
// }
