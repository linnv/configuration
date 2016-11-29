package demo

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head *ListNode
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"normal", args{GenerateList([]int{3, 2, 1}), 1, 3}, GenerateList([]int{1, 2, 3})},
		{"normal", args{GenerateList([]int{4, 3, 2, 1}), 2, 3}, GenerateList([]int{4, 2, 3, 1})},
		{"edge", args{GenerateList([]int{2, 1}), 1, 2}, GenerateList([]int{1, 2})},
		{"edge", args{GenerateList([]int{2, 1}), 1, 1}, GenerateList([]int{2, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
