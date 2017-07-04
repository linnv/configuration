// Package rwlo provides ...
package rwlock

import "fmt"

var rw_ = true //only read endable
var counter = 0

type Chip struct {
	rw_ bool //only read
	ce  bool //chip enable
}

func (this *Chip) CheckCE() bool {
	return this.ce
}

func checkR() bool {
	return rw_
}

func checkW() bool {
	return !rw_
}

func StriveOpportunity() {
	rw_ = !rw_
}

func AtomAdd() {
	if checkW() {
		counter++
	}
}

func GetCounter() (int, error) {
	if checkR() {
		return counter, nil
	}
	return 0, fmt.Errorf("write only")
}
