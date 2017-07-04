// Package pack provides ...
package pack

import "testing"

// func Test GroupedGlobals(t *testing.T) {
// 	GroupedGlobals()
// 	t.Log("done")
// }

func BenchmarkGroupedGlobalsMake(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GroupedGlobalsMake()
	}
}

func BenchmarkGroupedGlobalsFix(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GroupedGlobalsFix()
	}
}
