package demo

type twoQueueOneStack struct {
	queueX []int
	queueY []int
}

func (qs *twoQueueOneStack) Pop() (r int) {
	if qs.IsEmpty() {
		return -1
	}
	if qs.lengthX() != 0 {
		queueXLen := qs.lengthX()
		qs.queueY = make([]int, queueXLen-1)
		copy(qs.queueY, qs.queueX[:queueXLen-1])
		// qs.queueY = append(qs.queueY, qs.queueX[:queueXLen-1]...)
		r = qs.queueX[queueXLen-1]
		qs.queueX = qs.queueX[queueXLen:]

	} else {
		queueYLen := qs.lengthY()
		qs.queueX = make([]int, queueYLen-1)
		copy(qs.queueX, qs.queueY[:queueYLen-1])
		// qs.queueX = append(qs.queueX, qs.queueY[:qs.lengthY()-1]...)
		r = qs.queueY[queueYLen-1]
		qs.queueY = qs.queueY[queueYLen:]
	}
	return
}

func (qs *twoQueueOneStack) Push(n int) {
	if qs.lengthX() != 0 {
		qs.queueX = append(qs.queueX, n)
	} else {
		qs.queueY = append(qs.queueY, n)
	}
}

func (qs *twoQueueOneStack) IsEmpty() bool {
	return qs.lengthX() == 0 && qs.lengthY() == 0
}

func (qs *twoQueueOneStack) lengthX() int {
	return len(qs.queueX)
}

func (qs *twoQueueOneStack) lengthY() int {
	return len(qs.queueY)
}

func (qs *twoQueueOneStack) Size() int {
	return qs.lengthX() + qs.lengthY()
}
