// Package main provides ...
package demo

import (
	"fmt"
	"testing"
)

// func TestJustDemo(t *testing.T) {
// 	// JustDemo()
// 	p := NewPerson(NewBody(100), "jialin")
// 	fmt.Printf("p.DoGetEnergy(): %+v\n", p.DoGetEnergy())
// 	p.DoFillEnergy(100)
// 	fmt.Printf("p.DoGetEnergy(): %+v\n", p.DoGetEnergy())
// 	err := p.DoConsumeEnergy(50)
// 	if err != nil {
// 		fmt.Printf("err.Error(): %+v\n", err.Error())
// 		return
// 	}
// 	fmt.Printf("p.DoGetEnergy(): %+v\n", p.DoGetEnergy())
// 	p.AllMember()
//
// }

func BenchmarkJustDemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := NewPerson(NewBody(100), "jialin")
		fmt.Printf("p.DoGetEnergy(): %+v\n", p.DoGetEnergy())
		p.DoFillEnergy(100)
		fmt.Printf("p.DoGetEnergy(): %+v\n", p.DoGetEnergy())
		err := p.DoConsumeEnergy(50)
		if err != nil {
			fmt.Printf("err.Error(): %+v\n", err.Error())
			return
		}
		fmt.Printf("p.DoGetEnergy(): %+v\n", p.DoGetEnergy())
		p.AllMember()
	}
}
