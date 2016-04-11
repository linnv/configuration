// Package main provides ...
package newDir

import (
	"fmt"
	"sort"
	"time"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

func Today() string {
	return time.Now().Format("20060102")
}

type sortStrSlice []string

func (s sortStrSlice) Less(i, j int) bool {
	var iLen int
	var jLen int
	ir, jr := []rune(s[i]), []rune(s[j])
	for _, v := range ir {
		iLen += int(v)
	}
	for _, v := range jr {
		jLen += int(v)
	}
	return iLen < jLen
}

func (s sortStrSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortStrSlice) Len() int {
	return len(s)
}

func SortStrSlice(str []string) []string {
	sort.Sort(sortStrSlice(str))
	return str
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {
	w1 := "bcad"
	w2 := SortString(w1)

	fmt.Println(w1)
	fmt.Println(w2)
}

// func getDataFromMysql(day string) (data map[int64]*Camp, err error) {
// 	// var cfg dbConf
// 	// if day == today {
// 	// 	cfg = todayCfg
// 	// } else {
// 	// 	cfg = historyCfg
// 	// }
// 	macedonianDB, err := sql.Open("mymysql", fmt.Sprintf("tcp:%s:%s*%s/%s/%s",
// 		cfg.Host,
// 		cfg.Port,
// 		cfg.Database,
// 		cfg.UserName,
// 		cfg.Password))
// 	if err != nil {
// 		log.Errorf("连接 mysql 失败 :%s", err.Error())
// 		return
// 	}
// 	defer macedonianDB.Close()
// 	sqlStr := fmt.Sprintf("select campaign_id,sum(cost) as cost,sum(cost_over2) as costOver from trend_campaign_only where date = %s group by campaign_id", day)
// 	result, err := macedonianDB.Query(sqlStr)
// 	if err != nil {
// 		log.Errorf("查询数据失败:%s", err.Error())
// 		return
// 	}
// 	data = map[int64]*Camp{}
// 	for result.Next() {
// 		var id int64
// 		row := &Camp{}
// 		result.Scan(&id, &row.Cost, &row.CostOver)
// 		data[id] = row
// 	}
// 	return
// }
