package problem_18

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []int
	target   int
	expected [][]int
}

func TestFourSum(t *testing.T) {
	tests := [...]caseType{
		{
			input:  []int{1, 0, -1, 0, -2, 2},
			target: 0,
			expected: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			input:  []int{-1, 0, 1, 2, -1, -4},
			target: 0,
			expected: [][]int{
				{-1, -1, 0, 2},
			},
		},
		{
			input:  []int{0, 0, 0, 0},
			target: 0,
			expected: [][]int{
				{0, 0, 0, 0},
			},
		},
		{
			input:  []int{-2, 0, 0, 2, 2, 2},
			target: 0,
			expected: [][]int{
				{-2, 0, 0, 2},
			},
		},
		{
			input:  []int{-2, 0, 0, 2, 2, 2, 2},
			target: 0,
			expected: [][]int{
				{-2, 0, 0, 2},
			},
		},
		{
			input:  []int{-3, -2, -1, 0, 0, 1, 2, 3},
			target: 0,
			expected: [][]int{
				{-3, -2, 2, 3},
				{-3, -1, 1, 3},
				{-3, 0, 0, 3},
				{-3, 0, 1, 2},
				{-2, -1, 0, 3},
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			input:  []int{1, -2, -5, -4, -3, 3, 3, 5},
			target: -11,
			expected: [][]int{
				{-5, -4, -3, 1},
			},
		},
	}
	for _, tc := range tests {
		output := fourSum(tc.input, tc.target)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
