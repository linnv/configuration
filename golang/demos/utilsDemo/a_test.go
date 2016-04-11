// Package main provides ...
package newDir

import (
	"fmt"
	"testing"
)

func TestJustDemo(t *testing.T) {
	fmt.Printf("Today(): %+v\n", Today())
	// SortStrSlice([]string{"3", "4"})
	strl := []string{"20151010", "20141013", "20151011"}
	// strl := []string{"", "b1", "b5"}
	r := SortStrSlice(strl)
	fmt.Printf("  r: %+v\n", r)
}

func BenchmarkJustDemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		println(i, " round")
	}
}
