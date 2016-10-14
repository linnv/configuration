// version 1.0
package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

type TerminalCilent struct {
	sendChan     chan SendStruct
	responcechan chan ResponceStruct
}

type SendStruct struct {
	Result   int64 `json:"Result"`
	ChanId   int64
	ExtraMSG interface{} `json:"ExtraMSG "`
}

type ResponceStruct struct {
	Result   int64 `json:"Result"`
	ChanId   int64
	ExtraMSG interface{} `json:"ExtraMSG "`
}

var GolalTicketCount int64

var exitAllGoers bool

func BuckerMonitor(m *TerminalCilent) {
	var pauseDisplay int64
	for {
		m.sendChan <- SendStruct{}
		select {
		case r := <-m.responcechan:
			if pauseDisplay != r.Result {

				fmt.Println()
				log.Printf("monitoring total ticker in db: %+v\n", r.Result)
				fmt.Println()
			}
			pauseDisplay = r.Result
			// time.Sleep(time.Microsecond * 500)
			// m.sendChan <- SendStruct{}
			if exitAllGoers {
				log.Println("monitor exits: works\n")
				return
			}
		}

		// time.Sleep(time.Microsecond * 500)
		time.Sleep(time.Second)
	}
}

func DataCenter(t *TerminalCilent, m *TerminalCilent) {
	for {
		select {
		case r := <-t.sendChan:
			//@TODO do somethine with data: write to db etc.
			log.Printf("request from : %+v\n", r.ChanId)
			GolalTicketCount += r.Result
			t.responcechan <- ResponceStruct{Result: GolalTicketCount,
				ChanId:   r.ChanId,
				ExtraMSG: "node for demo"}
			if exitAllGoers {
				log.Println("DC exits: works\n")
				return
			}
			time.Sleep(time.Second * time.Duration(1))

		case <-m.sendChan:
			m.responcechan <- ResponceStruct{Result: GolalTicketCount,
				ChanId:   0,
				ExtraMSG: "just for monitor"}
			if exitAllGoers {
				log.Println("DC exits: works\n")
				return
			}
		}
	}
}

func ConsumeNode(cid int64, t *TerminalCilent, wg *sync.WaitGroup) {
	t.sendChan <- SendStruct{Result: 10 * cid,
		ChanId:   cid,
		ExtraMSG: "node for demo"}
	select {
	case <-t.responcechan:
		wg.Done()
	}
}

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(2)
	n := 100
	wg.Add(n)

	exit := make(chan bool)
	exitAllGoers = false
	t := &TerminalCilent{
		sendChan:     make(chan SendStruct),
		responcechan: make(chan ResponceStruct),
	}

	m := &TerminalCilent{
		sendChan:     make(chan SendStruct),
		responcechan: make(chan ResponceStruct),
	}
	//i is int not in64
	for i := 0; i < n; i++ {
		go ConsumeNode(int64(i), t, &wg)
	}
	go DataCenter(t, m)
	go BuckerMonitor(m)
	wg.Wait()
	<-exit
	// exitAllGoers = true
}
