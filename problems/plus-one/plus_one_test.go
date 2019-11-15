package problem66

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []int
	want []int
}

func TestPlusOne(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 3},
			want: []int{1, 2, 4},
		},
		{
			in:   []int{4, 3, 2, 1},
			want: []int{4, 3, 2, 2},
		},
		{
			in:   []int{1, 2, 9, 9},
			want: []int{1, 3, 0, 0},
		},
		{
			in:   []int{9, 9},
			want: []int{1, 0, 0},
		},
	}

	for _, tt := range tests {
		got := plusOne(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
