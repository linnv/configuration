package demo

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/linnv/logx"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	a := int(123)
	b := int64(123)
	c := "foo"
	d := struct {
		FieldA float32
		FieldB string
	}{0, "bar"}

	fmt.Printf("a: %T, %d Bytes\n", a, unsafe.Sizeof(a))
	fmt.Printf("b: %T, %d Bytes\n", b, unsafe.Sizeof(b))
	fmt.Printf("c: %T, %d Bytes\n", c, unsafe.Sizeof(c))
	fmt.Printf("d: %T, %d Bytes\n", d, unsafe.Sizeof(d))

	fmt.Printf("a: %T, %d Bytes\n", a, reflect.TypeOf(a).Size())
	fmt.Printf("b: %T, %d Bytes\n", b, reflect.TypeOf(b).Size())
	fmt.Printf("c: %T, %d Bytes\n", c, reflect.TypeOf(c).Size())
	fmt.Printf("d: %T, %d Bytes\n", d, reflect.TypeOf(d).Size())

	const lenbs = 1190
	bs := make([]byte, lenbs)
	bs[0] = 'b'
	for i := 1; i < lenbs-1; i++ {
		bs[i] = 'a'
	}
	bs[lenbs-1] = 'e'
	bsStr := string(bs)
	fmt.Printf("bsStr: %T, %d Bytes\n", bsStr, reflect.TypeOf(bsStr).Size())
	logx.Debug("bsStr[0]: %+v\n", bsStr[0])
	logx.Debug("bsStr[lenbs-1]: %+v\n", bsStr[lenbs-1])
	println("-----------------------------JustDemo end>>>")
	return
}
