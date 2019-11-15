package problem167

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
			want:   []int{1, 2},
		},
		{
			nums:   []int{2, 7, 11, 15},
			target: 10,
			want:   nil,
		},
		{
			nums:   []int{1, 3, 5, 7, 9, 12},
			target: 16,
			want:   []int{4, 5},
		},
		{
			nums:   []int{1, 2, 3, 4},
			target: 6,
			want:   []int{2, 4},
		},
	}

	for _, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.nums, got, tt.want)
		}
	}
}
