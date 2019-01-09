package sort_array_by_parity_ii

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []int
	expected []int
}

func TestSortArrayByParityII(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{4, 2, 5, 7},
			expected: []int{4, 5, 2, 7},
		},
		{
			input:    []int{2, 3, 1, 1, 4, 0, 0, 4, 3, 3},
			expected: []int{2, 3, 0, 1, 4, 1, 0, 3, 4, 3},
		},
	}
	for _, tc := range tests {
		output := sortArrayByParityII(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
