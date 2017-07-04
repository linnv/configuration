package foo

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 20 {
		t.Error("test foo:Addr failed")
	} else {
		t.Log("test foo:Addr pass")
	}
}

func BenchmarkAdd(b *testing.B) {
	// 如果需要初始化，比较耗时的操作可以这样：
	// b.StopTimer()
	// .... 一堆操作
	// b.StartTimer()
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

func benchmarkFib(a, r int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Add(a, r)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, 12, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, 23, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, 33, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, 10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, 20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, 40, b) }
