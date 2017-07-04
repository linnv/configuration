package demo

import "testing"

func TestLRUDemo(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// {"normal", args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LRUDemo(tt.args.i); got != tt.want {
				t.Errorf("LRUDemo() = %v, want %v", got, tt.want)
			}
		})
	}
}
