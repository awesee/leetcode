package binary_prefix_divisible_by_5

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []int
	expected []bool
}

func TestPrefixesDivBy5(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{0, 1, 1},
			expected: []bool{true, false, false},
		},
		{
			input:    []int{1, 1, 1},
			expected: []bool{false, false, false},
		},
		{
			input:    []int{0, 1, 1, 1, 1, 1},
			expected: []bool{true, false, false, false, true, false},
		},
		{
			input:    []int{1, 1, 1, 0, 1},
			expected: []bool{false, false, false, false, false},
		},
		{
			input:    []int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0},
			expected: []bool{false, false, true, false, false, false, false, false, false, false, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, true, false, false, true, true, true, true, true, true, true, false, false, true, false, false, false, false, true, true},
		},
	}
	for _, tc := range tests {
		output := prefixesDivBy5(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
