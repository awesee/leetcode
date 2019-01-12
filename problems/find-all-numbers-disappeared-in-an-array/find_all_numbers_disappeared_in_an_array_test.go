package find_all_numbers_disappeared_in_an_array

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []int
	expected []int
}

func TestFindDisappearedNumbers(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{4, 3, 2, 7, 8, 2, 3, 1},
			expected: []int{5, 6},
		},
		{
			input:    []int{5, 4, 2, 3, 1},
			expected: []int{},
		},
		{
			input:    []int{1, 1},
			expected: []int{2},
		},
	}
	for _, tc := range tests {
		output := findDisappearedNumbers(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
