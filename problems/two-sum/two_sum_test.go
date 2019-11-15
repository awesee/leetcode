package problem1

import (
	"reflect"
	"testing"
)

type testType struct {
	nums   []int
	target int
	want   []int
}

func TestTwoSum(t *testing.T) {
	tests := [...]testType{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{2, 7, 11, 15},
			target: 10,
			want:   nil,
		},
		{
			nums:   []int{1, 2, 3, 4},
			target: 5,
			want:   []int{1, 2},
		},
	}
	for _, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.nums, got, tt.want)
		}
	}
}
