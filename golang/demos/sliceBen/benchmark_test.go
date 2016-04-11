package ben

import (
	"demo/demos/utility"
	"fmt"
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const filePath = "/Users/Jialin/Desktop/ints.log"

type Num struct {
	Index []int `json:"num"`
}

func BenchmarkExampleDemo(b *testing.B) {
	//@TODO invoke this benchmark func by executing  go test -bench=Demo     xxx
	N := &Num{}
	//notice   结构的成员变量必须为导出，否则unmarsh不到数据
	utility.JsonFromFileDemo(filePath, N)
	//@toDelete
	fmt.Printf("N.Index): %+v\n", len(N.Index))

	// sort.IntSlice(N.Index).Sort()
	for i := 0; i < b.N; i++ {
		SortDeleteDuplicate(N.Index)
	}
}

func BenchmarkUniqueIntArray(b *testing.B) {
	//@TODO invoke this benchmark func by executing  go test -bench=Demo     xxx

	// var src = []int{1, 1, 2, 3, 2, 3, 4, 3, 4, 5, 2, 1}

	N := &Num{}
	utility.JsonFromFileDemo(filePath, N)
	fmt.Printf("N.Index): %+v\n", len(N.Index))
	for i := 0; i < b.N; i++ {
		UniqueIntArray(N.Index)
	}
}

func TestSortDeleteDuplicate(t *testing.T) {
	var src = []int{1, 1, 2, 3, 2, 3, 4, 3, 4, 5, 2, 1}
	sort.IntSlice(src).Sort()
	Convey("Slice should be uniq", t, func() {
		var res = SortDeleteDuplicate(src)
		So(res[0], ShouldEqual, 1)
		So(res[4], ShouldEqual, 5)

	})

}
