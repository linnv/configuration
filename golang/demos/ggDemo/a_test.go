package newDir

import "testing"

func TestC(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := C(); got != tt.want {
				t.Errorf("C() = %v, want %v", got, tt.want)
			}
		})
	}
}
