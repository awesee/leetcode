package problem283

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []int
	want []int
}

func TestMoveZeroes(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			in:   []int{0, 1, 0, 2, 3, 3, 0, 7, 8},
			want: []int{1, 2, 3, 3, 7, 8, 0, 0, 0},
		},
		{
			in:   []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		got := make([]int, len(tt.in))
		copy(got, tt.in)
		moveZeroes(got)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
