package utils

import (
	"io"
	"testing"
)

func TestError(t *testing.T) {
	vars := []struct {
		in   error
		want string
	}{
		{
			WrapError(nil, "a"),
			"a",
		},
		{
			WrapError(io.EOF, "a"),
			"EOF->a",
		},
	}

	for _, v := range vars {
		if v.in.Error() != v.want {
			t.Errorf("%v expects %s but %s given", v, v.want, v.in.Error())
		}
	}
}
