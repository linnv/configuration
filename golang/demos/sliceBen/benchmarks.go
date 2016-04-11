// Package main provides ...
package ben

import (
	"fmt"
	"sort"
)

func ExampleDemo() {
	fmt.Printf("example test\n")
}

func SortDeleteDuplicate(a []int) []int {
	aLen := len(a)
	index := 0
	sort.IntSlice(a).Sort()
	for i := 0; i < aLen; i++ {
		if a[index] != a[i] {
			index++
			a[index] = a[i]
		}
	}

	if aLen != index {
		return a[:index+1]
	}
	return a
}

func UniqueIntArray(a []int) []int {
	al := len(a)
	if al == 0 {
		return a
	}

	ret := make([]int, al)
	index := 0

loopa:
	for i := 0; i < al; i++ {
		for j := 0; j < index; j++ {
			if a[i] == ret[j] {
				continue loopa
			}
		}
		ret[index] = a[i]
		index++
	}

	return ret[:index]
}
