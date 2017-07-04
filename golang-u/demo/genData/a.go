// Package main provides ...
package demo

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}
func JustDemo() {
	println("<<<JustDemo start---------------------------")
	GenerateDataDemo()
	println("-----------------------------JustDemo end>>>")
	return
}

type CC struct {
	Index []int `json:"num"`
}

const randomCount = 1000000
const dataRange = 100000
const filepath = "./ints.log"

func GenerateDataDemo() {
	println("//<<-------------------------GenerateDataDemo start-----------")
	start := time.Now()
	c := &CC{}
	// tmp := make([]int, 0, randomCount)
	c.Index = make([]int, 0, randomCount)

	for i := 0; i < randomCount; i++ {
		// tmp = append(tmp, random.Intn(dataRange))
		c.Index = append(c.Index, random.Intn(dataRange))
	}
	// copy(c.Index, tmp)
	c.Index = append(c.Index, c.Index...)
	bs, err := json.Marshal(c)
	CheckError(err)

	err = SaveBytesToFile(bs, filepath)
	CheckError(err)

	fmt.Printf(" %v\n", time.Since(start))
	println("//---------------------------GenerateDataDemo end----------->>")
}

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
		return
	}
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
