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

type LRULink struct {
	NodeCount int
	Head      *Node
}

func NewNode(data string) *Node {
	return &Node{Data: data, Next: nil}
}

func NewLRULink() *LRULink {
	return &LRULink{NodeCount: 1,
		Head: NewNode("head"),
	}
}

func (lrul *LRULink) PrintAll() {
	node := lrul.Head
	for node != nil {
		println(node.Data)
		node = node.Next
	}
	return
}

func (lrul *LRULink) Upsert(data string) *Node {
	if lrul.Head == nil {
		return nil
	}
	var pre *Node
	target := lrul.Head
	for target.Next != nil {
		if target.Data == data {
			if pre != nil {
				//set target data to first one
				pre.Next = target.Next
				target.Next = lrul.Head
				lrul.Head = target
			}
			return target
		}
		pre = target
		target = target.Next
	}

	if lrul.NodeCount == 100 {
		// delete the last node
		pre.Next = target.Next
	}

	//node locate at first
	target = NewNode(data)
	target.Next = lrul.Head
	lrul.Head = target
	if lrul.NodeCount != 100 {
		lrul.NodeCount++
	}
	return target.Next
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
