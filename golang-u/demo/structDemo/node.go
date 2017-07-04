// Package  demo
package demo

//playground:`https://play.golang.org/p/Bj4al1ARXo`
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

	//jus thinking theory of relativity: 1x stands still from view of 2x,2x move aroung the cycle with speed 1x, so they will meet finally
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
	//just thinking: when 2x move from meetpoint(just think one node in the cycle) through cyclepoint to meetpoint,
	// this distance is equal to the distance from starting to the meetpoint,
	// and 1x,2x they all move throungh the same path which from cyclepoint to meetpoint,
	// so we can say:[distance from startpoint to cyclepoint] is equal to [distance from current location of speed1x(meetpoint) to cyclepoint]
	//question: which speed of 2x works for this finnal step, but 3x not working, because 3x works same as 1x actually for 1x in cycle
	log.Printf("ret speed1x.Val: %+v, 2x: %+v\n", speed1x.Val, speed2x.Val)
	log.Printf("after %v moves they meet\n", c)
	speed2x = head //no cyclepoint, but the startpoint
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
