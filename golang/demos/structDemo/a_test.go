package newDir

import "testing"

// func TestJustDemo(t *testing.T) {
// 	// JustDemo()
// 	// AListDemo()
// 	a := Base{Count: 1}
// 	a2 := &Base{Count: 2}
// 	b := Derive{}
// 	b.Base = a
// 	b.Two = a2
// 	b.UpdateCount(19)
// 	b.All()
// 	b.UpdateCountPointerReceiver(19)
// 	b.All()
//
// 	fmt.Printf("value member: works\n")
// 	// a.UpdateCountPointerReceiver(20)
// 	a.All()
// 	a.UpdateCountPointerReceiver(20)
// 	b.UpdateCount(222)
// 	a.All()
//
// 	fmt.Printf("pointer member: works\n")
// 	b.All()
// 	a2.All()
// 	a2.UpdateCountPointerReceiver(20)
// 	a2.UpdateCount(222)
// 	a2.All()
// 	b.All()
// }

func TestJustDemo(t *testing.T) {
	// p := &User{"Damon", "damon@xxoo.com"}
	// p.Notify()
	// log.Println(p.Name)
	//
	// p.PointerNotify()
	// log.Println(p.Name)
	//
	// println("-------value ----")
	// u := User{"Damon", "damon@xxoo.com"}
	// u.Notify()
	// log.Println(u.Name)
	//
	// u.PointerNotify()
	// log.Println(u.Name)
	// d := &Derive{Two: &Base{Count: 10}}
	// v, ok := d.(*Derive)
	// fmt.Printf(": works\n")
	// JustDemo()
}

func TestAssigneStructDemo(t *testing.T) {
	AssigneStructDemo()
}
