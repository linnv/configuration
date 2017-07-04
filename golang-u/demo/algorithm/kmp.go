package demo

import (
	"fmt"
	"strings"
	"time"
)

//key points: get a pattern table and use the pattern table
//compare count of chraracter match both prefix and suffix properly, e.g. `ababc` -> c:ababc:0,b:abab:2(ab),a:aba:1(a),b:ab:0,a:a:0(exception)
//get skip count from pattern table by `(matchedLength-patterTable[matchedLength-1])`
//format:[]int->[index]value
//Getpattern return a pattern table for str
func GetPattern(str string) (pattern []int) {
	strLen := len(str)
	//skip: skip count, also index beginning
	byteMatchIndex := func(s string, char byte, skip int) int {
		sLen := len(s)
		if skip >= sLen {
			return -1
		}
		for i := skip; i < sLen; i++ {
			if s[i] == char {
				return i
			}
		}
		return -1
	}

	pattern = make([]int, strLen)
	for i := strLen - 1; i >= 1; i-- {
		matchIndex := byteMatchIndex(str, str[i], 0)
		// logx.Debug("match %d ,first i: %+v c:%c\n", matchIndex, i, str[i])
		// logx.Debug("string(str[:matchIndex+1]): %+v\n", string(str[:matchIndex+1]))
		if matchIndex > -1 && matchIndex < i {
			goodJob := true
			for j := matchIndex; j >= 0; j-- {
				if str[i-matchIndex+j] != str[j] {
					goodJob = false
					// logx.Debug("not proper match: %s VS %s\n", string(str[:matchIndex+1]), str[i-matchIndex:])
					break
				}
			}
			if !goodJob {
				continue
			}
			pattern[i] = matchIndex + 1
		} else {
			continue
		}
		filled := false
	higher:
		if filled {
			continue
		}
		matchIndex = byteMatchIndex(str, str[i], matchIndex+1)
		// logx.Debug("higher match %d ,first i: %+v c:%c\n", matchIndex, i, str[i])
		// logx.Debug("string(str[:matchIndex+1]): %+v\n", string(str[:matchIndex+1]))
		if matchIndex > -1 && matchIndex < i {
			goodJob := true
			for j := matchIndex; j >= 0; j-- {
				if str[i-matchIndex+j] != str[j] {
					goodJob = false
					filled = true
					// logx.Debug("higher not proper match: %s VS %s\n", string(str[:matchIndex+1]), str[i-matchIndex:])
					break
				}
			}
			if !goodJob {
				continue
			}
			pattern[i] = matchIndex + 1
			goto higher
		}
	}
	return
}

func KMP(dst, sour string) int {
	patternTable := GetPattern(dst)
	dstLen := len(dst)
	sourLen := len(sour)
	for i := 0; i < sourLen; {
		// if sour[i] == dst[0] && (i+dstLen) <= sourLen {
		if sour[i] == dst[0] {
			if (i + dstLen - 1) >= sourLen {
				// logx.Debug("what: %+v\n", "what")
				// logx.Debug("i: %+v left str %s\n", i, string(sour[i:]))
				return -1
			}
			found := true
			for j := 1; j < dstLen; j++ {
				//ok we can decide how many steps to skip from querying pattern
				if sour[i+j] != dst[j] {
					ahead := j - patternTable[j-1]
					// logx.Debug("i %d matchLength: %+v\n", i, j)
					// logx.Debug("ahead: %+v\n", ahead)
					if ahead == 0 {
						ahead = 1
					}
					i += ahead
					found = false
					break
				}
			}
			if found {
				// logx.Debug("bingo i: %+v through str %s\n", i, string(sour[:i]))
				// logx.Debug("left string: %+v\n", string(sour[i:]))
				return i
			}
			continue
		}
		i++
	}
	return -1
}

func KMPDemo(d, s string) int {
	println("//<<-------------------------KMPDemo start-----------")
	start := time.Now()
	i := KMP(d, s)
	fmt.Printf("KMPDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------KMPDemo end----------->>")
	return i
}

func LibIndexDemo(d, s string) int {
	println("//<<-------------------------LibIndexDemo start-----------")
	start := time.Now()
	i := strings.Index(s, d)
	fmt.Printf("LibIndexDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------LibIndexDemo end----------->>")
	return i
}
