// Package main provides ...
package demo

import "testing"

func TestJustDemo(t *testing.T) {
	JustDemo()
}

// func TestStructTagJsonUnmarshalDemo(t *testing.T) {
// 	Convey("json unmarshal result", t, func() {
// 		b := &JPerson{"jialin", "p", 101}
// 		// b := &BPerson{"jialin", "p", 100}
// 		r, err := StructTagJsonDemo(b)
// 		So(err, ShouldEqual, nil)
// 		t.Log("json ret:", string(r))
// 		ub, err := StructTagJsonUnMarshalDemo(r)
// 		So(err, ShouldEqual, nil)
// 		t.Logf("uj: %+v\n", ub)
// 		// So(err, ShouldEqual, nil)
// 		So(err, ShouldNotEqual, nil)
// 	})
// }
//
// func TestStructTagBsonUnmarshalDemo(t *testing.T) {
// 	Convey("bson unmarshal result", t, func() {
// 		b := &BPerson{"jialin", "p", 101}
// 		// b := &BPerson{"jialin", "p", 100}
// 		// var tmp interface{}
// 		r, err := StructTagBsonDemo(b)
// 		So(err, ShouldEqual, nil)
// 		ub, err := StructTagBsonUnMarshalDemo(r)
// 		So(err, ShouldEqual, nil)
// 		t.Logf("%v", ub)
// 		// So(err, ShouldEqual, nil)
// 		So(err, ShouldNotEqual, nil)
// 	})
// }
//
// func TestStructTagBsonDemo(t *testing.T) {
// 	Convey("bson result", t, func() {
// 		b := &BPerson{"jialin", "p", 101}
// 		// b := &BPerson{"jialin", "p", 100}
// 		r, err := StructTagBsonDemo(b)
// 		t.Log(r)
// 		t.Logf("ret:%v", string(r))
// 		So(err, ShouldNotEqual, nil)
// 		// rr := `{"Name":"jialin","Phone":"p","Id":1}`
// 		// So(r, ShouldEqual, rr)
// 	})
// }
//
// func TestStructTagJsonDemo(t *testing.T) {
// 	Convey("json result", t, func() {
// 		b := &JPerson{"jialin", "p", 1}
// 		r, err := StructTagJsonDemo(b)
// 		t.Log(r)
// 		t.Logf("ret:%v", string(r))
// 		So(err, ShouldNotEqual, nil)
// 		// rr := `{"Name":"jialin","Phone":"p","Id":21}`
// 		// So(string(r), ShouldEqual, rr)
// 	})
// }
