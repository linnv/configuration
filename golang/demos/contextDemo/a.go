// Package main provides ...
package newDir

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/context"
)

func A(done <-chan struct{}, data chan interface{}, index int) {
	for {
		select {
		case <-done:
			println("exiting go ", index)
			return
		case d := <-data:
			log.Printf("d: %+v\n", d)
		}
	}
}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan error, 1)
	//do http request and deal with request result in f
	go func() { c <- f(client.Do(req)) }()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		return err
	}
}

// func handleSearch(w http.ResponseWriter, req *http.Request) {
// 	// ctx is the Context for this handler. Calling cancel closes the
// 	// ctx.Done channel, which is the cancellation signal for requests
// 	// started by this handler.
// 	var (
// 		ctx    context.Context
// 		cancel context.CancelFunc
// 	)
// 	timeout, err := time.ParseDuration(req.FormValue("timeout"))
// 	if err == nil {
// 		// The request has a timeout, so create a context that is
// 		// canceled automatically when the timeout expires.
// 		ctx, cancel = context.WithTimeout(context.Background(), timeout)
// 	} else {
// 		ctx, cancel = context.WithCancel(context.Background())
// 	}
// 	defer cancel() // Cancel ctx as soon as handleSearch returns.
// }

func CtxGetDemo() {
	println("//<<-------------------------CtxGetDemo start-----------")
	// func MainDemo() {
	// 	port := ":9999"
	// 	http.HandleFunc("/", Root)
	// 	log.Printf("listening %s\n", port)
	// 	http.ListenAndServe(port, nil)
	// }
	//
	// func Root(w http.ResponseWriter, r *http.Request) {
	// 	time.Sleep(time.Second * 1)
	// 	w.Write([]byte("good response"))
	// 	// e := r.FormValue("e")
	// 	// e = strings.TrimSpace(e)
	// }
	start := time.Now()
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	url := "http://127.0.0.1:9999/?e='true'"
	// data:='''
	// req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic(err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	timeout := time.Millisecond * 500
	// timeout, err := time.ParseDuration(req.FormValue("timeout"))
	// if err == nil {
	// The request has a timeout, so create a context that is
	// canceled automatically when the timeout expires.
	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	// } else {
	// 	ctx, cancel = context.WithCancel(context.Background())
	// }
	defer cancel() // Cancel ctx as soon as handleSearch returns.

	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		//@TODO resp.Body
		// fmt.Printf("resp.b: %+v\n",resp.b)
		bs, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("bs: %+v\n", string(bs))
		// if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		// 	return err
		// }

		return nil
	})

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("CtxGetDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------CtxGetDemo end----------->>")
}

func CTXDemo() {
	println("//<<-------------------------CTXDemo start-----------")
	start := time.Now()
	// var results Results

	//@TODO
	var req *http.Request
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		// The request has a timeout, so create a context that is
		// canceled automatically when the timeout expires.
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel() // Cancel ctx as soon as handleSearch returns.

	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		//@TODO resp.Body
		// fmt.Printf("resp.b: %+v\n",resp.b)
		bs, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("bs: %+v\n", bs)
		// if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		// 	return err
		// }

		return nil
	})
	// httpDo waits for the closure we provided to return, so it's safe to
	fmt.Printf("CTXDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------CTXDemo end----------->>")
}

func CtxDeadLineDemo() {
	println("//<<-------------------------CtxDeadLineDemo start-----------")
	start := time.Now()
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(500*time.Millisecond))
	time.Sleep(time.Second)
	select {
	case <-ctx.Done():
		log.Println("deadline exceeded done: works")
	default:
		log.Println("default: works")
	}
	cancel()
	fmt.Printf("CtxDeadLineDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------CtxDeadLineDemo end----------->>")
}

func CtxTimeOutDemo() {
	println("//<<-------------------------CtxTimeOutDemo start-----------")
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	time.Sleep(time.Second)
	select {
	case <-ctx.Done():
		log.Println("timeout done: works")
	default:
		log.Println("default: works")
	}
	cancel()
	fmt.Printf("CtxTimeOutDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------CtxTimeOutDemo end----------->>")
}

//coatzConjecture implements ...
func coatzConjecture(ctx context.Context) {
	i := 1859
	fmt.Println(i)
	for i != 1 {
		if i%2 == 0 {
			i := i / 2
			fmt.Println("i/2\t=", i)
		} else {
			i := 3*i + 1
			fmt.Println("3i + 1\t=", i)
		}
	}
}

func ollatzConjectureDemo() {
	println("//<<-------------------------ollatzConjectureDemo start-----------")
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	go coatzConjecture(ctx)
	time.Sleep(time.Second)
	select {
	case <-ctx.Done():
		println("good")
	default:
		println("bad: time out")
	}
	cancel()
	fmt.Printf("ollatzConjectureDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------ollatzConjectureDemo end----------->>")
}

func CtexCancelDemo() {
	println("//<<-------------------------CtexCancelDemo start-----------")
	start := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	select {
	case <-ctx.Done():
		log.Println("done: works")
	default:
		log.Println("default: works")
	}
	fmt.Printf("CtexCancelDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------CtexCancelDemo end----------->>")
}

func JustDemo() {
	println("//<<-------------------------JustDemo start-----------")
	start := time.Now()
	// done := make(chan struct{})
	// data := make(chan interface{})
	// for i := 0; i < 3; i++ {
	// 	go A(done, data, i)
	// }
	// close(done)

	fmt.Printf("JustDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------JustDemo end----------->>")
}
