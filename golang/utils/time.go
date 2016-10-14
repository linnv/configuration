package utils

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"
)

var dateLayout = "2006-01-02"

var SecondsPerDay int64 = 60 * 60 * 24

func ParseDate(str string) int64 {
	dayTime, err := time.ParseInLocation(dateLayout, str, time.Local)
	if err != nil {
		return 0
	}

	t := now.Now{Time: dayTime}
	beginTime, endTime := t.BeginningOfDay().Unix(), t.EndOfDay().Unix()
	fmt.Printf("beginTime: %+v,endTime: %+v of %s\n", beginTime, endTime, str)
	return t.Unix()
}

func ParseDate2Time(str string) (time.Time, error) {
	return time.ParseInLocation(dateLayout, str, time.Local)
}

func ConvertTimeint2str(seconds int64) string {
	return time.Unix(seconds, 0).Format("2006-01-02 15:04:05")
}
