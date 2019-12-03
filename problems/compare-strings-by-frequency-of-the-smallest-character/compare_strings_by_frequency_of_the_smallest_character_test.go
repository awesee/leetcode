package problem1170

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []string
	w    []string
	want []int
}

func TestNumSmallerByFrequency(t *testing.T) {
	tests := [...]testType{
		{
			in:   []string{"cbd"},
			w:    []string{"zaaaz"},
			want: []int{1},
		},
		{
			in:   []string{"bbb", "cc"},
			w:    []string{"a", "aa", "aaa", "aaaa"},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		got := numSmallerByFrequency(tt.in, tt.w)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
