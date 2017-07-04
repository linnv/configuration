// Package utilities provides ...
package utility

import "testing"

func TestCoerceInt64(t *testing.T) {
	str := []string{"1"}
	// str := []string{"1", "3", "7"}
	for i := 0; i < len(str); i++ {
		// r, err := CoerceInt64(reflect.ValueOf(str[i]).Interface())
		r, err := CoerceInt64(str[i])
		if r != 1 || err != nil {
			t.Error("test utility:CoerceInt64 failed")
		} else {
			t.Log("test utility:CoerceInt64  pass")
		}
	}
}
