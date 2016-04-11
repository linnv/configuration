// Package main provides ...
package newDir

import (
	"fmt"
	"strconv"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

func ParseDemo() {
	counter.PositionId, err = strconv.ParseInt(valArr[2], 10, 64)
	if err != nil {
		return counter, fmt.Errorf("parse field [%s] failed:%s", field, err.Error())
	}
}
