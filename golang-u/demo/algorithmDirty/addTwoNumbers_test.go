package demo

import (
	"reflect"
	"testing"
)

// func TestGenerateList(t *testing.T) {
// 	one := GenerateList([]int{1, 2, 3, 3, 5, 5, 6})
// 	two := GenerateList([]int{1, 9})
// 	echoList("input1", one)
// 	echoList("input2", two)
// 	ret := addTwoNumbers(one, two)
// 	echoList("output", ret)
// }

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		one *ListNode
		two *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"normal", args{GenerateList([]int{1, 2, 3, 3, 5, 5, 6}), GenerateList([]int{1, 9})}, GenerateList([]int{2, 1, 4, 3, 5, 5, 6})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.one, tt.args.two); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
