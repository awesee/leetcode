package flower_planting_with_no_adjacent

import (
	"reflect"
	"testing"
)

type caseType struct {
	n        int
	paths    [][]int
	expected []int
}

func TestGardenNoAdj(t *testing.T) {
	tests := [...]caseType{
		{
			n: 3,
			paths: [][]int{
				{1, 2},
				{2, 3},
				{3, 1},
			},
			expected: []int{1, 2, 3},
		},
		{
			n: 4,
			paths: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: []int{1, 2, 1, 2},
		},
		{
			n: 4,
			paths: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 1},
				{1, 3},
				{2, 4},
			},
			expected: []int{1, 2, 3, 4},
		},
	}
	for _, tc := range tests {
		output := gardenNoAdj(tc.n, tc.paths)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.n, output, tc.expected)
		}
	}
}
