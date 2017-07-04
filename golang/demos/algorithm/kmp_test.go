package demo

import (
	"reflect"
	"testing"
)

func TestGetPattern(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name        string
		args        args
		wantPattern []int
	}{
	// {"normal", args{"abababca"}, []int{1, 11, 11, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPattern := GetPattern(tt.args.str); !reflect.DeepEqual(gotPattern, tt.wantPattern) {
				t.Errorf("GetPattern([%s]) = %v, want %v", tt.args.str, gotPattern, tt.wantPattern)
			}
		})
	}
}

func TestKMP(t *testing.T) {
	type args struct {
		dst  string
		sour string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// {"normal", args{"abababca", "abababca"}, 2},
	// {"normal", args{"abababca", "bacbababaabcbababababca"}, 15},
	// {"normal", args{"abababca", "bbacbababaabcbababbacbababaabcbababbacbababaabcbababbacbababaabcbababbacbababaabcbababbacbababaabcbababbacbababaabcbababacbababaabcbababababca"}, 15},
	// {"normal", args{"abababca", "bacbababaabcbab,abababca"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := KMP(tt.args.dst, tt.args.sour); got != tt.want {
			if got := KMPDemo(tt.args.dst, tt.args.sour); got != tt.want {
				t.Errorf("KMP() = %v, want %v", got, tt.want)
			}
			if got := LibIndexDemo(tt.args.dst, tt.args.sour); got != tt.want {
				t.Errorf("KMP() = %v, want %v", got, tt.want)
			}
		})
	}
}
