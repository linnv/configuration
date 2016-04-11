// Package main provides ...
package newDir

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// const length string diff from  []
var strList = [...]string{
	"g",
	"d",
}

func BitCalucate() {
	fmt.Printf("A: %v B:%v C:%v \n", strconv.FormatInt(int64(A), 2), strconv.FormatInt(int64(B), 2), strconv.FormatInt(int64(C), 2))
	fmt.Printf("AB: %v\n", strconv.FormatInt(int64(AB), 2))
	fmt.Printf("AC: %v\n", strconv.FormatInt(int64(AC), 2))
	fmt.Printf("BC: %v\n", strconv.FormatInt(int64(BC), 2))
	fmt.Printf("ABC: %v\n", strconv.FormatInt(int64(ABC), 2))
}

const jj = "imge/"

func strFunc(fmtFunc func(format string, a ...interface{}) string, format string, args ...interface{}) {
	//@toDelete
	fmt.Printf("fmtFunc(format,args): %+v\n", fmtFunc(format, args))
}

func Llog(level string, fmtFunc func(...interface{}) string, args ...interface{}) {
	v := make([]interface{}, 1, len(args)+1)
	v[0] = "[" + level + "] "
	v = append(v, args...)

	str := fmtFunc(v...)
	fmt.Printf("str: %+v\n", str)
}

func Llogf(level, format string, fmtFunc func(string, ...interface{}) string, args ...interface{}) {
	v := make([]interface{}, 1, len(args)+1)
	v[0] = "[" + level + "] "
	v = append(v, args...)

	str := fmtFunc(format, v...)
	fmt.Printf("str: %+v\n", str)
}

func ReturnTrue() bool {
	os.Stdout.Write(append([]byte("true"), '\n'))
	return true
}

func ReturnFalse() bool {
	os.Stdout.Write(append([]byte("false"), '\n'))
	return false
}

