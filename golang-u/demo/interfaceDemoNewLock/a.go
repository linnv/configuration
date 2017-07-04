// Package main provides ...
package demo

import (
	"fmt"
	"math/rand"
	"sunteng/commons/log"
	"sync"
	"time"
)

type Humaner interface {
	FillEnergy(n uint)
	ConsumeEnergy(n uint) error
	GetEnergy() uint
}

type Body struct {
	Energy uint
}

func NewBody(life uint) *Body {
	return &Body{Energy: life}
}

func (b *Body) ConsumeEnergy(n uint) error {
	log.Errorf("b.Energy: %+v\n", b.Energy)
	b.latency()
	if b.Energy >= n {
		log.Errorf("b.Energy: %+v\n", b.Energy)
		b.Energy -= n
		return nil
	}
	return fmt.Errorf("jack is weak")
}

func (b *Body) FillEnergy(n uint) {
	log.Errorf("b.Energy: %+v\n", b.Energy)
	b.latency()
	log.Errorf("b.Energy: %+v\n", b.Energy)
	b.Energy += n
}

func (b *Body) latency() {
	<-time.After(time.Duration(rand.Intn(100)) * time.Millisecond)
}

func (b *Body) GetEnergy() uint {
	return b.Energy
}

type Person struct {
	body Humaner
	Name string `json:"Name"`
	lock sync.Mutex
}

func NewPerson(h Humaner, n string) *Person {
	tp := &Person{Name: n,
		body: h,
	}
	return tp
}

func (p *Person) DoFillEnergy(n uint) {
	println(p.Name, "'s power up ", n)
	p.lock.Lock()
	defer p.lock.Unlock()
	p.body.FillEnergy(n)
}

func (p *Person) DoConsumeEnergy(n uint) error {
	println(p.Name, "'s power down ", n)
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.body.ConsumeEnergy(n)
}

func (p *Person) DoGetEnergy() uint {
	// p.energyLock <- bodyChan{chanType: 9, content: nil}
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.body.GetEnergy()
}

func (p *Person) AllMember() {
	fmt.Printf("p: %+v\n", p.body)
}
