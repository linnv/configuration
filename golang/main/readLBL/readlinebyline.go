// Package main provides ...
package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

var bufs []bytes.Buffer

// var urls = make([]string, 774)
const workers = 6
const partion = 124

var urls = make([]string, 0, 774)

func readLine(path string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		count++
		// urls[count] = scanner.Text()
		urls = append(urls, scanner.Text())
		// if count %129==0 {
		// }
		// fmt.Println(scanner.Text())
	}
	println("there are ", count, " lines")
}

var mtx sync.Mutex

const path = "/Users/Jialin/myGit/OpenDemo/golang/main/readLBL/magnets"

func SaveTofiles(indexWorker int) error {
	mtx.Lock()
	defer mtx.Unlock()
	return ioutil.WriteFile(path, bufs[indexWorker].Bytes(), os.ModePerm)
}

const timeout = 20

var wg sync.WaitGroup

func saveFiles(indexWorker int, url []string) {
	defer wg.Done()
	// log.Printf("indexWorker: %+v\n", indexWorker)
	// log.Printf("len(url): %+v\n", len(url))
	// if indexWorker == 6 {
	// 	// log.Printf("url: %+v\n", url)
	// 	// log.Printf("url(len(url)): %+v\n", url[50:])
	// 	// return
	// }
	// log.Printf("url(len(url)): %+v\n", url[len(url)-1])

	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: timeout * time.Second,
		}).Dial,
		TLSHandshakeTimeout: timeout * time.Second,
	}

	hcl := http.Client{Timeout: timeout * time.Second,
		Transport: netTransport,
	}
	cal := 0
	for _, v := range url {
		log.Printf("getting by worker %d  v: %+v\n", indexWorker-1, v)
		gotCal++ //issue
		r, err := hcl.Get(v)
		if err != nil {
			log.Printf("worker %d break at index %d err.Error(): %+v\n", indexWorker-1, cal, err.Error())
			return
		}
		res, err := ioutil.ReadAll(r.Body)
		_, err = bufs[indexWorker-1].Write(res)
		if err != nil {
			log.Printf("worker %d break at index %d err.Error(): %+v\n", indexWorker-1, cal, err.Error())
			return
		}
		cal++
	}
	return
}

var gotCal = 0

func main() {
	bufs = make([]bytes.Buffer, workers)
	readLine("/Users/Jialin/originalGit/myDemo/python/webpageVister/544la/s.sh0")

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go saveFiles(i, urls[((i-1)*partion):(i*partion)])
	}
	println("get all pages done good")
	wg.Wait()
	for i := 0; i < workers; i++ {
		SaveTofiles(i)
	}
	println("go count", gotCal)
	println("good")
}
