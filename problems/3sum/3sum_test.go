package problem15

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []int
	want [][]int
}

func TestThreeSum(t *testing.T) {
	tests := [...]testType{
		{
			in: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			in: []int{0, 0, 0, 0},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			in: []int{-2, 0, 0, 2, 2, 2},
			want: [][]int{
				{-2, 0, 2},
			},
		},
		{
			in: []int{-2, 0, 0, 2, 2, 2, 2},
			want: [][]int{
				{-2, 0, 2},
			},
		},
	}
	for _, tt := range tests {
		got := threeSum(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
