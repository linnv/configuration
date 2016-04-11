package mongoDemo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStructTagDemo(t *testing.T) {
	Convey("elegant result", t, func() {
		// var tmp interface{}
		// err := StructTagDemo()
		// So(err, ShouldEqual, nil)
		// err := OrOperationDemo()
		// err := RegularOperationDemo()
		// So(err, ShouldEqual, nil)
		// b := A{N: good"}
		// r := `{"name":"good"}`
		// ret := MarshalDemo(b)
		// // fmt.Printf("rest: %+v\n", ret)
		// // t.Log(ret)
		// So(r, ShouldNotEqual, ret)
		// So(r, ShouldEqual, ret)
		// SuntengFindOperationDemo()
		// SuntengUpdateOperationDemo()
		AggregateQueryDemo()
	})
}
