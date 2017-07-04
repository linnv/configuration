package pack

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type A struct {
	N int
}

//go 1.4.2
const lenFix = 100000000 //fix occurs error but make runs normally
// const lenFix = 1000000 //fix is slow than make   nearly haf
// const lenFix = 100000 //equal
// const lenFix = 10000 //fix is better than  make

func GroupedGlobalsMake() []int {
	b := make([]int, lenFix)
	for i := 0; i < lenFix; i++ {
		b[i] = i
	}
	return b
}

func GroupedGlobalsFix() [lenFix]int {
	b := [lenFix]int{}
	for i := 0; i < lenFix; i++ {
		b[i] = i
	}
	return b
}

func GroupedGlobals() {
	var config struct {
		APIKey string
		Count  int
	}

	config.APIKey = "BADC0C0A"
	fmt.Printf("config: %+v\n", config)
}

func EmbeddedLock() {
	var hits struct {
		sync.Mutex
		n int
	}

	hits.Lock()
	hits.n++
	hits.Unlock()
	fmt.Printf("hits: %+v\n", hits)
}

//------------------method expressions start----------------
type T struct{}

func (T) Foo(s string)    { println("it works", s) }
func (T) NoneParameters() { println("it works x") }

var fn func(T, string) = T.Foo
var nfn func(T) = T.NoneParameters

func MethodExpressions() {
	a := T{}
	fn(a, "ssss")
	nfn(a)
}

//------------------method expressions end------------------

//------------------send and receive on the same channel start----------------

var battle = make(chan string)

func Warrior(name string, done chan struct{}) {
	select {
	case opponent := <-battle:
		fmt.Printf("%s beat %s\n", name, opponent)
	case battle <- name:
		// I lost :-(
	}
	done <- struct{}{}
}

//------------------send and receive on the same channel end------------------

//------------------using close to broadcast start----------------

func Waiter(i int, block, done chan struct{}) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	fmt.Println(i, "waiting...")
	<-block
	fmt.Println(i, "done!")
	done <- struct{}{}
}

//------------------

//------------------using close to broadcast end------------------
