package add_to_array_form_of_integer

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []int
	k        int
	expected []int
}

func TestAddToArrayForm(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 0, 0},
			k:        34,
			expected: []int{1, 2, 3, 4},
		},
		{
			input:    []int{2, 1, 5},
			k:        806,
			expected: []int{1, 0, 2, 1},
		},
		{
			input:    []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			k:        1,
			expected: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tc := range tests {
		output := addToArrayForm(tc.input, tc.k)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
