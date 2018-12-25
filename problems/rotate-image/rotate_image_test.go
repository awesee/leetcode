package rotate_image

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    [][]int
	expected [][]int
}

func TestRotate(t *testing.T) {
	tests := [...]caseType{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			input: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}

	for _, tc := range tests {
		l := len(tc.input)
		output := make([][]int, l)
		for i := 0; i < l; i++ {
			output[i] = make([]int, l)
			copy(output[i], tc.input[i])
		}
		rotate(output)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
