// Package main provides ...
package demo

import "testing"

const protype = "jialinwu"

func BenchmarkFmtEncode(b *testing.B) {
	bs := getSum(protype)
	for i := 0; i < b.N; i++ {
		FmtEncode(bs)
	}
}

func BenchmarkHexEncode(b *testing.B) {
	bs := getSum(protype)
	for i := 0; i < b.N; i++ {
		HexEncode(bs)
	}
}

// func TestJustDemo(t *testing.T) {
// 	JustDemo()
// 	// page := "100"
// 	var j int
// 	for i := 0; i < 10; i++ {
// 		j = i * 11
// 		fmt.Printf("MD5Demo(): %s\n", MD5EncodeDemo(strconv.Itoa(j)))
// 	}
// 	_ = MD5DecodeDemo("jialin")
// }
