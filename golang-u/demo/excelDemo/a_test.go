package demo

import (
	"reflect"
	"testing"
)

func TestJustDemo(t *testing.T) {
	ReadExcelDemo()
	// JustDemo()
}

func TestUnique(t *testing.T) {
	type args struct {
		s []*Slot
	}
	tests := []struct {
		name string
		args args
		want []*Slot
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
