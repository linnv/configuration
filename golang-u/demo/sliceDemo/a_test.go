package demo

import "testing"

func TestJustDemo(t *testing.T) {
	TrimDemo()
	// HashDemo()

	// GotoDemo()
	// IndexDemo()
	// MultiSliceDemo()
	// SearchIssueDemo()
	// pageDemo()
	// ConvertDemo()
	// IssueDemo()
	JustDemo()
	// DiffPosStrDemo()

	// SortInt64Demo()
	// s := "jjjj	xx|xxx yy\rrrr"
	// CharactersDemo(s)

	// EmptySliceDemo()
	// LoopDemo()
	// a := make([]int, 0, 2)
	// a := make([]int, 2)
	// for i := 0; i < 10; i++ {
	// 	a = append(a, i)
	// }
	// fmt.Printf("  a: %+v\n", a)
	// BufferDemo()
	// BitCalucate()
	// r := BinarySearch()
	// fmt.Printf("r: %+v\n", r)

	// r := LongConsecutiveSequence()
	// fmt.Printf("r: %+v\n", r)

	// SlicePointer()

	// fmt.Printf("r: %+v\n", r)

	// r := TwoSum()
	// fmt.Printf("r: %+v\n", r)

	// strFormat := "%sxxxxxoo"
	// // args := []string{"jialin"}
	// args := "jialin"
	// // strFunc(fmt.Sprintf, strFormat, args)
	// Llog("first", fmt.Sprint, args)
	// //issue %!(EXTRA string=jialin)
	// Llogf("first", strFormat, fmt.Sprintf, args)

	// tErr(nil, nil)

	// ReverseLinkList()

	// ar := RetSlice()
	// //@todoDelelte
	// fmt.Printf("ar: %+v,cap(ar):%d,len(ar):%d  \n", ar, cap(ar), len(ar))
	//@toDelete
	// fmt.Printf("ar[7]: %+v\n", ar[7])

	// SliceDemo()
	//<<-------------------------pointers slice and instance demo start-----------
	// a := make([]int, 0, 3)
	// for i := 0; i < 3; i++ {
	// 	a = append(a, 0)
	// 	// a[i] = 0
	// }
	//
	// fmt.Printf("  a: %+v\n", a)
	// Array(a)
	// fmt.Printf("  a: %+v\n", a)
	//
	// b := [3]int{0, 0, 0}
	// fmt.Printf("  b: %+v\n", b)
	// ArrayInstance(&b)
	// fmt.Printf("  b: %+v\n", b)

	// ArrayInstance(a)
	// fmt.Printf("  a: %+v\n", a)
	// str := "ffe..'fejf.jfe.png"
	// r := path.Ext(str)
	// r := GetSuffixTypeByDot(str)
	// fmt.Printf("  r: %+v\n", r)
	// fmt.Printf("  str: %+v\n", str)
	// fmt.Printf("  len(str): %+v\n", len(str))

	// ReplaceStr()
	// SliceDemo()
	// Trap()
	// aps := []*A{
	// 	&A{Name: "j"},
	// 	&A{Name: "i"},
	// }
	// println("------aps start\n========================\n")
	// for k, v := range aps {
	// 	fmt.Printf("%+v: %+v\n", k, v)
	// }
	// PointersSliceDemo(aps)
	// println("------aps end\n========================\n")
	// for k, v := range aps {
	// 	fmt.Printf("%+v: %+v\n", k, v)
	// }
	//
	// as := []A{
	// 	A{Name: "j"},
	// 	A{Name: "i"},
	// }
	// println("------as start\n========================\n")
	// for k, v := range as {
	// 	fmt.Printf("%+v: %+v\n", k, v)
	// }
	// instanceSliceDemo(as)
	// println("------as end\n========================\n")
	// for k, v := range as {
	// 	fmt.Printf("%+v: %+v\n", k, v)
	// }
	//---------------------------pointers slice and instance demo end----------->>

	// a := []int{1, 2, 3, 4, 5}
	// b := []int{2, 3}
	// fmt.Printf("  a1: %+v\n", a)
	// c := IntArraySubtract(a, b)
	// a = IntArraySubtract(a, b)
	// fmt.Printf("  c: %+v\n", c)
	// fmt.Printf("  a2: %+v\n", a)

	//peek get the last element but no pop it
	// sl := []int{1, 2, 3333, 3, 4, 566}
	// fmt.Printf("sl[:1]: %+v\n", sl[:1])
	// fmt.Printf("sl[:0]: %+v\n", sl[:0])
	// fmt.Printf("sl[1]: %+v\n", sl[1])
	// fmt.Printf("sl[2:4]: %+v\n", sl[2:4])

}

//  go test -run=CountOneDemo
func TestCountOneDemo(t *testing.T) {
	// CountOneDemo()
}

func TestMaxElementDemo(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{[]int{3, 1, 4, 2, 5, 3, 2}}, 5},
		{"normal", args{[]int{3, 1, 4, 2, 5, 3, 9}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxElementDemo(tt.args.s); got != tt.want {
				t.Errorf("MaxElementDemo() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_kpHashDemo(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			kpHashDemo()
// 		})
// 	}
// }

func BenchmarkStrConvertDemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrConvertDemo(i)
	}
}

func BenchmarkStrConvertFmtDemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrConvertFmtDemo(i)
	}
}

func BenchmarkBybytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := ByCombine(benStr)
		if r == "" {

		}
	}
}

func BenchmarkByappend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := Byappend(benStr)
		if r == "" {

		}
	}
}
