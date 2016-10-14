package newDir

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

type A struct {
	Name    string `json:"name"`
	Flatten string
}

func (a *A) afunc() {
	fmt.Println(" struct a func")
	// fmt.Printf("a.Name: %+v\n", a.Name)
}

// func (a A) Update() {
// 	a.Flatten = "update by instance"
// }
func (a *A) Update() {
	a.Flatten = "update by instance"
}

func (a *A) UpdatePointer() {
	a.Flatten = "update by refference"
}

type Derive struct {
	Name string `json:"Name"`
	Base
	Two *Base
}

func (this Derive) All() {
	fmt.Printf("this: %+v\n", this)
	fmt.Printf("this.Two: %+v\n", this.Two)
}

func (this Base) All() {
	fmt.Printf("this: %+v\n", this)
}

type Base struct {
	Count int `json:"Count"`
}

func (this Base) UpdateCount(n int) {
	this.Count = n
}

func (this *Base) UpdateCountPointerReceiver(n int) {
	this.Count = n
}

type AList []*A

func AListDemo() {
	println("<<<AListDemo start---------------------------")
	al := make(AList, 10)
	for i := 0; i < 10; i++ {
		a := &A{Name: strconv.Itoa(i)}
		al = append(al, a)
	}
	b := &struct {
		AList
		Sky string `json:"Sky "`
	}{
		al,
		"jialin",
	}
	fmt.Printf("b: %+v\n", b)
	mb, _ := json.Marshal(b)
	fmt.Printf("mb: %+v\n", string(mb))
	println("-----------------------------AListDemo end>>>")
	return
}

type User struct {
	Name  string
	Email string
}

func (u User) Notify() error {
	println("inner")
	println(u.Name)
	u.Name = "alimon"
	return nil
}

func (u *User) PointerNotify() error {
	println("inner pointer")
	println(u.Name)
	u.Name = "pointer alimon"
	return nil
}

type T struct {
	Name string `json:"Name"`
}

func (t T) ShowMember() {
	println(t.Name)
}

func (t *T) ShowPoniterMember() {
	println(t.Name)
}

type TT struct {
	AA string `json:"A"`
	B  string `json:"B"`
	*A
	Strs []string `json:"s"`
}
type smapList map[int]*A

func (sl smapList) Get() {
	for k, v := range sl {
		fmt.Printf("k: %+v ==v: %v\n", k, v)
	}
}

func (sl smapList) Update() {
	for k, _ := range sl {
		sl[k].Name = "uuuuu"
	}
}

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	// mm := make(map[int]*A)
	// for i := 0; i < 10; i++ {
	// 	mm[i] = &A{Name: "xxxx"}
	// }
	// smapList(mm).Get()
	// os.Stdout.Write(append([]byte("line"), '\n'))
	// smapList(mm).Update()
	// smapList(mm).Get()
	// os.Stdout.Write(append([]byte("line2"), '\n'))
	// for k, v := range mm {
	// 	fmt.Printf("%+v: %+v\n", k, v)
	// }

	a := TT{AA: "jialin", Strs: []string{"xxxx", "nnn"}}
	a.A = &A{Name: "jjj"}
	a.afunc()
	b := a.A
	b.Name = "343er3"
	fmt.Printf("a.Name: %+v\n", a.Name)
	fmt.Printf("a.A: %+v\n", a.A)
	// fmt.Printf("a: %+v\n", a)
	// ar, err := json.Marshal(a)
	// if err != nil {
	// 	return
	// }
	// fmt.Printf("  string(a): %+v\n", string(ar))
	// 同一个包内同访问内联结构的非导出方法以及成员变量的值
	//if uncomment line 15th,  runtime
	// "error: invalid memory address or nil pointer dereference "
	//will occur with no instance of *A initialize
	// a.afunc()
	// a.Afunc()
	// a.aa.Afunc()
	// t := &T{Name: "jjj"}
	// t.ShowPoniterMember()
	// t.ShowMember()

	// pointerInstance := &A{}
	// pointerInstance.Update()
	// fmt.Printf("pointerInstance.Flatten: %+v\n", pointerInstance.Flatten)
	// pointerInstance.UpdatePointer()
	// fmt.Printf("pointerInstance.Flatten: %+v\n", pointerInstance.Flatten)

	// pointerInstance := A{}
	// pointerInstance.Update()
	// fmt.Printf("pointerInstance.Flatten: %+v\n", pointerInstance.Flatten)
	// pointerInstance.UpdatePointer()
	// fmt.Printf("pointerInstance.Flatten: %+v\n", pointerInstance.Flatten)

	println("-----------------------------JustDemo end>>>")
	return
}

// type UkL struct {
// 	A ...bson.M
// 	// B ...[]int
// }

type stack []uintptr

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

type ErrerWrapper struct {
	e   error
	msg string
	s   *stack
}

// func (e *ErrerWrapper) Cause() error  { return e.e }
func (e *ErrerWrapper) Error() string { return e.msg }

func EWDemo(err error, m string) error {
	println("//<<-------------------------EWDemo start-----------")
	start := time.Now()
	fmt.Printf("EWDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------EWDemo end----------->>")
	return &ErrerWrapper{
		e:   err,
		msg: m,
		s:   callers(),
	}
}
