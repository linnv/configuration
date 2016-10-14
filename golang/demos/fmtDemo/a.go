package newDir

import (
	"fmt"
	"os"
	"time"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	a := []int{1, 2, 3}
	//@toDelete
	fmt.Printf("  a: %d\n", a)
	fmt.Printf(": %g\n", 1.222)
	fmt.Printf(": %G\n", 1.222)
	fmt.Printf(": %e\n", 1.222)
	fmt.Printf(": %E\n", 1.222)
	fmt.Printf(": %f\n", 1.222)
	fmt.Printf(": %.1f\n", 1.222)
	fmt.Printf(": %b\n", 4)
	var s interface{}
	s = fmt.Sprintf("%.2f", 1.33333)
	fmt.Printf("s: %+v\n", s)
	time.Sleep(time.Second * 10)
	switch s.(type) {
	case string:
		os.Stdout.Write(append([]byte("string"), '\n'))
	case float32, float64:
		os.Stdout.Write(append([]byte("float"), '\n'))
	default:
		os.Stdout.Write(append([]byte("unkown"), '\n'))
	}
	//don't have to use strconv.Itoa()
	r1 := fmt.Sprintln("xxx", 333, 333.444, " str at end")
	r2 := fmt.Sprint("xxx ", 333)
	fmt.Printf("r1: %+v\n", r1)
	fmt.Printf("r2: %+v\n", r2)
	return
}
