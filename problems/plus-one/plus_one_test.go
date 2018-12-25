package plus_one

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []int
	expected []int
}

func TestPlusOne(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			input:    []int{4, 3, 2, 1},
			expected: []int{4, 3, 2, 2},
		},
		{
			input:    []int{1, 2, 9, 9},
			expected: []int{1, 3, 0, 0},
		},
		{
			input:    []int{9, 9},
			expected: []int{1, 0, 0},
		},
	}

	for _, tc := range tests {
		output := plusOne(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
