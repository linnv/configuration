package demo

import "testing"

func TestJustDemo(t *testing.T) {
	// JustDemo()
	// PhoneDemo( "18290015121")
	// PersonalIDDemo("")
	// CarNOValidationDemo("")
}

func TestCarNOValidationLocalDemo(t *testing.T) {
	type args struct {
		cn string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{"6228481698729890079"}, true},
		{"normal", args{"6228481698729890"}, true},
		{"normal", args{"622848169872989a"}, false},
		{"normal", args{"622848169872989007b"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CarNOValidationLocalDemo(tt.args.cn); got != tt.want {
				t.Errorf("CarNOValidationLocalDemo(%s) = %v, want %v", tt.args.cn, got, tt.want)
			}
		})
	}
}

func TestPhoneDemo(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{"18290015121"}, true},
		{"normal", args{"11290015121"}, false},
		{"normal", args{"13290015121"}, true},
		{"normal", args{"0777-8581189"}, true},
		{"normal", args{"07778581189"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PhoneDemo(tt.args.number); got != tt.want {
				t.Errorf("PhoneDemo(%s) = %v, want %v", tt.args.number, got, tt.want)
			}
		})
	}
}

func TestRDemo(t *testing.T) {
	type args struct {
		cn string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{"1"}, false},
		{"normal", args{"a"}, false},
		{"normal", args{"2a"}, false},
		{"normal", args{"22"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RDemo(tt.args.cn); got != tt.want {
				t.Errorf("RDemo(%s) = %v, want %v", tt.args.cn, got, tt.want)
			}
		})
	}
}

func TestPersonalIDDemo(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{"35052419880112712X"}, true},
		{"normal", args{"450722199111152834"}, true},
		{"normal", args{"05052419880112712X"}, false},
		{"normal", args{"43048119930926003X"}, true},
		{"normal", args{"450222199207250028"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PersonalIDDemo(tt.args.id); got != tt.want {
				t.Errorf("PersonalIDDemo(%s) = %v, want %v", tt.args.id, got, tt.want)
			}
		})
	}
}
