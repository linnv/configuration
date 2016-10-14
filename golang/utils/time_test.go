package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestParseDate2Time(t *testing.T) {
	// date := "2016-04-19"
	// var eclipseDay int64 = 1
	// _t, err := ParseDate2Time(date)
	// CheckErr(err)
	// fmt.Printf("_t: %+v\n", _t)
	// fmt.Println(_t.Unix())
	// fmt.Println(eclipseDay, "day(s) before ", _t.Unix(), " is ", _t.Unix()-SecondsPerDay*eclipseDay)

	// ParseDate(date)

	nowTime := time.Now()
	fmt.Printf("nowTime.Unix(): %+v\n", nowTime.Unix())
	d := time.Duration(int(nowTime.YearDay()*5)+1) * 24 * time.Hour
	nn := nowTime.Truncate(time.Hour).Add(d).Unix()
	fmt.Printf("nn: %+v\n", nn)

}
