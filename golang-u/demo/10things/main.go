package main

import (
	"demo/10things/pack"
	"fmt"
	"time"
)

func main() {
	fmt.Println("dd")
	pack.GroupedGlobalsFix()
	// pack.GroupedGlobals()
	// pack.EmbeddedLock()
	// pack.MethodExpressions()
	//
	//------------------send and receive on the same channel start----------------
	// done := make(chan struct{})
	// langs := []string{"Go", "C", "C++", "Java", "Perl", "Python"}
	// for _, l := range langs {
	// 	go pack.Warrior(l, done)
	// }
	// for _ = range langs {
	// 	<-done
	// }
	//------------------send and receive on the same channel end------------------

	//------------------using close to broadcast start----------------
	block, done := make(chan struct{}), make(chan struct{})
	for i := 0; i < 4; i++ {
		go pack.Waiter(i, block, done)
	}
	time.Sleep(5 * time.Second)
	close(block)
	for i := 0; i < 4; i++ {
		<-done
	}
	//------------------using close to broadcast end------------------
}
