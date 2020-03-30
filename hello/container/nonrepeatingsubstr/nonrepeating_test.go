package main

import "testing"

func Test_lenthOfNonRepeatingSubStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name          string
		args          args
		wantMaxLength int
	}{
		// TODO: Add test cases.
		{"normal cases", args{"abcabcbb"}, 3},
		{"edge cases", args{""}, 0},
		{"chinese support cases", args{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMaxLength := lenthOfNonRepeatingSubStr(tt.args.s); gotMaxLength != tt.wantMaxLength {
				t.Errorf("lenthOfNonRepeatingSubStr() = %v, want %v", gotMaxLength, tt.wantMaxLength)
			}
		})
	}
}
