package demo

import "testing"

func TestGenerateList(t *testing.T) {
	one := GenerateList([]int{1, 2, 3, 3, 5, 5, 6})
	two := GenerateList([]int{1, 9})
	ret := addTwoNumbers(one, two)
	echoList(ret)
}
