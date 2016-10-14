// Package utilities
package utils

import (
	// "fmt"
	"sort"
	"strconv"
	"strings"
)

func Sort(ids []int) []int {
	is := sort.IntSlice(ids)
	is.Sort()
	return is
}

func StrArrayToStrWithSpliter(str []string, spliter string) (result string) {
	for i, v := range str {
		// fmt.Printf("v %+v %d  len %d\n", v, i, len(v))
		if len(strings.Trim(v, " ")) != 0 { //skip empty string
			v = strings.Trim(v, " ")
			if i != 0 {
				result = result + spliter + v
			} else {
				result = result + v
			}
		}
	}
	return
}

func StrToStringArray(str, spileType string) []string {
	tmpArray := strings.Split(strings.Trim(str, " "), spileType)
	var result = make([]string, 0, len(tmpArray))
	for _, v := range tmpArray {
		if len(strings.Trim(v, " ")) != 0 { //skip empty string
			v = strings.Trim(v, " ")
			result = append(result, v)
		}
	}
	return result
}

func stringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	return jsons
}
