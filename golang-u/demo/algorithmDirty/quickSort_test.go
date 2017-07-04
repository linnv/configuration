package main

import "testing"

func TestQuickSortStr(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
	}{
		{"normal", args{[]string{"bba", "baa", "abc", "ccc"}, []string{"abc", "baa", "bba", "ccc"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSortStr(tt.args.str)
		})
	}
}
