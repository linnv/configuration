// Package main provides ...
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// var i int
	// fmt.Scan(&i)
	var i *int = new(int)
	fmt.Scan(i)
	fmt.Println("read number", *i, "from stdin")

	// wg.Add(1) //count of goroutines
	// c := make(chan string)
	// go func() {
	// 	defer wg.Done()
	// 	var o string
	// 	fmt.Scanln(&o)
	// 	c <- o
	// }()
	// str := <-c
	// fmt.Printf("str from stdin: %+v\n", str)
	// wg.Wait()

	// bio := bufio.NewReader(os.Stdin)
	// line, hasMoreInLine, err := bio.ReadLine()
	// if err != nil {
	// 	return
	// }
	//
	// fmt.Printf("str from stdin: %+v\n", string(line))
	// if hasMoreInLine {
	// 	fmt.Printf("more lines:\n")
	// }
}
