// Package main provides ...
package newDir

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func MD5EncodeDemo(protype string) string {
	println("<<<MD5Demo---------------------------")

	// name := "win27v@gmail.com"
	// name := "win27v@gmail.com"
	h := md5.New()
	h.Write([]byte(protype))
	println("-----------------------------MD5Demo>>>")
	return hex.EncodeToString(h.Sum(nil))
}

func MD5DecodeDemo(md5Str string) string {
	println("<<<MD5Demo---------------------------")

	h := md5.New()
	// bs, err := hex.DecodeString(md5Str)
	fmt.Printf("md5Str: %+v\n", md5Str)
	h.Write([]byte(md5Str))
	hrefBs := h.Sum(nil)

	fmt.Printf("gen bs: %+v\n", hrefBs)
	href := hex.EncodeToString(hrefBs)

	fmt.Printf("encode to str: %+v\n", href)
	bs, err := hex.DecodeString(href)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("decode bs: %+v\n", bs)
	println("-----------------------------MD5Demo>>>")
	return hex.EncodeToString(h.Sum(nil))
}

// func Base64UrlSafeDecodersrc string) (dest []byte, err error) {
// 	src = strings.Replace(src, "-", "+", -1)
// 	src = strings.Replace(src, "_", "/", -1)
// 	/* return base64.URLEncoding.DecodeString(src) */
// 	return Base64UrlDecoder(src)
// }

func JustDemo() {
	println("<<<JustDemo---------------------------")
	a := 19
	fmt.Printf("a: %+v\n", a)
	// s := MD5EncodeDemo("wujialin")
	//vpn password
	// fmt.Printf("s: %+v\n", s[:10])
	MD5DecodeDemo("wujialin")
	println("-----------------------------JustDemo>>>")
	return
}
