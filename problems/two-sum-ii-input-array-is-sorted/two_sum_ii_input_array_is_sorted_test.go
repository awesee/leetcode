package two_sum_ii_input_array_is_sorted

import (
	"reflect"
	"testing"
)

type caseType struct {
	nums     []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	tests := [...]caseType{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			nums:     []int{2, 7, 11, 15},
			target:   10,
			expected: nil,
		},
		{
			nums:     []int{1, 3, 5, 7, 9, 12},
			target:   16,
			expected: []int{4, 5},
		},
		{
			nums:     []int{1, 2, 3, 4},
			target:   6,
			expected: []int{2, 4},
		},
	}

	for _, tc := range tests {
		output := twoSum(tc.nums, tc.target)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.nums, output, tc.expected)
		}
	}
}
