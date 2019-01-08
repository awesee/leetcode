package sort_array_by_parity

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []int
	expected []int
}

func TestSortArrayByParity(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{3, 1, 2, 4},
			expected: []int{2, 4, 3, 1},
		},
	}
	for _, tc := range tests {
		output := sortArrayByParity(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
