package demo

import (
	"fmt"
	"net/http"
	"strconv"
)

//just put a parameter to closure function f, now it's unclear what what the closure function do with this parameter.
//When using this closureDemo function, user must implement the closure function f
func ClosureDemo(a int, f func(string) string) string {
	return f(strconv.Itoa(a))
}

type functionTemplate func(int)

func fs(n int) {
	fmt.Println("fs", n)
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}
func Trap() {
	var closures [2]func()
	var i int
	for i = 0; i < 2; i++ {
		// tmp := i
		closures[i] = func() {
			fmt.Println(i)
			// fmt.Println(tmp)
		}
	}
	fmt.Printf("  i: %+v\n", i)
	closures[0]()
	closures[1]()
	closures[1] = func() {
		fmt.Printf("  19: \n")
	}
	closures[1]()
}

func BytesBuffer(resp http.Request) {
	//using automatical increase slice
	// body, _ := ioutil.ReadAll(resp.Body)
	// vs
	//using mamnually allocation,but ad least 512kb due to the minRead in bytes.Buffer()
	// buffer := bytes.NewBuffer(make([]byte, 0, resp.ContentLength)
	// buffer.ReadFrom(res.Body)
	// body := buffer.Bytes()

	// body := make([]byte, 0, resp.ContentLength)
	// _, err := io.ReadFull(res.Body, body)
	// vs
	// buffer := bytes.NewBuffer(make([]byte, 0, 65536))
	// io.Copy(buffer, r.Body)
	// temp := buffer.Bytes()
	// length := len(temp)
	// var body []byte
	// //are we wasting more than 10% space?
	// if cap(temp) > (length + length / 10) {
	//   body = make([]byte, length)
	//   copy(body, temp)
	// } else {
	//   body = temp
	// }
}
