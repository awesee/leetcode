package k_closest_points_to_origin

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    [][]int
	k        int
	expected [][]int
}

func TestKClosest(t *testing.T) {
	tests := [...]caseType{
		{
			input: [][]int{
				{1, 3},
				{-2, 2},
			},
			k: 1,
			expected: [][]int{
				{-2, 2},
			},
		},
		{
			input: [][]int{
				{3, 3},
				{5, -1},
				{-2, 4},
			},
			k: 2,
			expected: [][]int{
				{3, 3},
				{-2, 4},
			},
		},
	}
	for _, tc := range tests {
		output := kClosest(tc.input, tc.k)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
