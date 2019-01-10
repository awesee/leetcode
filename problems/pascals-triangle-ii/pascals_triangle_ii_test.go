package pascals_triangle_ii

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    int
	expected []int
}

func TestGetRow(t *testing.T) {
	tests := [...]caseType{
		{
			input:    2,
			expected: []int{1, 2, 1},
		},
		{
			input:    3,
			expected: []int{1, 3, 3, 1},
		},
		{
			input:    5,
			expected: []int{1, 5, 10, 10, 5, 1},
		},
	}
	for _, tc := range tests {
		output := getRow(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
