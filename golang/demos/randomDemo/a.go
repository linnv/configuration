// Package main provides ...
package newDir

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"time"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

const n = 10

type CC struct {
	Index []int `json:"num"`
}

const randomCount = 1000000
const randomGen = 100000
const filepath = "./ints.log"

func UniqueMap() {
	// randomCount := random.Intn(10000)
	start := time.Now()
	c := &CC{}
	tmp := make([]int, 0, randomCount)
	c.Index = make([]int, randomCount)
	err := JsonFromFileDemo(filepath, &c)
	CheckError(err)
	fmt.Printf("  c.Index): %+v\n", len(c.Index))

	fmt.Printf("  using map: \n")
	tmp = UniqueSliceMap(c.Index)
	fmt.Printf(" %v\n", time.Since(start))
	fmt.Printf("  len(tmp): %+v\n", len(tmp))
	fmt.Printf("  tmp[len(tmp)-1]: %+v\n", tmp[len(tmp)-1])
}
func UniqueAppend() {
	// randomCount := random.Intn(10000)
	start := time.Now()
	c := &CC{}
	tmp := make([]int, 0, randomCount)
	c.Index = make([]int, randomCount)
	err := JsonFromFileDemo(filepath, &c)
	CheckError(err)
	fmt.Printf("  c.Index): %+v\n", len(c.Index))

	fmt.Printf("  using append unique slice\n")
	for i := 0; i < len(c.Index); i++ {
		tmp = AppendIfMissing(tmp, c.Index[i])
	}

	fmt.Printf(" %v\n", time.Since(start))
	fmt.Printf("  len(tmp): %+v\n", len(tmp))
	fmt.Printf("  tmp[len(tmp)-1]: %+v\n", tmp[len(tmp)-1])
}

func SimpleUnique() {
	// randomCount := random.Intn(10000)
	start := time.Now()
	c := &CC{}
	tmp := make([]int, 0, randomCount)
	c.Index = make([]int, randomCount)
	err := JsonFromFileDemo(filepath, &c)
	CheckError(err)
	fmt.Printf("  c.Index): %+v\n", len(c.Index))

	fmt.Printf("  using simple slice unique: \n")
	tmp = UniqueSlice(c.Index)

	fmt.Printf(" %v\n", time.Since(start))
	fmt.Printf("  len(tmp): %+v\n", len(tmp))
}

func UniqueSlice(s []int) []int {
	//don't put len(s) out of for() like l:=len(s),or panic will occur
	Sort(s)
	for i := 0; i < len(s); i++ {
		// for j := len(s) - 1; j > i; j-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				s = append(s[:j], s[j+1:]...)
			}
		}
	}
	return s
}

func GenerateDataDemo() {
	println("//<<-------------------------GenerateDataDemo start-----------")
	start := time.Now()
	c := &CC{}
	tmp := make([]int, 0, randomCount)
	c.Index = make([]int, randomCount)

	for i := 0; i < randomCount; i++ {
		tmp = append(tmp, random.Intn(randomGen))
	}
	copy(c.Index, tmp)
	bs, err := json.Marshal(c)
	CheckError(err)

	err = SaveBytesToFile(bs, filepath)
	CheckError(err)

	fmt.Printf(" %v\n", time.Since(start))
	println("//---------------------------GenerateDataDemo end----------->>")
}

func SaveBytesToFile(bs []byte, filePath string) error {
	// f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err

	}
	defer f.Close()
	_, err = f.Write(bs)
	f.Sync()
	return nil
}

func JsonFromFileDemo(filePath string, c interface{}) (err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}

	var bs []byte

	bs, err = ioutil.ReadAll(f)
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, &c)
	return
}

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
		return
	}
}

func Sort(ids []int) []int {
	is := sort.IntSlice(ids)
	is.Sort()
	return is
}

func AppendIfMissing(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func UniqueSliceMap(a []int) []int {
	var res = make([]int, len(a))
	var mp = make(map[int]bool)
	var index = 0
	for _, i := range a {
		if !mp[i] {
			res[index] = i
			index++
		}
		mp[i] = true
	}
	return res[:index]
}
