// Package main provides ...
package newDir

import (
	"fmt"
	"testing"
)

func TestJustDemo(t *testing.T) {
	JustDemo()
	str := "jialin"
	encodeBs := base64Encode([]byte(str))
	fmt.Printf("bs: %+v\n", encodeBs)
	decodeStr, _ := base64Decode(encodeBs)
	fmt.Printf("decodeStr: %+v\n", string(decodeStr))

}
