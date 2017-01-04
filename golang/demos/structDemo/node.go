// Package  demo
package newDir

//playground:`https://play.golang.org/p/7EF4qAIC3u`
import (
	"fmt"
	"log"
	"time"
)

type Node struct {
	Mark bool
	Val  int
	Next *Node
}

func NewNode(v int) *Node {
	return &Node{Val: v}
}

func NewNodeWithNext(v int, next *Node) *Node {
	return &Node{Val: v, Next: next}
}

func TravelDemo() {
	println("//<<-------------------------TravelDemo start-----------")
	start := time.Now()
	head := NewNode(0)
	cyclePoint := NewNode(1)
	head.Next = cyclePoint
	cyclePoint.Next = NewNodeWithNext(2, NewNodeWithNext(3, NewNodeWithNext(4, cyclePoint)))
	// cyclePoint.Next = NewNodeWithNext(2, NewNodeWithNext(3, head))
	var speed1x, speed2x *Node
	speed1x = head.Next
	speed2x = head.Next.Next
	c := 0
	for speed1x != speed2x {
		c++
		log.Printf("speed1x.Val: %+v, 2x: %+v\n", speed1x.Val, speed2x.Val)
		speed1x = speed1x.Next
		speed2x = speed2x.Next.Next
		if speed2x.Val == cyclePoint.Val {
			log.Println("cycle: works")
		}
	}
	//now [distance from start point to cycle point] is equal to [distance from current location of speed1x to cycle point]
	log.Printf("ret speed1x.Val: %+v, 2x: %+v\n", speed1x.Val, speed2x.Val)
	log.Printf("after %v moves they meet\n", c)
	speed2x = head //no cycle point, but the start point
	c = 0
	for speed1x != speed2x {
		c++
		speed1x = speed1x.Next
		speed2x = speed2x.Next
	}
	log.Printf("ret speed1x.Val: %+v, 2x: %+v\n", speed1x.Val, speed2x.Val)
	log.Printf("distance from start point to cycle point: %+v\n", c)
	fmt.Printf("TravelDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------TravelDemo end----------->>")
}
