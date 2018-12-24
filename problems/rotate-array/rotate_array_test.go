package rotate_array

import (
	"reflect"
	"testing"
)

type castType struct {
	nums     []int
	k        int
	expected []int
}

func TestRotate(t *testing.T) {
	tests := [...]castType{
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
		{
			nums:     []int{-1},
			k:        2,
			expected: []int{-1},
		},
	}

	for _, tc := range tests {
		output := make([]int, len(tc.nums))
		copy(output, tc.nums)
		rotate(output, tc.k)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.nums, tc.k, output, tc.expected)
		}
	}
}
