package image_smoother

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    [][]int
	expected [][]int
}

func TestImageSmoother(t *testing.T) {
	tests := [...]caseType{
		{
			input: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{3, 3, 4},
				{4, 5, 5},
				{6, 6, 7},
			},
		},
	}
	for _, tc := range tests {
		output := imageSmoother(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
