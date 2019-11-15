package problem189

import (
	"reflect"
	"testing"
)

type castType struct {
	nums []int
	k    int
	want []int
}

func TestRotate(t *testing.T) {
	tests := [...]castType{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums: []int{-1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, -1, -100},
		},
		{
			nums: []int{-1},
			k:    2,
			want: []int{-1},
		},
	}

	for _, tt := range tests {
		got := make([]int, len(tt.nums))
		copy(got, tt.nums)
		rotate(got, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
