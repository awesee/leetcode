package flipping_an_image

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    [][]int
	expected [][]int
}

func TestFlipAndInvertImage(t *testing.T) {
	tests := [...]caseType{
		{
			input: [][]int{
				{1, 1, 0},
				{1, 0, 1},
				{0, 0, 0},
			},
			expected: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
		},
		{
			input: [][]int{
				{1, 1, 0, 0},
				{1, 0, 0, 1},
				{0, 1, 1, 1},
				{1, 0, 1, 0},
			},
			expected: [][]int{
				{1, 1, 0, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 1},
				{1, 0, 1, 0},
			},
		},
	}
	for _, tc := range tests {
		output := flipAndInvertImage(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
