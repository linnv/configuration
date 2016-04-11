// Package main provides ...
package main

import "fmt"

type A struct {
	N int
	// N string
}

func main() {
	arr := make([]*A, 0, 1)
	for i := 0; i < 10; i++ {
		arr = append(arr, new(A))
		arr[i].N = i
		// arr = append(arr, &A{N: i})
	}

	brr := make([]*A, 0, 1)
	for i := 0; i < 5; i++ {
		brr = append(brr, arr[i])
	}
	for k, v := range brr {
		fmt.Printf("%+v: %+v\n", k, v)
	}
	fmt.Printf(": works\n")
	for k, v := range arr {
		fmt.Printf("%+v: %+v\n", k, v)
	}

}
