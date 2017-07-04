// Package main provides ...
package demo

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func getSum(protype string) []byte {
	h := md5.New()
	h.Write([]byte(protype))
	return h.Sum(nil)
}

//FmtEncode implements ...
func FmtEncode(bs []byte) {
	_ = fmt.Sprintf("%x", bs)
}

func HexEncode(bs []byte) {
	_ = hex.EncodeToString(bs)
}

func MD5EncodeDemo(protype string) string {
	println("<<<MD5Demo---------------------------")

	// name := "win27v@gmail.com"
	h := md5.New()
	h.Write([]byte(protype))
	fmt.Printf("%s\n")
	println("-----------------------------MD5Demo>>>")
	return hex.EncodeToString(h.Sum(nil))
}

// func MD5DecodeDemo(md5Str string) string {
// 	println("<<<MD5Demo---------------------------")
// 	//
// 	// h := md5.New()
// 	// // bs, err := hex.DecodeString(md5Str)
// 	// h.Write([]byte(md5Str))
// 	// hrefBs := h.Sum(nil)
// 	// fmt.Printf("hrefBs: %+v\n", hrefBs)
// 	// href := hex.EncodeToString(hrefBs)
// 	// fmt.Printf("href: %+v\n", href)
// 	bs, err := hex.DecodeString(md5Str)
// 	if err != nil {
// 		panic(err.Error())
// 		return ""
// 	}
// 	fmt.Printf("bs: %+v\n", bs)
// 	println("-----------------------------MD5Demo>>>")
// 	// return hex.EncodeToString(h.Sum(nil))
// }
//
// // func Base64UrlSafeDecodersrc string) (dest []byte, err error) {
// // 	src = strings.Replace(src, "-", "+", -1)
// // 	src = strings.Replace(src, "_", "/", -1)
// // 	/* return base64.URLEncoding.DecodeString(src) */
// // 	return Base64UrlDecoder(src)
// // }
//
// func JustDemo() {
// 	println("<<<JustDemo---------------------------")
// 	a := 19
// 	fmt.Printf("a: %+v\n", a)
// 	println("-----------------------------JustDemo>>>")
// 	return
// }
