package utils

import (
	"fmt"
	"testing"
)

func TestParseDate2Time(t *testing.T) {
	var eclipseDay int64 = 1
	date := "2016-04-16"
	_t, err := ParseDate2Time(date)
	CheckErr(err)
	fmt.Printf("_t: %+v\n", _t)
	fmt.Println(_t.Unix())
	fmt.Println(eclipseDay, "day(s) before ", _t.Unix(), " is ", _t.Unix()-SecondsPerDay*eclipseDay)
}
