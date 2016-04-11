// Package main provides ...
package main

import (
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	Data string
	Next *Node
}

func NewNode(d string) *Node {
	return &Node{Data: d, Next: nil}
}

type LinkList struct {
	Header *Node
}

func (ll *LinkList) Sort() {
	if ll.Header == nil {
		return
	}

	oddNode := ll.Header
	curNode := ll.Header.Next
	for curNode != nil && curNode.Next != nil {
		curNext := curNode.Next

		curNode.Next = curNext.Next
		curNext.Next = oddNode.Next
		oddNode.Next = curNext

		oddNode = oddNode.Next
		curNode = curNode.Next
	}

}
func (ll *LinkList) Print() {
	//wrong operation with  address
	// for ll.Header != nil {
	// 	//@toDelete
	// 	fmt.Printf("ll.Header.Data: %+v\n", ll.Header.Data)
	// 	ll.Header = ll.Header.Next
	// }

	tmpNode := new(Node)
	tmpNode = ll.Header
	for tmpNode != nil {
		fmt.Printf("tmpNode.Data: %+v\n", tmpNode.Data)
		tmpNode = tmpNode.Next
	}

}

func (ll *LinkList) Append(n *Node) {

	//wrong operation with  address
	// for ll.Header.Next != nil {
	// 	ll.Header = ll.Header.Next
	// }
	// ll.Header.Next = n

	tmpNode := ll.Header
	for tmpNode.Next != nil {
		tmpNode = tmpNode.Next
	}
	tmpNode.Next = n
}

func (ll *LinkList) Delete(d string) {
	if ll.Header != nil && ll.Header.Data == d {
		ll.Header = ll.Header.Next
		return
	}
	preNode := ll.Header //equals to preNode=new(Node) preNode=ll.Header
	curNode := ll.Header
	for curNode != nil && curNode.Data != d {
		preNode = curNode
		curNode = curNode.Next
	}
	preNode.Next = curNode.Next
}

func main() {
	ll := new(LinkList)
	// ll.Header = NewNode("a")
	// ll.Append(NewNode("b"))
	// ll.Append(NewNode("c"))
	// ll.Delete("c")
	// ll.Print()
	// ll.Delete("b")
	// ll.Print()
	// ll.Delete("a")
	// ll.Print()

	ll.Header = NewNode("1")
	for i := 2; i < 100; i++ {
		ll.Append(NewNode(strconv.Itoa(i)))
	}
	ll.Print()

	os.Stdout.Write(append([]byte("after"), '\n'))

	ll.Sort()
	ll.Print()
}
