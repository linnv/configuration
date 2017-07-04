// Package  provides ...
package mix

import (
	"fmt"
)

var MemoryData *MemoryPool

type MemoryPool struct {
	count   int
	idArray []int
	mmap    map[int]interface{}
	// mmap    map[int]string
}

func InitMemory() {
	MemoryData = new(MemoryPool)
	MemoryData.mmap = make(map[int]interface{}, 0)
}

func AddWithId(id int, str string) {
	MemoryData.count++
	MemoryData.idArray = append(MemoryData.idArray, id)
	MemoryData.mmap[id] = str
}

func Add(str string) {
	if _, err := MemoryData.mmap[MemoryData.count]; err {
		// MemoryData.count--
		return
	}
	MemoryData.count++
	MemoryData.idArray = append(MemoryData.idArray, MemoryData.count)
	MemoryData.mmap[MemoryData.count] = str
}

func PrintAll() {
	println("count:", MemoryData.count)
	// println("len array:", len(MemoryData.idArray))
	fmt.Printf("array:%+v\n", MemoryData.idArray)
	// for k, v := range MemoryData.mmap {
	// 	fmt.Printf("%d->%s\n", k, v)
	// }
}

func GetNameById(id int) interface{} {
	if v, err := MemoryData.mmap[id]; err {
		return v
	}
	AddWithId(id, "new name from db")
	return MemoryData.mmap[id]
}
