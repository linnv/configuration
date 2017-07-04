package demo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJustDemo(t *testing.T) {
	Convey("JustDemo", t, func() {
		JustDemo()
		MarshalDemo()
		MarshalDemoP()

		MarshalDemoA()
		MarshalDemoAP()

		bs := []byte(`{"name":"a","bName":"bn"}`)
		err := UnMarshalDemoP(bs)
		So(err, ShouldEqual, nil)
		err = UnMarshalDemo(bs)
		So(err, ShouldEqual, nil)

		err = UnMarshalDemoAP(bs)
		So(err, ShouldEqual, nil)
		err = UnMarshalDemoA(bs)
		So(err, ShouldEqual, nil)
	})
}
