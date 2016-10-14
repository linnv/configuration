package main

import (
	"fmt"
	"sync"
	"time"
)

type D struct {
	Arr int
}

func main() {
	loopDemo()
	// ArrayPinter_PointerArrarDemo()
	// ArrayPointerDemo()
	// PointerArrayDemo()
}

func Iplus() {

	s := make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		s[8*i] = i * 10
	}

}

func ArrayPinter_PointerArrarDemo() {
	println("<<<-----------------ArrayPinter_PointerArrarDemo----------")
	println("-------copy value------")
	//modification will work during range,copy value?
	a := make([]D, 0, 1)
	a = append(a, D{Arr: 1})
	a = append(a, D{Arr: 134324})

	for k, v := range a {
		fmt.Printf("vxxx%+v: %+v arr:%p\n", k, v, &v.Arr)
		v.Arr = 191
	}

	for k, v := range a {
		fmt.Printf("vooo%+v: %+v arr:%p\n", k, v, &v.Arr)
	}
	// modification will work during range,copy address?
	println("-------copy address------")
	b := make([]*D, 0, 1)
	b = append(b, &D{Arr: 1})
	b = append(b, &D{Arr: 134324})

	for k, v := range b {
		fmt.Printf("pxxx%+v: %+v arr:%p\n", k, v, &v.Arr)
		v.Arr = 191
	}

	for k, v := range b {
		fmt.Printf("pooo%+v: %+v arr:%p\n", k, v, &v.Arr)
	}
	println("-----------------ArrayPinter_PointerArrarDemo------------>>>")
}

type PostData struct {
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	CreateDate time.Time `json:"create_date"`
	Hit        int       `json:"hit"`
	Content    string    `json:"content"`
	Source     string    `json:"source"`
	Url        string    `json:"url"`
}

type PostBody struct {
	Key  string      `json:"key"`
	Data *[]PostData `json:"data"`
}

type PostBodyNew struct {
	Key  string      `json:"key"`
	Data []*PostData `json:"data"`
}

var key = "65b4780597f3c3d23e629f8b1d30002f"

func ArrayPointerDemo() {
	println("<<<-----------------ArrayPointerDemo----------")
	pd := PostBody{Key: "xxx",
		Data: &[]PostData{PostData{Title: "demotitle", Author: "xxxx", CreateDate: time.Now(), Hit: 1999, Source: "xxxxxsource", Url: "fewwwurl "},
			PostData{Title: "demotitle", Author: "xxxx", CreateDate: time.Now(), Hit: 1999, Source: "xxxxxsource", Url: "fewwwurl "}}}

	fmt.Printf("before pd: %+v\n", pd.Data)

	pd.Key = key
	cd := time.Now()
	tmpd := pd.Data
	for k, _ := range *pd.Data { //addre not support range or index
		(*pd.Data)[k].CreateDate = cd
	}
	fmt.Printf("afer tmpd: %+v\n", tmpd)

	println("-----------------ArrayPointerDemo------------>>>")
}

func PointerArrayDemo() {
	println("<<<-----------------PointerArrayDemo----------")
	pd := PostBodyNew{Key: "xxx",
		Data: []*PostData{&PostData{Title: "demotitle", Author: "xxxx", CreateDate: time.Now(), Hit: 1999, Source: "xxxxxsource", Url: "fewwwurl "},
			&PostData{Title: "demotitle", Author: "xxxx", CreateDate: time.Now(), Hit: 1999, Source: "xxxxxsource", Url: "fewwwurl "}}}
	fmt.Printf("before pd:")
	pd.Key = key
	cd := time.Now()
	for k, v := range pd.Data { //addre not support range or index
		fmt.Printf("%+v: %+v\n", k, v)
		v.CreateDate = cd
		v.Title = "damn it"
	}
	fmt.Printf("afer tmpd:")
	for k, v := range pd.Data {
		fmt.Printf("%+v: %+v\n", k, v)
	}
	println("-----------------PointerArrayDemo------------>>>")
}

func loopDemo() {
	println("<<<loopDemo---------------------------")
	// the value of the iteration variable v changes at each step so the most probable output is first value of a same being greeted three times. You must capture the current value of v
	//v的值在每一次Loop中是改变的，但是每个goroutine传入的都同一个变量v，在不同的goroutine中传入同一个变量(地址一样)，所获得的这个变量的值是不会发生更新的
	var wg sync.WaitGroup
	a := []string{"3", "4r3r", "nbn"}
	for k, v := range a {
		fmt.Printf("%+v: %+v addr %p\n", k, v, &v)
		//method one: copy v to p,thus p is a new variable with new address and the updated value of v
		p := a
		fmt.Printf("p: %p\n", &p)
		wg.Add(1)
		//method two: invoke goroutine with one parameter, as the parameter has its address and value,each time v is put into goroutine as paramter,its value will be assgined to the parameter variable (s of func)
		go func(s string) {
			fmt.Printf("goroutine %+v: %+v addr %p\n", k, s, &s)
			wg.Done()
		}(v)
	}
	wg.Wait()
	println("-----------------------------loopDemo>>>")
	return
}
