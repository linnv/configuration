// Package main provides ...
package newDir

import (
	"math/rand"
	"sync"
	"time"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

type Accounter interface {
	Draw(uint)
	Deposit(uint)
	Balance() int
}

type SimpleAccount struct {
	balance int `json:"Balance"`
}

func NewSimpleAccount(n int) *SimpleAccount {
	return &SimpleAccount{balance: n}
}

func (s *SimpleAccount) setBalance(n int) {
	s.Latency()
	s.balance = n
}

func (s *SimpleAccount) Draw(n uint) {
	if s.balance >= int(n) {
		s.setBalance(s.balance - int(n))
		return
	}
	panic("jack is poor")
}

func (s *SimpleAccount) Deposit(n uint) {
	s.setBalance(s.balance + int(n))
}

func (s *SimpleAccount) Balance() int {
	return s.balance
}

func (s *SimpleAccount) Latency() {
	<-time.After(time.Duration(rand.Intn(100)) * time.Millisecond)
}

type Card struct {
	lock    sync.Mutex
	account Accounter
}

func NewCard(ac Accounter) *Card {
	return &Card{account: ac}
}
func (c *Card) Draw(n uint, s string) {
	println(s, " drawing", n)
	c.lock.Lock()
	defer c.lock.Unlock()
	c.account.Draw(n)
}

func (c *Card) Deposit(n uint, s string) {
	println(s, " deposting", n)
	c.lock.Lock()
	defer c.lock.Unlock()
	c.account.Deposit(n)
}

func (c *Card) Balance() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.account.Balance()
}

type CardNew struct {
	account *SimpleAccount
	deposit chan uint
	draw    chan uint
	balance chan chan int
}

func (c *CardNew) listenChan() {
	go func() {
		for {
			select {
			case dep := <-c.deposit:
				c.account.Deposit(dep)
			case draw := <-c.draw:
				c.account.Draw(draw)
			case bac := <-c.balance:
				bac <- c.account.Balance()
			}
		}
	}()
}

func NewCardChan(ac *SimpleAccount) *CardNew {
	card := &CardNew{account: ac,
		deposit: make(chan uint),
		draw:    make(chan uint),
		balance: make(chan chan int),
	}
	card.listenChan()
	return card
}

func (c *CardNew) Draw(n uint, s string) {
	println(s, " chan drawing", n)
	c.draw <- n
}

func (c *CardNew) Deposit(n uint, s string) {
	println(s, "chan deposit", n)
	c.deposit <- n
}

func (c *CardNew) Balance() int {
	tmpChan := make(chan int)
	c.balance <- tmpChan
	return <-tmpChan
}
