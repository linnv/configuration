// Package main provides ...
package newDir

import (
	"fmt"
	"testing"
)

func TestJustDemo(t *testing.T) {
	JustDemo()

	println("<<<mmap Demo start---------------------------")

	done := make(chan bool)
	b := NewCard(NewSimpleAccount(1))
	b.Deposit(10, "jialin")
	b.Deposit(30, "jialin wu")
	fmt.Printf("monney left init map: %+v\n", b.Balance())

	go func() {
		b.Draw(10, "jialin wu")
		done <- true
	}()

	go func() {
		b.Draw(12, "jialin")
		done <- true
	}()
	<-done
	<-done
	fmt.Printf("monney left done map: %+v\n", b.Balance())
	println("-----------------------------mmap Demo end>>>")

	println("<<<chan Demo start---------------------------")

	donec := make(chan bool)
	bc := NewCardChan(NewSimpleAccount(1))
	bc.Deposit(10, "jialin")
	bc.Deposit(30, "jialin wu")
	fmt.Printf("monney left initial channel: %+v\n", bc.Balance())

	go func() {
		bc.Draw(10, "jialin wu")
		donec <- true
	}()

	go func() {
		bc.Draw(12, "jialin")
		donec <- true
	}()
	<-donec
	<-donec
	fmt.Printf("monney left done channel: %+v\n", bc.Balance())
	println("-----------------------------chan Demo end>>>")
}
