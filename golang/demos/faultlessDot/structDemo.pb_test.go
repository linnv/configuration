// Package main provides ...
package faultlessDot

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// func TestMarshalDemo(t *testing.T) {
//
// 	b := A{N: "good"}
// 	r := `{"name":"good"}`
// 	ret := MarshalDemo(b)
// 	fmt.Printf("rest: %+v\n", ret)
// 	// t.Log(ret)
// 	if ret != r {
// 		t.Error("test foo:Addr failed")
// 	} else {
// 		t.Log("test foo:Addr pass")
// 	}
// }

func TestMarshalDemo(t *testing.T) {
	Convey("elegant result", t, func() {
		b := A{N: "good"}
		r := `{"name":"good"}`
		ret := MarshalDemo(b)
		// fmt.Printf("rest: %+v\n", ret)
		// t.Log(ret)
		So(r, ShouldNotEqual, ret)
		// So(r, ShouldEqual, ret)

	})
}
