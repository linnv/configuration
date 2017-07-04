package ben

import "testing"

func BenchmarkExampleDemo(b *testing.B) {
	//@TODO invoke this benchmark func by executing  go test -bench=Demo     xxx

	for i := 0; i < b.N; i++ {
		ExampleDemo()
	}
}
