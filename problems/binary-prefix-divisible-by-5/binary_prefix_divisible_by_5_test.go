package problem1018

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []int
	want []bool
}

func TestPrefixesDivBy5(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{0, 1, 1},
			want: []bool{true, false, false},
		},
		{
			in:   []int{1, 1, 1},
			want: []bool{false, false, false},
		},
		{
			in:   []int{0, 1, 1, 1, 1, 1},
			want: []bool{true, false, false, false, true, false},
		},
		{
			in:   []int{1, 1, 1, 0, 1},
			want: []bool{false, false, false, false, false},
		},
		{
			in:   []int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0},
			want: []bool{false, false, true, false, false, false, false, false, false, false, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, true, false, false, true, true, true, true, true, true, true, false, false, true, false, false, false, false, true, true},
		},
	}
	for _, tt := range tests {
		got := prefixesDivBy5(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
