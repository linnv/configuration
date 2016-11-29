package demo

func echoList(str string, node *ListNode) {
	print(str, "`")
	var tmp *ListNode
	tmp = node
	for tmp != nil {
		print(tmp.Val, " ")
		tmp = tmp.Next
	}
	println("`")
}

func GenerateList(numbers []int) *ListNode {
	if len(numbers) < 1 {
		return nil
	}
	ret := &ListNode{Val: numbers[0], Next: nil}
	var tmp *ListNode
	tmp = ret
	for i := 1; i < len(numbers); i++ {
		tmp.Next = &ListNode{Val: numbers[i], Next: nil}
		tmp = tmp.Next
	}
	return ret
}

func addTwoNumbers(one, two *ListNode) *ListNode {
	if two == nil {
		return one
	}

	if one == nil {
		one = &ListNode{0, nil}
	}

	var carry, lastNodeVal, tmpSum int
	var tmpOne, tmpTwo, ret *ListNode
	tmpOne, tmpTwo = one, two
	ret = tmpTwo

	for tmpOne != nil || tmpTwo != nil || carry > 0 {
		if tmpOne != nil {
			tmpSum += tmpOne.Val
			tmpOne = tmpOne.Next
		}
		tmpSum += tmpTwo.Val
		tmpSum += carry
		lastNodeVal = tmpSum % 10
		carry = tmpSum / 10

		tmpSum = 0
		if tmpTwo.Next == nil && (tmpOne != nil || carry > 0) {
			tmpTwo.Next = &ListNode{}
		}
		tmpTwo.Val = lastNodeVal
		tmpTwo = tmpTwo.Next
	}

	return ret
}
