package demo

import "testing"

func TestJustDemo(t *testing.T) {
	JustDemo()
}

func TestFlagDemo(t *testing.T) {
	type args struct {
		name string
		ss   []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"none", args{name: "p0", ss: []string{"p0=1"}}, -1}, //only - and -- works
		{"-", args{name: "p1", ss: []string{"-p1=11"}}, 11},
		// {"--", args{name: "p2", ss: []string{"--p2=1"}}, 1},
		{"--", args{"p2", []string{"--p2=12"}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagDemo(tt.args.name, tt.args.ss); got != tt.want {
				t.Errorf("FlagDemo() = %v, want %v", got, tt.want)
			}
		})
	}
}
