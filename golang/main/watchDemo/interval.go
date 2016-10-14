package main

import (
	"sync"
	"time"
)

const (
	TypeTick = iota
	TypeHurry
)

type Content struct {
	Type int `json:"Type"`
	Girl string
}

type Body struct {
	Item      chan Content
	resetTick chan bool
	done      chan struct{}
	lifeTime  uint

	fiLock sync.Mutex
	wg     sync.WaitGroup
}

func (c *Body) subLifeTime(n uint) {
	c.fiLock.Lock()
	defer c.fiLock.Unlock()
	c.lifeTime = c.lifeTime - n
}

func (c *Body) getLifeTime() uint {
	c.fiLock.Lock()
	defer c.fiLock.Unlock()
	return c.lifeTime

}
func NewBody(lifetime uint) *Body {
	return &Body{
		Item:      make(chan Content),
		resetTick: make(chan bool),
		done:      make(chan struct{}),
		lifeTime:  lifetime,
		wg:        sync.WaitGroup{},
		fiLock:    sync.Mutex{},
	}
}

func (c *Body) Ticker() {
	c.wg.Add(1)
	// tick := time.NewTicker(time.Second * time.Duration(c.lifeTime))
	tick := time.NewTimer(time.Second * time.Duration(c.getLifeTime()))
	println("at beginging you have ", c.getLifeTime(), "seconds to escapse  from ", time.Now().String())
	for {
		select {
		case <-c.done:
			c.wg.Done()
			return
		case <-c.resetTick:
			if !tick.Stop() {
				<-tick.C
			}
			tick.Reset(time.Second * time.Duration(c.getLifeTime()))
			println("tick reset,you have ", c.getLifeTime(), "seconds to escapse  from ", time.Now().String())
		case <-tick.C:
			select {
			case c.Item <- Content{TypeTick, "阿贵"}:

			default:
			}
			println("time out ", time.Now().String())
		default:
		}
	}
}

func (c *Body) DoSomething() {
	for i := 0; i < 4; i++ {
		// select {
		// case
		//
		// case <-c.done:
		// 	println(" break ", time.Now().String())
		// 	break
		// // default:
		// // 	time.Sleep(time.Second * time.Duration(i))
		// }
		c.Item <- Content{TypeHurry, "阿贵"}
	}
	println("done agui quit")
	c.wg.Wait()
	println("exiting", time.Now().String())
}

func (c *Body) HandlerEvent() {
	c.wg.Add(1)
	for {
		select {
		case <-c.done:
			c.wg.Done()
			return
		case h := <-c.Item:
			switch h.Type {
			case TypeTick:
				println("game over", h.Girl, time.Now().String())
				//anyone who select reading this channel whill exit
				close(c.done)
				c.wg.Done()
				return
			case TypeHurry:
				//only here operate lifeTime
				c.subLifeTime(1)
				if c.getLifeTime() == 0 {
					println("congratulation all pass")
					return
				}
				println("hurry up now ", h.Girl)
				c.resetTick <- true
			}
		}
	}

}

func main() {
	g := NewBody(10)
	go g.HandlerEvent()
	go g.Ticker()
	g.DoSomething()
}
