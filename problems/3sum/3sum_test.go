package problem_15

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []int
	expected [][]int
}

func TestThreeSum(t *testing.T) {
	tests := [...]caseType{
		{
			input: []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			input: []int{0, 0, 0, 0},
			expected: [][]int{
				{0, 0, 0},
			},
		},
		{
			input: []int{-2, 0, 0, 2, 2, 2},
			expected: [][]int{
				{-2, 0, 2},
			},
		},
		{
			input: []int{-2, 0, 0, 2, 2, 2, 2},
			expected: [][]int{
				{-2, 0, 2},
			},
		},
	}
	for _, tc := range tests {
		output := threeSum(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