func JustDemo() {
	println("//<<-------------------------JustDemo start-----------")
	start := time.Now()
	CountOneDemo()
	// SwapMetrix()
	// str := "-33fjefej"
	// bs := []byte(str)
	// const minuteChart = 45 //"-"
	// for k, v := range bs {
	// 	if v == minuteChart {
	// 		fmt.Printf("string(v): %+v\n", string(v))
	// 	}
	// 	fmt.Printf("%+v: %+v\n", k, v)
	// }
	// i := 0
	// for i < 100 {
	// 	fmt.Printf("i: %+v\n", i)
	// 	i += 90
	// }

	// fmt.Printf("1<<5: %+v\n", 1<<5)
	// fmt.Printf("1<<7: %+v\n", 1<<7)
	// fmt.Printf("1<<0: %+v\n", 1<<0)
	// fmt.Printf("5e9: %f\n", 5e3)

	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 100; j++ {
	// 		if j == 33 {
	// 			os.Stdout.Write(append([]byte("break at 33"), '\n'))
	// 			break
	// 		}
	// 	}
	// }

	// m := make(map[int]string, 10)
	// fmt.Printf("  len(m): %+v\n", len(m))
	// s := "[%d]:[%v]"
	// a := fmt.Sprintf(s, 11, []int{1, 24})
	// //@toDelete
	// fmt.Printf("  a: %+v\n", a)
	// str := []string{"a", "b", "c"}
	// tmpStr := make([]string, len(str), len(str))
	// // tmpStr := make([]string, 0, len(str))
	// // var tmpStr []string
	// copy(tmpStr, str)
	// // copy(str, tmpStr)
	// fmt.Printf("tmpStr: %+v\n", tmpStr)
	// for i := 0; i < len(tmpStr); i++ {
	// 	tmpStr[i] = strings.ToUpper(tmpStr[i])
	// }
	// fmt.Printf("tmpStr: %+v\n", tmpStr)

	// name := "1jfje/2fejfej/8fejfe/a"
	// arr := strings.Split(name, "/")
	// //@toDelete
	// fmt.Printf("arr: %+v\n", arr)
	// rarr := strings.Join(arr[1:], "/")
	// //@toDelete
	// fmt.Printf("rarr: %+v\n", rarr)
	//
	// // host := "xxxx:eere"
	// host := "xxxx"
	// hp := strings.Split(host, ":")
	// fmt.Printf("hp: %+v\n", hp)
	//
	// // a, b := true, false
	// if ReturnFalse() && ReturnTrue() {
	// 	os.Stdout.Write(append([]byte("both true"), '\n'))
	// }

	// fmt.Printf("  strList[FIRST]: %+v\n", strList[FIRST])
	//
	// s := jj + "333"
	// fmt.Printf("  s: %+v\n", s)
	// var appIds, normalIds, illegalIds []int
	// appIds, normalIds, illegalIds = make([]int, 0, 33)
	fmt.Printf("JustDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------JustDemo end----------->>")
}

func BufferDemo() {

	var b bytes.Buffer
	for i := 0; i < 10; i++ {
		fmt.Fprintf(&b, ";line %d", i)
	}
	fmt.Printf("  b.String(): %+v\n", b.String())
	fmt.Printf("  b.Bytes(): %+v\n", b.Bytes())
}

func Trap() {
	arr := []int{1, 2, 3, 4, 5, 10}
	slice := arr[1:2]
	fmt.Println(slice)
	//now the underly pointer points to index 2th of arr, append() add or update  element from this index of arr
	slice = append(slice, 6, 7, 8)
	fmt.Println(slice)
	fmt.Println(arr)
}

func InIntArray(i int, ints []int) bool {
	for _, v := range ints {
		if i == v {
			return true
		}
	}
	return false
}

func IntArraySubtract(a, b []int) (c []int) {
	c = []int{}

	for _, _a := range a {
		if !InIntArray(_a, b) {
			c = append(c, _a)
		}
	}

	return
}

func IntArrayCoincide(a, b []int) (c []int) {
	c = []int{}

	for _, _a := range a {
		if InIntArray(_a, b) {
			c = append(c, _a)
		}
	}

	return
}

func SliceDemo() {
	// appAdslotIds := make([][]int, 2)
	// for i := 0; i < len(appAdslotIds); i++ {
	// 	appAdslotIds[i] = make([]int, 0, 10)
	// 	// appAdslotIds[i] = make([]int, 10)
	// }
	// fmt.Printf("  appAdslotIds[0]: %+v cap:%v\n", appAdslotIds[0], cap(appAdslotIds[0]))
	//
	// fmt.Printf("  appAdslotIds[1]: %+v cap:%v\n", appAdslotIds[1], cap(appAdslotIds[0]))

	a, b := make([]string, 0, 2), make([]string, 0, 2)

	fmt.Printf("  len(a) %+v,cap(b): %+v\n", len(a), cap(a))
	fmt.Printf("  len(a) %+v,cap(b): %+v\n", len(b), cap(b))

	c := "[\""
	for i := 0; i < 10; i++ {
		c += strconv.Itoa(i) + "\",\""
	}
	c += "]"
	fmt.Printf("  c: %+v\n", c)
}

func PointersSliceDemo(as []*AA) {
	for k, v := range as {
		v.Name = "update" + strconv.Itoa(k)
	}
}
func instanceSliceDemo(as []AA) {
	for k, v := range as {
		v.Name = "update" + strconv.Itoa(k)
	}
}

func Array(a []int) {
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
}

func ArrayInstance(a *[3]int) {
	for i := 0; i < len(a); i++ {
		a[i] = 3
	}
}

func ReplaceStr() {

	str := `"Bid_M.render({
	"id": "323903",
	"name": "{PRODUCT_NAME}",
	"desc": "{PRODUCT_SUMMARY}",
	"icon": "{PRODUCT_ICON_NAME}",
	"banner": "{BANNER_NAME}",
	"link": "{WEB_URL}",
	"isLink": {IS_OPEN_URL},
	"star": {PRODUCT_RATE},
	"count": {PRODUCT_DOWNLOADS},
	"ios": {
		"download": "{ITUNES_URL}",
		"dplink": "{IOS_DEEP_URL}"
	},
	"android": {
		"download": "{APK_URL}",
		"dplink": "{ANDROID_DEEP_URL}"
	}
});
	"`

	// c := "[\""
	// for i := 0; i < 10; i++ {
	// 	if i != 9 {
	// 		c += strconv.Itoa(i) + "\",\""
	// 	} else {
	// 		c += strconv.Itoa(i) + "\""
	// 	}
	//
	// }
	// c += "]"
	c := `\["sss","eee"\]`

	b, err := json.Marshal(c)
	if err != nil {
		return
	}
	cstr := strconv.Quote(string(b))
	fmt.Printf("  cstr: %+v\n", cstr)
	// fmt.Printf("  string(b): %+v\n", string(b))
	// c := []string{"a", "b", "c"}
	replacer := strings.NewReplacer(
		"{PRODUCT_NAME}", "xx",
		"{PRODUCT_SUMMARY}", "xx",
		"{PRODUCT_ICON_NAME}", "33",
		"{BANNER_NAME}", cstr[1:len(cstr)-1],
		// "{BANNER_NAME}", c,
		"{WEB_URL}", "xx",
		"{IS_OPEN_URL}", "isOpenUrl",
		"{PRODUCT_RATE}", "strconv",
		"{PRODUCT_DOWNLOADS}", "xx",
		"{ITUNES_URL}", "params",
		"{IOS_DEEP_URL}", "params",
		"{APK_URL}", "xxe3",
		"{ANDROID_DEEP_URL}", "AwkDeepUrl",
	)
	str = replacer.Replace(str)
	fmt.Printf("  str: %+v\n", str)
	fmt.Printf("  str: %v\n", str)
}

func GetSuffixTypeByDot(str string) string {
	// r := strings.LastIndexByte(str, byte("."))
	r := strings.LastIndex(str, ".")
	return str[r:]
	// fmt.Printf("  r: %+v\n", r)
	// fmt.Printf("  str[r:]: %+v\n", str[r:])
}

var (
	invalidReplacer *strings.Replacer
)

func CharactersReplaceDemo(s string) (str string) {
	println("//<<-------------------------CharactersDemo start-----------")
	fmt.Printf("  oriign s: %+v\n", s)
	invalidReplacer = strings.NewReplacer("\t", "", "|", "\\|", "\r", "~")
	str = invalidReplacer.Replace(s)
	fmt.Printf("result  str: %+v\n", str)
	println("//---------------------------CharactersDemo end----------->>")
	return
}

func EmptySliceDemo() {
	println("//<<-------------------------EmptySliceDemo start-----------")
	emptyStr := []string{""}
	//@todoDelelte
	fmt.Printf("emptyStr,cap(emptyStr):%d,len(emptyStr):%d  : %+v\n", cap(emptyStr), len(emptyStr), emptyStr)
	// ss := make([]string, 0, 2)
	// //@toDelete
	// fmt.Printf("  len(ss): %+v\n", len(ss))
	// ss = append(ss, emptyStr)
	// fmt.Printf("  after len(ss): %+v\n", len(ss))
	println("//---------------------------EmptySliceDemo end----------->>")
}

func SortInt64Demo() {
	println("//<<-------------------------SortInt64Demo start-----------")
	ii := make(Int64Slice, 0, 100)
	var i int64
	for i = 0; i < 199; i++ {
		ii = append(ii, i)
	}
	sort.Sort(ii)
	//@toDelete
	fmt.Printf("  ii: %+v\n", ii)
	ii = ii[:0]
	ii = append(ii, 11)
	//@toDelete
	fmt.Printf(" after ii: %+v\n", ii)
	c := time.Now().UnixNano()
	//@toDelete
	fmt.Printf("  c: %+v\n", c)
	fmt.Printf("  c: %+v\n", strconv.FormatInt(c, 10)[7:10])
	println("//---------------------------SortInt64Demo end----------->>")
}

func LoopDemo() {
	println("//<<-------------------------LoopDemo start-----------")
	var dspIdsNotExists []int
	dspIds := []int{1, 2, 3}
	dspList := []int{2, 3}
loopPostDspIds:
	for _, id := range dspIds {
		//@toDelete
		fmt.Printf("  id: %+v\n", id)
		for _, d := range dspList {
			if d == id {
				continue loopPostDspIds
			}
		}
		dspIdsNotExists = append(dspIdsNotExists, id)
	}
	//@toDelete
	fmt.Printf("  dspIdsNotExists: %+v\n", dspIdsNotExists)
	println("//---------------------------LoopDemo end----------->>")
}

func RetSlice() []int {
	// as := make([]int, 0, len(10))
	var as []int
	for i := 0; i < 10; i++ {
		as = append(as, i)
	}
	//@todoDelelte
	fmt.Printf("as: %+v,cap(as):%d,len(as):%d  \n", as, cap(as), len(as))
	return as[:5]
}

var as []int

func init() {
	as = make([]int, 0, 100)
	for i := 6; i < 12; i++ {
		as = append(as, i)
	}

	for i := 0; i < 6; i++ {
		as = append(as, i)
	}
}

func BinarySearch() int {
	target := 0
	first, mid, last := 0, 0, len(as)
	//@todoDelelte
	fmt.Printf("as: %+v,cap(as):%d,len(as):%d  \n", as, cap(as), len(as))
	for first != last {
		mid = first + (last-first)/2
		if as[mid] == target {
			return mid
		}
		if as[mid] < as[first] {
			if as[first] <= target && target < as[mid] {
				last = mid
			} else {
				first = mid + 1
			}
		} else if as[mid] > as[first] {
			if as[mid] <= target && target < as[last] {
				first = mid + 1
			} else {
				last = mid
			}
		} else {
			first++
		}
	}

	return -1
}

func LongConsecutiveSequence() int {
	giveSlice := []int{100, 4, 200, 1, 3, 2, 5}
	longest := 0
	used := make(map[int]bool)
	giveSliceLen := len(giveSlice)

	for i := 0; i < giveSliceLen; i++ {
		used[giveSlice[i]] = false
	}

	for i := 0; i < giveSliceLen; i++ {
		if used[giveSlice[i]] {
			continue
		}

		curLength, t := 0, giveSlice[i]

		for m := i + 1; m < giveSliceLen; m++ {
			t++
			if _, ok := used[t]; ok {
				curLength++
				used[giveSlice[t]] = true
			}
		}
		for m := i - 1; m > 0; m-- {
			t--
			if _, ok := used[t]; ok {
				curLength++
				used[giveSlice[t]] = true
			}
		}

		longest = Max(longest, curLength)
	}

	return longest
}

func TwoSum() []int {
	giveSlice, target := []int{2, 7, 11, 15}, 9
	giveSliceLen := len(giveSlice)
	sliceMap, ret := make(map[int]int, giveSliceLen), make([]int, 0, 2)

	for i := 0; i < giveSliceLen; i++ {
		sliceMap[giveSlice[i]] = i
	}
	for i := 0; i < giveSliceLen; i++ {
		gap := target - giveSlice[i]
		if j, ok := sliceMap[gap]; ok && j > i {
			ret = append(ret, i+1, j+1)
			break
		}
	}

	return ret
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SlicePointer() {
	ar := []int{1, 2, 34, 66}
	fmt.Printf("ar: %+v,cap(ar):%d,len(ar):%d  \n", ar, cap(ar), len(ar))
	ip := new(int)

	ip = &ar[0]
	fmt.Printf("ip: %+v\n", ip)
	fmt.Printf("*ip: %+v\n", *ip)

}

// func TreeSum() []int {
// 	giveSlice, target := []int{2, 7, 11, 15}, 9
// 	giveSliceLen := len(giveSlice)
// 	nextIndex, lastIndex, min = 0, 0, 0
//
// 	for i := 0; i < giveSliceLen; i++ {
// 		nextIndex, lastIndex = i+1, giveSliceLen-1
// 		gap := target - giveSlice[i]
// 	}
//
// }

func NewNode(str string) *Node {
	return &Node{str, nil}
}

func GetLinkList() *Node {
	root := NewNode("root")
	tmp := root
	for i := 0; i < 10; i++ {
		tmp.Next = NewNode(strconv.Itoa(i))
		tmp = tmp.Next
	}
	return root
}

func ReverseLinkList() {
	ll := GetLinkList()
	ll.PrintAll()
	pre := new(Node)
	target := ll
	for target != nil {
		tmp := target.Next
		target.Next = pre
		pre = target
		target = tmp
	}
	os.Stdout.Write(append([]byte("reverse"), '\n'))
	pre.PrintAll()
}

func tErr(a []int, err ...error) {
	if len(a) < 1 {
		os.Stdout.Write(append([]byte("good"), '\n'))
		return
	}
	if len(err) > 0 && err[0] != nil {
		panic(err[0].Error())
		return
	}
	os.Stdout.Write(append([]byte("good"), '\n'))
}

func CountOneDemo() {
	println("//<<-------------------------CountOneDemo start-----------")
	start := time.Now()
	// fmt.Printf("strconv.IntSize: %+v\n", strconv.IntSize)
	num := 2
	bitsTotal := uint(unsafe.Sizeof(int64(0)))
	// fmt.Printf("size(int64): %+v\n", unsafe.Sizeof(int64(0)))
	bitCountArray := make([]int, bitsTotal)
	var i uint
	for i = 0; i < bitsTotal; i++ {
		bitCountArray[i] += (num >> i) & 1
	}
	fmt.Printf("bitCountArray: %+v\n", bitCountArray)

	fmt.Printf("CountOneDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------CountOneDemo end----------->>")
}

func SwapMetrix() {
	edgeLenth := 50
	metrixSlice := make([][]string, edgeLenth)
	// metrixSlice := [edgeLenth][]string{}
	for i := 0; i < edgeLenth; i++ {
		metrixSlice[i] = make([]string, edgeLenth)
		for j := 0; j < edgeLenth; j++ {
			if i == j {
				// os.Stdout.Write([]byte("*"))
				metrixSlice[i][j] = "-"
				if i == 0 && j == 0 {
					fmt.Printf("metrixSlice[i][j] : %+v\n", metrixSlice[i][j])
				}
				continue
			} else if i < j {
				metrixSlice[i][j] = "8"
				// os.Stdout.Write([]byte("-"))
				continue
			}
			metrixSlice[i][j] = "6"
			// os.Stdout.Write([]byte("#"))
		}
		// println()
	}

	for i := 0; i < edgeLenth; i++ {
		fmt.Printf("%s\n", metrixSlice[i])
	}
	println("------ pause \n========================\n")

	// reverse metrix
	for i := 0; i < edgeLenth/2; i++ {
		for j := 0; j < edgeLenth; j++ {
			tmp := metrixSlice[i][j]
			metrixSlice[i][j] = metrixSlice[edgeLenth-1-i][edgeLenth-1-j]
			metrixSlice[edgeLenth-1-i][edgeLenth-1-j] = tmp
		}
	}

	for i := 0; i < edgeLenth; i++ {
		fmt.Printf("%s\n", metrixSlice[i])
	}
}
