package pascals_triangle

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    int
	expected [][]int
}

func TestGenerate(t *testing.T) {
	tests := [...]caseType{
		{
			input: 5,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}
	for _, tc := range tests {
		output := generate(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
