// Package main provides ...
package main

import (
	"math/rand"

	"github.com/linnv/logx"
)

//key point: recursive iterator
// func partion(slice []int, left, right int) (indexLastPivot int) {
// 	indexLastPivot = left
// 	pivot := slice[right]
//
// 	for i := left; i < right; i++ {
// 		if slice[i] < pivot {
// 			slice[indexLastPivot], slice[i] = slice[i], slice[indexLastPivot]
// 			indexLastPivot++
// 		}
// 	}
// 	slice[indexLastPivot], slice[right] = slice[right], slice[indexLastPivot]
// 	// now left of slice is less than pivot ,right is greater ,but both are unsort
// 	return indexLastPivot
// }
//
// func sort(slice []int, left, right int) {
// 	if left > right {
// 		return
// 	}
//
// 	indexPivot := partion(slice, left, right)
// 	logx.Debug("indexPivot: %+v\n", indexPivot)
// 	logx.Debug("end slice: %+v\n", slice)
// 	logx.Debug("sorting the left part\n")
// 	sort(slice, left, indexPivot-1)
// 	logx.Debug("sorting the right part\n")
// 	sort(slice, indexPivot+1, right)
// }

// func partion(s []int, left, right int) (indexPivot int) {
// 	pivotValue := s[right]
// 	indexPivot = left
//
// 	// don't start from inde 0, because sort the right part isn't from index 0
// 	for i := left; i < right; i++ {
// 		if s[i] < pivotValue {
// 			s[i], s[indexPivot] = s[indexPivot], s[i]
// 			indexPivot++
// 		}
// 	}
// 	s[right], s[indexPivot] = s[indexPivot], s[right]
// 	return
// }
//
// func sort(s []int, left, right int) {
// 	if left > right {
// 		return
// 	}
//
// 	indexPivot := partion(s, left, right)
// 	sort(s, left, indexPivot-1)
// 	sort(s, indexPivot+1, right)
// }

var (
	allchars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func RandStr(p []byte, limit string) {
	pLen := len(p)
	limitLen := len(limit)
	for i := 0; i < pLen; i++ {
		p[i] = limit[rand.Intn(limitLen)]
	}
	return
}

func main() {
	// s := []string{"bba", "baa", "abc", "ccc"}
	s := make([]string, 0, 20)
	p := make([]byte, 5)
	for i := 0; i < 20; i++ {
		RandStr(p, "abcdefghijklmnopqrstuvwxyz")
		s = append(s, string(p))
	}
	logx.Debug("s: %+v\n", s)
	// QuickSortStr(s)
	DictSortInsertion(s)
	logx.Debug("s end: %+v\n", s)
	// rand.Read(p)
	logx.Debug("string(p): %+v\n", string(p))
	// ints := rand.Perm(20)
	// logx.Debug("ints: %+v\n", ints)
	// // sort(ints, 0, len(ints)-1)
	// quickSort(ints, 0, len(ints)-1)
	// logx.Debug("end ints: %+v\n", ints)
}

func partion(s []int, left, right int) (indexPivot int) {
	indexPivot = left
	valuePivot := s[right]
	for i := left; i < right; i++ {
		if s[i] < valuePivot {
			s[i], s[indexPivot] = s[indexPivot], s[i]
			indexPivot++
		}
	}
	s[right], s[indexPivot] = s[indexPivot], s[right]
	return indexPivot
}

func quickSort(s []int, left, right int) {
	if left > right {
		return
	}
	indexPivot := partion(s, left, right)
	logx.Debug("indexPivot: %+v\n", indexPivot)
	logx.Debug("end slice: %+v\n", s)
	logx.Debug("sorting the left part\n")
	quickSort(s, left, indexPivot-1)
	quickSort(s, indexPivot+1, right)
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

func QuickSortStr(str []string) {
	sortStr(str, 0, len(str)-1)
}

func sortStr(str []string, left, right int) {
	if left > right {
		return
	}
	pivot := partionStr(str, left, right)
	sortStr(str, left, pivot-1)
	sortStr(str, pivot+1, right)
}

func DictSortInsertion(str []string) {
	strLen := len(str)
	if strLen < 1 {
		return
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
	return
}
