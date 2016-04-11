package ben

import (
	"fmt"
	"testing"
	"time"
)

func Benchmarks(b *testing.B) {
	// customTimerTag := false
	// if customTimerTag {
	// 	b.StopTimer()
	// }
	// b.SetBytes(1024)
	// time.Sleep(time.Second)
	// if customTimerTag {
	// 	b.StartTimer()
	// }
}

// func BenchmarkDemoXXXFuncs(b *testing.B) {
// 	//@TODO invoke this benchmark func by executing  go test -bench=Demo     xxx
// }

func ExampleDemo() {
	fmt.Printf("example test\n")
	//Output: just example test
}

// func main() {
// 	fmt.Println("benchmark")
// 	return
// }
