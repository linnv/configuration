package newDir

import (
	"fmt"
	"runtime/debug"
	"sunteng/commons/log"
	"unsafe"
)

func pointerIntDemo(a *int) {
	println("<<<pointerIntDemo start---------------------------")
	log.Logf("a: %+v\n", a)
	*a = 10
	println("-----------------------------pointerIntDemo end>>>")
	return
}

func ReturnPointerDemo() *int {
	println("<<<ReturnPointerDemo start---------------------------")
	var a *int = new(int)
	log.Logf("a: %+v\n", a)
	*a = 11
	log.Logf("a: %+v\n", a)
	println("-----------------------------ReturnPointerDemo end>>>")
	return a
}
func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

type AA struct {
	Name string `json:"Name"`
}

func byval(q *int) {
	debug.PrintStack()
	fmt.Printf("3. byval -- q %T: &q=%p q=&i=%p  *q=i=%v\n", q, &q, q, *q)
	*q = 4143
	fmt.Printf("4. byval -- q %T: &q=%p q=&i=%p  *q=i=%v\n", q, &q, q, *q)
	/*
		In function main, byval(p) is a function call which assigns the value (p=&i) 0xf800000040 of the argument at memory location (&p) 0xf8000000f0 to the function byval parameter q at memory location (&q) 0xf8000000d8. In other words, memory is allocated for the byval parameter q and the value of the main byval argument p is assigned to it; the values of p and q are initially the same, but the variables p and q are distinct.

		In function byval, using pointer q (*int), which is a copy of pointer p (*int), integer *q (i) is set to a new int value 4143. At the end before returning. the pointer q is set to nil (zero value), which has no effect on p since q is a copy.
	*/
	q = nil
	// fmt.Printf("nil. byval -- q %T: &q=%p q=&i=%p  *q=i=%v\n", q, &q, q, *q)
}

func PointerDetailDemo() {
	// a := 4
	// fmt.Printf("  value of a: %x\n", a)
	// aNewAddr := unsafe.Pointer(&a)
	// fmt.Printf("  address of aNew: %x\n", aNewAddr)

	// i := int(42)
	// fmt.Printf("1. main  -- i  %T: &i=%p i=%v\n", i, &i, i)
	// p := &i
	// fmt.Printf("2. main  -- p %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p)
	// byval(p)
	// fmt.Printf("5. main  -- p %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p)
	// fmt.Printf("6. main  -- i  %T: &i=%p i=%v\n", i, &i, i)

	// an := new(int)
	// p := unsafe.Pointer(&an)
	// i := 19
	// an = &i
	// tmpP := &i
	// *an = 10
	// *tmpP = 1199
	// fmt.Printf("  *an: %v\n", *an)
	// fmt.Printf("  i: %v\n", i)
	// fmt.Printf("  p: %p\n", p)
	//
	// ip := unsafe.Pointer(&i)
	// fmt.Printf("  an: %v\n", an)
	// fmt.Printf("  ip: %v\n", ip)
	// //@toDelete
	// fmt.Printf("  tmpP: %v\n", tmpP)

	target := new(int)
	fmt.Printf("*target: %+v\n", *target)
	fmt.Printf("value of target(self value): %+v\n", target)
	pointerToTarget := unsafe.Pointer(&target)
	fmt.Printf("pointerToTarget: %+v\n", pointerToTarget)
	var value int = 19
	target = &value
	fmt.Printf("value: %+v\n", value)
	fmt.Printf("*target: %+v\n", *target)
	fmt.Printf("value of target(self value): %+v\n", target)
	addressOfValue := unsafe.Pointer(&value)
	fmt.Printf("addressOfValue: %+v\n", addressOfValue)

	targetc := new(int)
	targetc = target

	fmt.Printf("targetc: %+v\n", targetc)
	fmt.Printf("*targetc: %+v\n", *targetc)

	// println("//<<-------------------------PointerDetailDemo start-----------")
	// aNew := new(AA)
	// fmt.Printf("  value of aNew: %+v\n", aNew)
	// aNewAddr := unsafe.Pointer(&aNew)
	// fmt.Printf("  address of aNew: %+v\n", aNewAddr)
	//
	// a := AA{Name: "origin"}
	// aNew = &a
	// fmt.Printf("  value of aNew: %v\n", aNew)
	// aAddr := unsafe.Pointer(&a)
	// fmt.Printf("  value of a: %+v\n", a)
	// aNew.Name = "change by new pointer"
	// fmt.Printf("  value of a: %+v\n", a)
	// fmt.Printf("  address of a: %+v\n", aAddr)
	//
	// ap := &AA{Name: "origin"}
	// aNew = ap
	// fmt.Printf("  value of aNew: %v\n", aNew)
	// fmt.Printf("  value of *aNew: %v\n", *aNew)
	// apAddr := unsafe.Pointer(&ap)
	// fmt.Printf("  value of ap: %+v\n", ap)
	// fmt.Printf("  address of ap: %+v\n", apAddr)
	// fmt.Printf("  address of aNew: %+v\n", aNewAddr)
	// println("//---------------------------PointerDetailDemo end----------->>")
}

func changeValueDemo(a AA) {
	println("//<<-------------------------changeValueDemo start-----------")
	p := unsafe.Pointer(&a)
	//@toDelete
	fmt.Printf("func  address of a : %+v\n", p)
	a.Name = "change through value"
	println("//---------------------------changeValueDemo end----------->>")
}

func changePointerDemo(a *AA) {
	println("//<<-------------------------changeValueDemo start-----------")
	p := unsafe.Pointer(&a)
	//@toDelete
	fmt.Printf(" func  address of &a : %+v\n", p)
	pp := unsafe.Pointer(a)
	fmt.Printf(" func  address of a : %+v\n", pp)
	a.Name = "change through pointer"
	println("//---------------------------changeValueDemo end----------->>")
}
