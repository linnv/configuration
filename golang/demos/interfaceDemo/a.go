// Package main provides ...
package newDir

import (
	"fmt"
	"sunteng/commons/log"
	"time"
)

type Timer interface {
	CheckDuration() error
	GetDuration() int
}

type Moniter interface {
	GetConsumption() int
}

type Man struct {
	Timer
	Moniter
}

func (this *Man) Health() {
	if this.CheckDuration() != nil {
		fmt.Printf("Consumption: %v in duraiton%v\n", this.GetConsumption(), this.GetDuration())
	}
}

// func ReturnInterface() interface{} {
func ReturnInterface() *MyTime {
	return &MyTime{start: 111, end: 333}
}

type K uint

const (
	aSlice = iota
)

func ReturnConst() K {
	return aSlice
}

func ReturnSlice() interface{} {
	return []int{1, 23, 4}
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

type MyTime struct {
	start int `json:"Start"`
	end   int `json:"End"`
}

func (this *MyTime) CheckDuration() error {
	if this.start < 0 || this.end < 0 {
		return fmt.Errorf("ilegal time")
	}
	return nil
}

func (this *MyTime) GetDuration() int {
	return this.end - this.start
}

type Body struct {
	energy      int64 `json:"Energy"`
	consumption int64 `json:"Consumption"`
	stop        bool
	relax       chan bool
	live        bool `json:"Live"`
}

func (this *Body) AddEnergy(n int64) {
	this.energy += n
}

func (this *Body) GetConsumption() (n int64) {
	return this.consumption
}

func (this *Body) Stop() {
	for {
		select {
		case t := <-this.relax:
			if t {
				this.stop = true
				return

			}
		}
	}
}

func (this *Body) StarFiring() {
	go func() {
		for !this.live {
			if !this.stop {
				log.Logf("fired: %+v\n", this.consumption)
				this.consumption = this.consumption + 2
			}
			// this.consumption++
			time.Sleep(time.Second)
		}
	}()
}

func (this *Body) AllMembers() {
	fmt.Printf("this: %+v\n", *this)
}

type TimeInstanceA struct {
	Count int64 `json:"Count"`
}

func (ti *TimeInstanceA) CheckDuration() error {
	fmt.Printf("  time instance a: check duration\n")
	return nil
}

func (ti *TimeInstanceA) GetDuration() int {
	fmt.Printf("  time instance a:get durantin \n")
	fmt.Printf("  ti.Count: %+v\n", ti.Count)
	return 99
}

type TimeInstanceInheritanceA struct {
	TimeInstanceA
}

func (tia *TimeInstanceInheritanceA) GetDuration() int {
	fmt.Printf("  time instance inheritant a:get durantin \n")
	fmt.Printf("  tia.Count: %+v\n", tia.Count)
	tia.Count = 1999
	return 100
}

func (tia *TimeInstanceInheritanceA) UpdateCount() int64 {
	tia.Count = 1999
	return tia.Count
}

type Inheritan struct {
	N string `json:"N"`
}

func (i Inheritan) GetConsumption() int {
	return 1000
}

func (i Inheritan) AxGetConsumption() int {
	return 1022
}
