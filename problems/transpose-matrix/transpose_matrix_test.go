package transpose_matrix

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    [][]int
	expected [][]int
}

func TestTranspose(t *testing.T) {
	tests := [...]caseType{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expected: [][]int{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
	}
	for _, tc := range tests {
		output := transpose(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
