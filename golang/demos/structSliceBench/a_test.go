// Package main provides ...
package newDir

import "testing"

func TestJustDemo(t *testing.T) {
	JustDemo()
}

// func BenchmarkJustDemo(b *testing.B) {
// 	a := A{Name: "xj"}
// 	t := make([]A, 0, 1)
// 	for i := 0; i < b.N; i++ {
// 		println(i, " round")
// 		t = append(t, a)
// 		fmt.Printf("cap(t): %+v\n", cap(t))
// 		fmt.Printf("len(t): %+v\n", len(t))
// 	}
// }
