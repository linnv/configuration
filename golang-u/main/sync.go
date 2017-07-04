package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	urls := []string{
		// "http://www.reddit.com/r/aww.json",
		// "http://www.reddit.com/r/funny.json",
		// "http://www.reddit.com/r/programming.json",

		"http://www.baidu.com",
		"http://www.360.cn",
		"http://www.zhihu.com",
	}
	jsonResponses := make(chan string)
	// jsonResponses := make(chan string, 3)

	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			} else {
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				now := time.Now().String()
				if err != nil {
					jsonResponses <- ("httpError" + now + url + "------end---->\n\n")
					log.Fatal(err)
				} else {
					// fmt.Printf("get url done: %+v\n", url)
					jsonResponses <- (string(body) + now + url + "------end---->\n\n")
				}
			}
		}(url)
	}

	go func() {
		defer wg.Done()
		// for _ = range jsonResponses {
		for response := range jsonResponses {
			// _ := response
			fmt.Println(response)
		}
	}()

	wg.Wait()
}
