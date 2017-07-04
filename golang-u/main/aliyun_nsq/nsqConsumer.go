package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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

func (in Int64Slice) ToString() (str string) {
	str += "["
	for i := 0; i < len(in); i++ {
		if i == 0 {
			str += strconv.FormatInt(in[i]/1000000, 10)
			continue
		}
		str += "," + strconv.FormatInt(in[i]/1000000, 10)
	}
	str += "]"
	return
}

var pageCount, count int

func main() {

	pageCount = 0
	countFlag := flag.Int("count", 10, "count")
	flag.Parse()
	count = *countFlag

	wg := &sync.WaitGroup{}
	wg.Add(1)

	timesConsumption = make(Int64Slice, 0, count)
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
		//@toDelete
		fmt.Printf("  receive: %+v\n", ti)
		//@toDelete
		fmt.Printf(" local time: %+v\n", tn)
		//@toDelete
		fmt.Printf("nsq consumes %+v nanoseconds\n", sub)
		timesConsumption = append(timesConsumption, sub)
		// log.Printf("Got a message: %s\n", string(message.Body))
		// wg.Done()
		if ti.Th >= count-1 {
			SortAndPrintTimeConsumed()
			timesConsumption = timesConsumption[:0]
			pageCount++
		}
		return nil

	}))
	// err := q.ConnectToNSQD("127.0.0.1:4150")
	// err := q.ConnectToNSQD("192.168.10.41:4150")
	err := q.ConnectToNSQD("182.92.76.126:4150")
	if err != nil {
		log.Panic("Could not connect")

	}
	wg.Wait()

}

func SortAndPrintTimeConsumed() {
	basic := "/Users/Jialin/web/demo/chart/samples/"
	absPath := basic + "sample_line.html"
	_, htmlTemp := FileToStringDemo(absPath)
	rStr := timesConsumption.ToString()
	replacer := strings.NewReplacer(
		"{{dataToFill}}", rStr,
	)
	str := replacer.Replace(htmlTemp)

	//@toDelete
	fmt.Printf("  count: %+v\n", count)
	// randname := "chart" + strconv.FormatInt(time.Now().UnixNano(), 10)[7:10] + ".html"
	randname := "chart" + strconv.Itoa(count) + strconv.Itoa(pageCount) + ".html"
	SaveStringToFile(str, basic+randname)
	//@toDelete
	fmt.Printf("  timesConsumption origin: %+v\n", timesConsumption)
	originStr := " timesConsumption original:" + timesConsumption.ToString() + "\n"
	logFile := basic + "nsq_" + strconv.Itoa(count) + ".stats"
	SaveStringToFile(originStr, logFile)
	//@toDelete
	sort.Sort(timesConsumption)
	sortStr := strconv.FormatInt(time.Now().Unix(), 10) + " timesConsumption sorted:" + timesConsumption.ToString() + "\n-------------------------\n"
	SaveStringToFile(sortStr, logFile)
	//@toDelete
	fmt.Printf("  timesConsumption sort: %+v\n", timesConsumption)
}

func FileToStringDemo(absPath string) (err error, str string) {
	f, err := os.Open(absPath)
	if err != nil {
		return

	}
	thisb, err := ioutil.ReadAll(f)

	return nil, string(thisb)
}

func SaveStringToFile(str, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err

	}
	defer f.Close()
	_, err = f.WriteString(str)
	f.Sync()
	return nil
}
