package newDir

import (
	"fmt"
	"testing"
)

func TestJustDemo(t *testing.T) {
	// JustDemo()

	// r := isSubsequenceDemo()
	// r := decodeStringDemo()

	// s := "ababbc"
	// s := "weitong"
	// s := "aabcabb"
	// s := "aaabb"
	// s := "caaabb"
	s := "abcdedghijklmnopqrstuvwxyz"
	k := 1
	r := longestSubstringDemo(s, k)
	fmt.Printf("r: %+v\n", r)
}
