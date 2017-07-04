// Package main provides ...
package demo

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	allchars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var defaulInts []int

func init() {
	// const count = 50000
	const count = 10
	defaulInts = make([]int, 0, count)
	for i := 0; i < count; i++ {
		defaulInts = append(defaulInts, rand.Intn(count))
	}
}

func GetIntslice() []int {
	ns := make([]int, len(defaulInts))
	copy(ns, defaulInts)
	return ns
}

func RandStr(p []byte, limit string) {
	pLen := len(p)
	limitLen := len(limit)
	for i := 0; i < pLen; i++ {
		p[i] = limit[rand.Intn(limitLen)]
	}
	return
}

func swapStr(s, p string) bool {
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		if int(s[i]) == int(p[i]) {
			continue
		}
		if int(s[i]) < int(p[i]) {
			return true
		}
		return false
	}
	return false
}

func partionStr(str []string, left, right int) int {
	value := str[right]
	indexPivot := left
	for i := left; i < right; i++ {
		if swapStr(str[i], value) {
			str[i], str[indexPivot] = str[indexPivot], str[i]
			indexPivot++
		}
	}
	str[right], str[indexPivot] = str[indexPivot], str[right]
	return indexPivot
}

func QuickSortDict(str []string) []string {
	return sortStr(str, 0, len(str)-1)
}

func sortStr(str []string, left, right int) []string {
	if left > right {
		return str
	}
	pivot := partionStr(str, left, right)
	sortStr(str, left, pivot-1)
	sortStr(str, pivot+1, right)
	return str
}

func InsertionSortDict(str []string) []string {
	strLen := len(str)
	if strLen < 1 {
		return str
	}

	for i := 1; i < strLen; i++ {
		sortingStr := str[i]
		sortingStrLen := len(sortingStr)
		j := i - 1
		for j >= 0 {
			for n := 0; n < sortingStrLen; n++ {
				if int(str[j][n]) == int(sortingStr[n]) {
					continue
				}
				if int(str[j][n]) > int(sortingStr[n]) {
					str[j+1] = str[j]
					j--
					break
				}
				goto endLoop
			}
		}
	endLoop:
		str[j+1] = sortingStr
	}
	return str
}

func partion(ints []int, left, right int) int {
	value := ints[right]
	indexPivot := left
	for i := left; i < right; i++ {
		if ints[i] < value {
			ints[i], ints[indexPivot] = ints[indexPivot], ints[i]
			indexPivot++
		}
	}
	ints[right], ints[indexPivot] = ints[indexPivot], ints[right]
	return indexPivot
}

func quickSort(ints []int, left, right int) []int {
	if left > right {
		return ints
	}
	pivot := partion(ints, left, right)
	quickSort(ints, left, pivot-1)
	quickSort(ints, pivot+1, right)
	return ints
}

func QuickSortInt(ints []int) []int {
	// start := time.Now()
	ns := quickSort(ints, 0, len(ints)-1)
	// fmt.Printf("quicksortDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	return ns
}

func InsertionSortInt(ints []int) []int {
	// start := time.Now()
	intsLen := len(ints)
	if intsLen < 1 {
		return ints
	}
	for i := 1; i < intsLen; i++ {
		j := i - 1
		value := ints[i]
		for j >= 0 && value < ints[j] {
			ints[j+1] = ints[j]
			j--
		}
		ints[j+1] = value
	}
	// fmt.Printf("insertionDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	return ints
}

func BubbleSort(ints []int) []int {
	intsLen := len(ints)
	for i := intsLen - 1; i >= 0; i-- {
		sorted := true
		for j := 1; j <= i; j++ {
			if ints[j] < ints[j-1] {
				ints[j], ints[j-1] = ints[j-1], ints[j]
				sorted = false

			}
		}
		if sorted {
			return ints
		}
	}
	return ints
}

func BenchSort() {
	start := time.Now()
	const loopAmount = 100
	insertionInts := make([]int, len(defaulInts))
	for i := 0; i < loopAmount; i++ {
		copy(insertionInts, defaulInts)
		openSourceBubblesort(insertionInts)
		_ = insertionInts
	}
	fmt.Printf("insertionDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))

	start = time.Now()
	bubbleInts := make([]int, len(defaulInts))
	for i := 0; i < loopAmount; i++ {
		copy(bubbleInts, defaulInts)
		openSourceInsertionsort(bubbleInts)
		_ = bubbleInts
	}

	fmt.Printf("bubbleDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))

	start = time.Now()
	quickInts := make([]int, len(defaulInts))
	for i := 0; i < loopAmount; i++ {
		copy(quickInts, defaulInts)
		openSourcequicksort(quickInts)
		_ = quickInts
	}
	fmt.Printf("quicksortDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
}
