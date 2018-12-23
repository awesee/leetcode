package remove_element

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []int
	val      int
	expected []int
}

func TestRemoveElement(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{3, 2, 2, 3},
			val:      3,
			expected: []int{2, 2},
		},
		{
			input:    []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: []int{0, 1, 3, 0, 4},
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			val:      6,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			input:    []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			val:      4,
			expected: []int{1, 2, 2, 3, 3, 3},
		},
		{
			input:    []int{},
			val:      1,
			expected: []int{},
		},
	}

	for _, tc := range tests {
		nums := make([]int, len(tc.input))
		copy(nums, tc.input)
		l := removeElement(nums, tc.val)
		output := nums[:l]
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
