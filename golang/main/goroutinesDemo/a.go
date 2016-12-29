// Package main provides ...
package main

import (
	"fmt"
	"sync"
)

//RInt implements ...
var gWrite int

func RInt(a []*A, tc *A) {
	for i := 0; i < len(a); i++ {
		print(a[i].E)
	}
	tc = NewA()
	//@TODO write lock
	// gWrite = 2
	println()
}

type A struct {
	E int `json:"E "`
}

//NewA implements ...
func NewA() *A {
	return &A{8888}
}

func main() {
	len := 100
	a := make([]*A, 0, 100)
	for i := 0; i < len; i++ {
		a = append(a, &A{i})
	}
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		RInt(a[:30], NewA())
	}()
	go func() {
		defer wg.Done()
		RInt(a[30:60], NewA())
	}()
	go func() {
		defer wg.Done()
		RInt(a[60:], NewA())
	}()
	wg.Wait()
	//@toDelete
	fmt.Printf("good: works\n")
	return
}
