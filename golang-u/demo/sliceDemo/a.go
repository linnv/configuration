package demo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/linnv/logx"
)

// const length string diff from  []
var strList = [...]string{
	"g",
	"d",
	4: "e",
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

func FunSlice(a []int) []int {
	return a[1:]
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
	// arr := []int{1, 2, 3, 4, 5, 10}
	// slice := arr[1:2]
	// fmt.Println(slice)
	// //now the underly pointer points to index 2th of arr, append() add or update  element from this index of arr
	// slice = append(slice, 6, 7, 8)
	// fmt.Println(slice)
	// fmt.Println(arr)
	ss := make([]int, 10)
	for i := 0; i < 10; i++ {
		ss[i] = i
	}
	//@todoDelelte
	fmt.Printf("ss: %+v,cap(ss):%d,len(ss):%d  \n", ss, cap(ss), len(ss))
	ss = ss[:0]
	fmt.Printf("after ss: %+v,cap(ss):%d,len(ss):%d  \n", ss, cap(ss), len(ss))
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

func PinterDemo() {
	println("//<<-------------------------PinterDemoDemo start-----------")
	start := time.Now()
	h := new(Node)
	// var hh *Node
	// if hh == nil {
	// 	println(" var and nil")
	// }

	fmt.Printf("h: %p\n", h)
	h = NewNode("header")
	ah := NewNode(" auto header")
	fmt.Printf("header: %p\n", h)
	fmt.Printf(" auto header: %p\n", ah)
	h = ah
	fmt.Printf("header: %p\n", h)
	fmt.Printf("PinterDemoDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------PinterDemoDemo end----------->>")
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

func FindCenterNodeLinkList() {
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")
	d.Next = e
	c.Next = d
	b.Next = c
	a.Next = b
	a.PrintAll()

	var nSlow, nFast *Node
	nSlow, nFast = a, a
	for nFast.Next != nil {
		fmt.Printf("nSlow: %x\n", nSlow)
		fmt.Printf("nFast: %x\n", nFast)
		nSlow = nSlow.Next
		//even count of link nodes
		if nFast.Next.Next == nil {
			nFast = nFast.Next
			break
		}
		//prime count of link nodes
		nFast = nFast.Next.Next
	}

	println("nSlow ", nSlow.Data)
	println("nFast ", nFast.Data)
}

func SimpleReverseLinkList() {
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	b.Next = c
	a.Next = b
	a.PrintAll()
	println("------ pause \n========================\n")
	// tmp := new(Node)
	// pre, post := new(Node), a
	// for post != nil {
	// 	last := post.Next
	// 	post.Next = pre
	// 	pre = post
	// 	if last == nil {
	// 		break
	// 	}
	// 	post = last
	// }
	//
	// pre, post := a, a.Next
	// a.Next = nil
	// for post != nil {
	// 	last := post.Next
	// 	post.Next = pre
	// 	pre = post
	// 	if last == nil {
	// 		break
	// 	}
	// 	post = last
	// }

	// post := a.Next
	// a.Next = nil
	// for post != nil {
	// 	last := post.Next
	// 	post.Next = a
	// 	a = post
	// 	if last == nil {
	// 		break
	// 	}
	// 	post = last
	// }
	// a = post

	// pre := a
	// a = a.Next
	pre, a := a, a.Next
	pre.Next = nil
	for a != nil {
		last := a.Next
		a.Next = pre
		pre = a
		// if last == nil {
		// 	break
		// }
		a = last
	}
	a = pre
	a.PrintAll()
}

func ReverseLinkList() {
	ll := GetLinkList()
	ll.PrintAll()
	pre := new(Node)
	target := ll
	for target != nil {
		last := target.Next
		target.Next = pre
		pre = target
		target = last
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

func kpHashDemo() {
	println("//<<-------------------------kpHashDemo start-----------")
	start := time.Now()

	// primeRK is the prime base used in Rabin-Karp algorithm.
	const primeRK = 16777619

	// hashStr returns the hash and the appropriate multiplicative
	// factor for use in Rabin-Karp algorithm.
	hashStr := func(sep string) (uint32, uint32) {
		hash := uint32(0)
		for i := 0; i < len(sep); i++ {
			hash = hash*primeRK + uint32(sep[i])
		}
		var pow, sq uint32 = 1, primeRK
		for i := len(sep); i > 0; i >>= 1 {
			if i&1 != 0 {
				pow *= sq
			}
			sq *= sq
		}
		return hash, pow
	}

	// hashStrRev returns the hash of the reverse of sep and the
	// appropriate multiplicative factor for use in Rabin-Karp algorithm.
	hashStrRev := func(sep string) (uint32, uint32) {
		hash := uint32(0)
		for i := len(sep) - 1; i >= 0; i-- {
			hash = hash*primeRK + uint32(sep[i])
		}
		var pow, sq uint32 = 1, primeRK
		for i := len(sep); i > 0; i >>= 1 {
			if i&1 != 0 {
				pow *= sq
			}
			sq *= sq
		}
		return hash, pow
	}

	str := "aaa"
	hashStrA, powA := hashStr(str)
	fmt.Printf("hashStrA,: %+v\n", hashStrA)
	fmt.Printf("powA: %+v\n", powA)

	hpa, ppa := hashStrRev(str)
	fmt.Printf("hpa: %+v\n", hpa)
	fmt.Printf("ppa: %+v\n", ppa)

	fmt.Printf("10: %b\n", 10)
	for i := 10; i > 0; i >>= 1 {
		if i&1 != 0 {
			fmt.Println("non zero bits")
		}
		fmt.Printf("i: %+v\n", i)
	}
	fmt.Printf("kpHashDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------kpHashDemo end----------->>")
}

func LRULinkListDemo() {
	println("//<<-------------------------LRULinkListDemo start-----------")
	start := time.Now()

	fmt.Printf("LRULinkListDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------LRULinkListDemo end----------->>")
}

func BufReadDemoPointer(buf *[]int) {
	println("//<<-------------------------BufReadDemo start-----------")
	start := time.Now()
	// io.ReadFull()
	*buf = append(*buf, 222)
	// buf = buf[0:]
	// buf[0] = 112
	fmt.Printf("BufReadDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------BufReadDemo end----------->>")
}

func BufReadDemo(buf []int) {
	println("//<<-------------------------BufReadDemo start-----------")
	start := time.Now()
	// io.ReadFull()
	fmt.Printf("go in2 buf: %+v,cap(buf):%d,len(buf):%d arrd:%p \n", buf, cap(buf), len(buf), buf)
	// BufReadDemoPointer(&buf)
	// buf = buf[len(buf):]
	// buf = append(buf, 222)
	// buf = append(buf, 222)
	buf = buf[1:]
	buf[0] = 112
	fmt.Printf("in2 buf: %+v,cap(buf):%d,len(buf):%d arrd:%p \n", buf, cap(buf), len(buf), buf)
	fmt.Printf("BufReadDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------BufReadDemo end----------->>")
}

func JustDemo() {
	println("//<<-------------------------JustDemo start-----------")
	start := time.Now()
	a := make([][2]int, 0, 1)
	a = append(a, [2]int{1, 2})
	a = append(a, [2]int{1, 2})
	logx.Debug("a: %+v\n", a)
	// const good = "jfeifejfej fejfe efj"
	// bs := make([]byte, len(good))
	// pr, pw := io.Pipe()
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go func() {
	// 	n, err := pw.Write([]byte(good))
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	if len("good") > n {
	// 		panic("less")
	// 		return
	// 	}
	// 	wg.Done()
	// }()
	//
	// // pr.Read(bs)
	// io.ReadFull(pr, bs)
	// log.Printf("bs: %+v\n", bs)
	// log.Printf("bs: %+v\n", string(bs))
	// err := pr.Close()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// err = pw.Close()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// wg.Done()
	// wg.Wait()
	//
	// for i := 0; i < 5; i++ {
	// 	defer fmt.Printf("%d ", i)
	// }
	// buf := make([]int, 3, 10)
	// buf[0] = 111
	// fmt.Printf("1 buf: %+v,cap(buf):%d,len(buf):%d arrd:%v \n", buf, cap(buf), len(buf), &buf[0])
	// BufReadDemo(buf)
	// // BufReadDemoPointer(&buf)
	// // buf = append(buf, 222)
	// //@todoDelelte
	// fmt.Printf("final buf: %+v,cap(buf):%d,len(buf):%d arrd:%v \n", buf, cap(buf), len(buf), &buf[0])

	// rul := NewLRULink()
	// for i := 0; i < 201; i++ {
	// 	rul.Upsert(strconv.Itoa(i))
	// }
	// rul.Upsert("19")
	// rul.PrintAll()

	// kpHashDemo()
	// var s string = "jjj3"
	// fmt.Printf("len(s): %+v\n", len(s))
	// fmt.Printf("s[0]: %+v\n", s[0])
	// fmt.Printf("s[0]: %+v\n", string(s[0]))
	// if strings.Contains("aaeje j", "e j") {
	// 	fmt.Println("contain e j")
	// }

	// Trap()
	// FindCenterNodeLinkList()
	// SimpleReverseLinkList()
	// PinterDemo()
	// ReverseLinkList()
	// CountOneDemo()
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

	fmt.Printf("  strList[FIRST]: %+v\n", strList[FIRST])
	fmt.Printf("  strList[4]: %+v\n", strList[4])

	// s := jj + "333"
	// fmt.Printf("  s: %+v\n", s)
	// var appIds, normalIds, illegalIds []int
	// appIds, normalIds, illegalIds = make([]int, 0, 33)
	// a := []int{1, 24, 5}
	// b := [][]int{a}
	// fmt.Printf("b: %+v\n", b)
	// fmt.Printf("b[0]: %+v\n", b[0])

	// var b []int
	// a := []int{1, 2, 3, 4, 5}
	// fmt.Printf("len(a): %+v\n", len(a))
	// fmt.Printf("a[1:] %+v\n", a[1:])
	// for k, v := range a {
	// 	//@toDelete
	// 	fmt.Printf("%+v: %+v\n", k, v)
	// }
	// b = FunSlice(a)
	// fmt.Printf("a: %+v\n", a)
	// b[2] = 999 //under line array of b is a, this operation will affect slice a
	// fmt.Printf("a: %+v\n", a)
	// a[4] = 10000
	// fmt.Printf("b: %+v\n", b)

	// s := strings.FieldsFunc("jjfe fejf ejw s  jxoe fejw ", unicode.IsSpace)
	// s := strings.Split("jjfe,fejf,ejw s  jxoe fejw ", ",")
	// fmt.Printf("len(s): %+v\n", len(s))
	// fmt.Printf("s: %+v\n", s)

	// bs := make([]byte, 1)
	// bs[len(bs)-1] = 1
	// fmt.Printf("bs: %+v,cap(bs):%d,len(bs):%d arrd:%v \n", bs, cap(bs), len(bs), &bs[0])
	// bs = bs[:len(bs)-1]
	// log.Printf("len(bs): %+v\n", len(bs))
	// fmt.Printf("o bs: %+v,cap(bs):%d,len(bs):%d arrd:%v \n", bs, cap(bs), len(bs), &bs[0])
	// ss := "jia"
	// n := copy(bs, ss)
	// fmt.Printf("bs: %+v\n", string(bs))
	// fmt.Printf("bs: %v\n", bs)
	// fmt.Printf("bs: %T\n", bs)
	// fmt.Printf("n: %+v\n", n)

	// fmt.Println(strings.Join([]string{"j", "i", "a"}, "----"))

	fmt.Printf("JustDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------JustDemo end----------->>")
}
func DiffPosStrDemo() {
	println("//<<-------------------------DiffPosStrDemo start-----------")
	start := time.Now()
	a := "abc"
	b := "bac"
	var chara, charb uint
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			chara += uint(a[i])
		}
		for i := 0; i < len(b); i++ {
			charb += uint(b[i])
		}
		if chara == charb {
			println("good")
		}
	}

	fmt.Printf("DiffPosStrDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------DiffPosStrDemo end----------->>")
}

type BASTRUCT struct {
	BN string `json:"BN"`
}

type ASTRUCT struct {
	BASTRUCT *BASTRUCT `json:"BASTRUCT"`
	N        string    `json:"N"`
}

func IssueDemo() {
	println("//<<-------------------------IssueDemo start-----------")
	start := time.Now()
	// a := []int{1, 2}
	a := []*ASTRUCT{&ASTRUCT{&BASTRUCT{"bOne"}, "one"}, &ASTRUCT{&BASTRUCT{"bTwo"}, "two"}}
	b := []*ASTRUCT{&ASTRUCT{&BASTRUCT{"bOne"}, "one"}, &ASTRUCT{&BASTRUCT{"bTwo"}, "two"}}
	fmt.Printf("a: %+v,cap(a):%d,len(a):%d arrd:%v \n", a, cap(a), len(a), &a[0])
	fmt.Printf("b: %+v,cap(b):%d,len(b):%d arrd:%v \n", b, cap(b), len(b), &b[0])
	a, b = a[0:], b[0:]
	fmt.Printf("a: %+v,cap(a):%d,len(a):%d arrd:%v \n", a, cap(a), len(a), &a[0])
	fmt.Printf("b: %+v,cap(b):%d,len(b):%d arrd:%v \n", b, cap(b), len(b), &b[0])
	fmt.Printf("IssueDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------IssueDemo end----------->>")
}

func MultiSliceDemo() {
	println("//<<-------------------------MultiSliceDemo start-----------")
	start := time.Now()
	as := make([][]int, 0, 3)
	a := []int{1, 1, 1, 1}
	b := []int{2, 2, 2}
	c := []int{3, 3, 3}
	as = append(as, a)
	as = append(as, b)
	as = append(as, c)
	log.Printf("len(as): %+v\n", len(as))
	log.Printf("as[0]: %+v\n", as[0])
	log.Printf("as[1]: %+v\n", as[1])
	fmt.Printf("MultiSliceDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------MultiSliceDemo end----------->>")
}

func SearchIssueDemo() {
	// var htmlNospaceReplacementTable = []string{
	// 	0:    "&#xfffd;",
	// 	'\t': "&#9;",
	// 	'\n': "&#10;",
	// 	'\v': "&#11;",
	// 	'\f': "&#12;",
	// 	'\r': "&#13;",
	// 	' ':  "&#32;",
	// 	'"':  "&#34;",
	// 	'&':  "&amp;",
	// 	'\'': "&#39;",
	// 	'+':  "&#43;",
	// 	'<':  "&lt;",
	// 	'=':  "&#61;",
	// 	'>':  "&gt;",
	// 	// A parse error in the attribute value (unquoted) and
	// 	// before attribute value states.
	// 	// Treated as a quoting character by IE.
	// 	'`': "&#96;",
	// }
	// 	'\'': "&#39;",
	println("//<<-------------------------SearchIssueDemo start-----------")
	start := time.Now()
	const sentence = `<a href="javascript:void(0)" onclick="queryXian(&#39;1100&#39;)">北京市</a>`
	singleQuoteConvert := []string{"&#39;", "'"}
	r := strings.NewReplacer(singleQuoteConvert...).Replace(sentence)
	log.Printf("r: %+v\n", r)
	// decodeStr := url.QueryEscape(sentence)
	// decodeStr, err := url.QueryUnescape(sentence)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// log.Printf("sentence: %+v\n", sentence)
	// log.Printf("decodeStr: %+v\n", decodeStr)
	return
	// sentence := string(`onClick="queryXian('4502')">`)
	for i := 0; i < len(sentence); i++ {
		log.Printf("i%d: %s\n", i, string(sentence[i]))
	}
	index := strings.Index(sentence, `'`) + 1
	// sufIndex := index + strings.Index(sentence[index+1:], `'`)
	sufIndex := index + strings.Index(sentence[index:], `'`)
	log.Printf("index: %+v\n", index)
	log.Printf("sufIndex: %+v\n", sufIndex)
	log.Printf("sentence[index:sufIndex]: %+v\n", sentence[index:sufIndex])

	fmt.Printf("SearchIssueDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------SearchIssueDemo end----------->>")
}

func pageDemo() {
	println("//<<-------------------------pageDemo start-----------")
	start := time.Now()
	ss := `查询结果共：16条   当前为：1页  共2页`
	const (
		current = `前为：`
		total   = `共`
		end     = `页`
	)
	a := strings.Index(ss, current) + len(current)
	b := strings.Index(ss[a:], end) + a
	log.Printf("ss[a:b]: %+v\n", ss[a:b])

	a = b + strings.Index(ss[b:], total) + len(total)
	b = strings.Index(ss[a:], end) + a
	log.Printf("total ss[a:b]: %+v\n", ss[a:b])
	fmt.Printf("pageDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------pageDemo end----------->>")
}

func ConvertDemo() {
	println("//<<-------------------------ConvertDemo start-----------")
	start := time.Now()
	a := "f333efjXjfej"
	r := strings.ToUpper(a)
	log.Printf("r: %+v\n", r)
	fmt.Printf("ConvertDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------ConvertDemo end----------->>")
}

func IndexDemo() {
	println("//<<-------------------------IndexDemo start-----------")
	start := time.Now()
	// line := strings.TrimSpace("one line")
	// log.Printf("line: %+v\n", line)
	// return
	a := "play_duration"

	if (strings.Index(a, "cost") >= 0) || (strings.Index(a, "ratio") >= 0) || (strings.Index(a, "cpm") >= 0) {
		println("waht?")
	}
	r := strings.Index(a, "d")
	// rNil := strings.Index(a, "d")
	log.Printf(": %+v\n", r)

	ns := "曹渡;林颖;董越;梁丽丽;劳咏姗;罗颂雅;陈驹峥;李伟明;莫冬倩;林子良;谭荣棉;何建波;陈璐;刘日荣;唐比昌;林恒;徐彬;郑健豪;刘舒婷;钟思远;李华煜;刘伟;梁杰;苏创绩;周开伟;黄佳豪;韦万;何世坛;吴甲林;周凯帆;谢天;程辉;刘东方;许忠洲;李文莎;柏石先;林碧洪;张月丽;姚梓洋;廖梦思;郑园柳;张三三;姚欢;黄晓冬;何嘉静;文伟英;谢淑卿;罗梦飞;陈耿涛;陈伟平;黄腾飞;苏特;张婧;黄智杰;"
	ss := strings.Split(ns, ";")
	log.Printf("len(ss): %+v\n", len(ss))
	log.Printf("s[0]: %+v\n", ss[0])
	log.Printf("ss[len(ss)-1]: %+v\n", ss[len(ss)-1])
	log.Printf("ss[len(ss)-2]: %+v\n", ss[len(ss)-2])
	// log.Printf("rNil: %+v\n", rNil)
	// a := "/1/3/2_8.png"
	// log.Printf(": %+v\n", a[strings.LastIndex(a, "/"):])
	fmt.Printf("IndexDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------IndexDemo end----------->>")
}

func GotoDemo() {
	println("//<<-------------------------GotoDemo start-----------")
	start := time.Now()
	c := 0
retry:
	c++
	if c < 4 {
		log.Println("retry: works")
		goto retry
	}
	log.Printf("c: %+v\n", c)
	log.Println("good: works")
	fmt.Printf("GotoDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------GotoDemo end----------->>")
}

func MaxElementDemo(s []int) int {
	println("//<<-------------------------MaxElementDemo start-----------")
	start := time.Now()
	sLen := len(s)
	i := 0
	last := s[sLen-1]
	for i < sLen-1 {
		// max := s[i]
		// s[sLen-1] = max
		// i++
		// for s[i] < max { //tips: max will never be greater than itself,so i won't be greater thatn max(s[sLen-1])
		s[sLen-1] = s[i]
		i++
		for s[i] < s[sLen-1] { //tips: max will never be greater than itself,so i won't be greater thatn max(s[sLen-1])
			i++
		}
	}
	fmt.Printf("MaxElementDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------MaxElementDemo end----------->>")
	if last > s[sLen-1] {
		return last
	}
	return s[sLen-1]
}

func HashDemo() {
	println("//<<-------------------------HashDemo start-----------")
	start := time.Now()
	const (
		base     = 6
		maxIndex = 101
	)
	//you might use linkList as under data storage as well
	underArray := make([]*string, maxIndex)

	getIndex := func(str string) int {
		hash := 0
		for i := 0; i < len(str); i++ {
			hash = hash*base + i*int(str[i])
		}
		r := hash % maxIndex
		log.Printf("index for %v r: %+v\n", str, r)
		return r
	}

	sethash := func(key string, value string) {
		underArray[getIndex(key)] = &value
	}

	gethash := func(key string) string {
		return *underArray[getIndex(key)]
	}

	sethash("first", "JialinWu")
	r := gethash("first")
	log.Printf("r: %+v\n", r)

	sethash("frist", "AGui")
	r = gethash("frist")
	log.Printf("r: %+v\n", r)

	sethash("frits", "sAGui")
	r = gethash("frits")
	log.Printf("r: %+v\n", r)

	fmt.Printf("HashDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------HashDemo end----------->>")
}

func TrimDemo() {
	println("//<<-------------------------TrimDemo start-----------")
	start := time.Now()
	o := "a 3j 45 "
	ro := strings.TrimSpace(o)
	logx.Debug("ro: %+v\n", ro)
	fmt.Printf("TrimDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------TrimDemo end----------->>")
}

func StrConvertDemo(i int) string {
	return strconv.Itoa(i) + strconv.Itoa(i*10)
}

func StrConvertFmtDemo(i int) string {
	return fmt.Sprintf("%d%d", i, i*10)
}

var benStr []string

func init() {
	benStr = make([]string, 0, 100)
	for i := 10; i < 110; i++ {
		b := make([]byte, i)
		rand.Read(b)
		benStr = append(benStr, string(b))
	}
}

func ByCombine(str []string) string {
	var r string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		r += str[i] //allocate memory when need, inefficiency
	}
	return r
}

func Byappend(str []string) string {
	totalLen := 0
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		totalLen += len(str[i])
	}
	bs := make([]byte, 0, totalLen) //allocate all memory at one time
	for i := 0; i < strLen; i++ {
		bs = append(bs, str[i]...)
	}

	return string(bs)
}
