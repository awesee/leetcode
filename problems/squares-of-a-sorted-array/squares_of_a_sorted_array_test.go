package problem977

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []int
	want []int
}

func TestSortedSquares(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			in:   []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
	}
	for _, tt := range tests {
		got := sortedSquares(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
