package problem566

import (
	"reflect"
	"testing"
)

type testType struct {
	nums [][]int
	r    int
	c    int
	want [][]int
}

func TestMatrixReshape(t *testing.T) {
	tests := [...]testType{
		{
			nums: [][]int{
				{1, 2},
				{3, 4},
			},
			r: 1,
			c: 4,
			want: [][]int{
				{1, 2, 3, 4},
			},
		},
		{
			nums: [][]int{
				{1, 2},
				{3, 4},
			},
			r: 2,
			c: 4,
			want: [][]int{
				{1, 2},
				{3, 4},
			},
		},
	}
	for _, tt := range tests {
		got := matrixReshape(tt.nums, tt.r, tt.c)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.nums, got, tt.want)

		}
	}
}
