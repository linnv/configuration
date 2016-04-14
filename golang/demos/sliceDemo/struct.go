package newDir

import "fmt"

const (
	FIRST = iota
)

const (
	A = 1 << iota
	B
	C
)

var (
	AB  = A | B
	ABC = A | B | C
	AC  = A | C
	BC  = B | C
)

type AA struct {
	Name string `json:"Name"`
}

type Node struct {
	Data string `json:"Data"`
	Next *Node
}

func (n *Node) PrintAll() {
	if n == nil {
		return
	}
	tmp := n
	for tmp != nil {
		//@toDelete
		fmt.Printf("n.Data: %+v %x\n", tmp.Data, tmp)
		tmp = tmp.Next
	}
}

type Int64Slice []int64

var timesConsumption Int64Slice

func (in Int64Slice) Less(i, j int) bool {
	return in[i] > in[j]
}

func (in Int64Slice) Len() int {
	return len(in)
}

func (in Int64Slice) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}
