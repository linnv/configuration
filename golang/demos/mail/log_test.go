package mail

import (
	"sunteng/commons/unittest"
	"testing"
)

func TestInit(t *testing.T) {
	convey.Convey("test mail log", t, func() {
		err := Init()
		if err != nil {
			panic(err.Error())
			return
		}
		convey.So(err, convey.ShouldEqual, nil)
	})
}
