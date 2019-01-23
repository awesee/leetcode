package positions_of_large_groups

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    string
	expected [][]int
}

func TestLargeGroupPositions(t *testing.T) {
	tests := [...]caseType{
		{
			input: "abbxxxxzzy",
			expected: [][]int{
				{3, 6},
			},
		},
		{
			input:    "abc",
			expected: [][]int{},
		},
		{
			input: "abcdddeeeeaabbbcd",
			expected: [][]int{
				{3, 5},
				{6, 9},
				{12, 14},
			},
		},
		{
			input: "abcddd",
			expected: [][]int{
				{3, 5},
			},
		},
	}
	for _, tc := range tests {
		output := largeGroupPositions(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
