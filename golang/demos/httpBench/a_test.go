// Package main provides ...
package newDir

import "testing"

func TestJustDemo(t *testing.T) {
	JustDemo()
}

func BenchmarkJustDemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		println(i, " round")
	}
}
