package reshape_the_matrix

import (
	"reflect"
	"testing"
)

type caseType struct {
	nums     [][]int
	r        int
	c        int
	expected [][]int
}

func TestMatrixReshape(t *testing.T) {
	tests := [...]caseType{
		{
			nums: [][]int{
				{1, 2},
				{3, 4},
			},
			r: 1,
			c: 4,
			expected: [][]int{
				{1, 2, 3, 4},
			},
		},
		{
			nums: [][]int{
				{1, 2},
				{3, 4},
			},
			r: 2,
			c: 4,
			expected: [][]int{
				{1, 2},
				{3, 4},
			},
		},
	}
	for _, tc := range tests {
		output := matrixReshape(tc.nums, tc.r, tc.c)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.nums, output, tc.expected)

		}
	}
}
