package demo

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m > n {
		return head
	}
	var pre, suf, handlerStart, handlerEnd *ListNode
	pre = head
	i := 2
	if m < 3 { //edge cases
		switch m {
		case 1:
			handlerStart = head
		case 2:
			handlerStart = pre.Next

		}
		i = m
	} else {
		for i < m {
			pre = pre.Next
			i++
		}
		handlerStart = pre.Next
	}

	handlerEnd = handlerStart
	for i < n {
		handlerEnd = handlerEnd.Next
		i++
	}
	suf = handlerEnd.Next
	handlerEnd.Next = nil
	reverseListFunc := func(head *ListNode) *ListNode {
		var pre, cur, backup *ListNode
		cur = head
		for cur != nil {
			backup = cur.Next
			cur.Next = pre
			pre = cur
			cur = backup
		}
		return pre
	}
	reverseList := reverseListFunc(handlerStart)
	if pre.Next != nil {
		pre.Next = reverseList
	}
	handlerStart.Next = suf
	if m > 1 {
		return head
	}

	return handlerEnd
}
