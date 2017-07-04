package demo

import "testing"

func TestJustDemo(t *testing.T) {
	JustDemo()
	// structInfoDemo()
	// NewByReflectDemo()
}

func TestRuneCountDemo(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{"a一"}, 2},
		{"normal", args{"a2一"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RuneCountDemo(tt.args.str); got != tt.want {
				t.Errorf("RuneCountDemo() = %v, want %v", got, tt.want)
			}
		})
	}
}
