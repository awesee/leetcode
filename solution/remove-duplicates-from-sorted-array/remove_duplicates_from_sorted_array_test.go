package remove_duplicates_from_sorted_array

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []int
	expected []int
}

func TestRemoveDuplicates(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 1, 2},
			expected: []int{1, 2},
		},
		{
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			input:    []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tc := range tests {
		nums := make([]int, len(tc.input))
		copy(nums, tc.input)
		l := removeDuplicates(nums)
		output := nums[:l]
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
