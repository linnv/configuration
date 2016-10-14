// Package main provides ...
package newDir

import (
	"fmt"
	"strconv"
	"testing"
)

func TestJustDemo(t *testing.T) {
	JustDemo()
	// syscall()
	a := 22
	// s := uitoa(uint(a))
	s := strconv.Itoa(a)
	fmt.Printf("s: %T\n", s)
}
