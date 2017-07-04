package demo

import (
	"math/rand"
	"reflect"
	"testing"
)

func BenchmarkQuickSortDict(b *testing.B) {
	s := make([]string, 0, 20)
	p := make([]byte, 5)
	for i := 0; i < 20; i++ {
		RandStr(p, "abcdefghijklmnopqrstuvwxyz")
		s = append(s, string(p))
	}
	for i := 0; i < b.N; i++ {
		QuickSortDict(s)
	}
}

func BenchmarkInsertionSortDict(b *testing.B) {
	s := make([]string, 0, 20)
	p := make([]byte, 5)
	for i := 0; i < 20; i++ {
		RandStr(p, "abcdefghijklmnopqrstuvwxyz")
		s = append(s, string(p))
	}
	for i := 0; i < b.N; i++ {
		InsertionSortDict(s)
	}
}

func TestQuickSortDict(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"normal", args{[]string{"b", "c", "a"}}, []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSortDict(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSortDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertionSortDict(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"normal", args{[]string{"b", "c", "a"}}, []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSortDict(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSortDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertionSortInt(t *testing.T) {
	type args struct {
		ints []int
	}
	ns := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		ns = append(ns, rand.Intn(1000))
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {"normal", args{[]int{2, 5, 4, 3, 1}}, []int{1, 2, 3, 4, 5}},
		{"normal", args{ns}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSortInt(tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("InsertionSortInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSortInt(t *testing.T) {
	type args struct {
		ints []int
	}

	ns := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		ns = append(ns, rand.Intn(1000))
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {"normal", args{[]int{2, 5, 4, 3, 1}}, []int{1, 2, 3, 4, 5}},
		{"normal", args{ns}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSortInt(tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("QuickSortInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	ns := GetIntslice()
	newNs := make([]int, len(ns))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// BubbleSort(ns)
		copy(newNs, ns)
		openSourceBubblesort(ns)
	}
}

func BenchmarkQuickSortInt(b *testing.B) {
	ns := GetIntslice()
	newNs := make([]int, len(ns))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// QuickSortInt(ns)
		copy(newNs, ns)
		openSourcequicksort(newNs)
	}
}

func BenchmarkInsertionSortInt(b *testing.B) {
	ns := GetIntslice()
	newNs := make([]int, len(ns))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// InsertionSortInt(ns)
		copy(newNs, ns)
		openSourceInsertionsort(ns)
	}
}

func TestBubbleSort(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"normal", args{[]int{2, 5, 4, 3, 1}}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBenchSort(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"normal"},
	}
	for _ = range tests {
		BenchSort()
	}
}
