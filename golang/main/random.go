package main

import (
	"fmt"
	"math/rand"
	"time"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

const n = 10

func main() {
	for i := 0; i < 9; i++ {
		randomRatio := random.Intn(10)
		fmt.Printf("randomRatio: %+v\n", randomRatio)
	}
	var i int
	i = 199
	if i > n {
		println("bigger")
	}

}
