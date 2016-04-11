package main

import (
	. "demo/model/mix"
	"fmt"
)

func main() {
	InitMemory()
	Add("jialin")
	Add("wujialin")
	Add("jialinwu")
	// fmt.Printf("name %s by id %d\n", GetNameById(10000), 10000)
	fmt.Printf("name %s by id %d\n", GetNameById(10), 10)
	for i := 0; i < 20; i++ {
		Add("a" + string(i))
	}
	fmt.Printf("name %s by id %d\n", GetNameById(2), 2)
	PrintAll()
}
