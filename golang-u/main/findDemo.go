// Package main provides ...
package main

import "fmt"

func FindDemo(network int) bool {
	println("<<<FindDemo---------------------------")

	nt := []int{1, 3, 4, 5}
	count := len(nt)
	i, j := 0, count-1
	for i < j {
		h := i + (j-i)/2
		if nt[h] < network {
			i = h + 1
		} else {
			j = h
		}
	}
	println("-----------------------------FindDemo>>>")
	return nt[i] == network
}

func main() {

	fmt.Printf(": %+v\n", FindDemo(0))

}
