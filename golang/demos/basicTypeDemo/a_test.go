// Package main provides ...
package newDir

import (
	"log"
	"testing"
	"unsafe"
)

func TestJustDemo(t *testing.T) {
	JustDemo()
	a := int64(1)
	b := int(1)
	c := "ss"
	const d = "ss"
	log.Printf("unsafe.Sizeof(int64): %+v\n", unsafe.Sizeof(a))
	log.Printf("unsafe.Sizeof(int32): %+v\n", unsafe.Sizeof(b))
	log.Printf("unsafe.Sizeof(c): %+v\n", unsafe.Sizeof(c))
	log.Printf("unsafe.Sizeof(d): %+v\n", unsafe.Sizeof(d))

}
