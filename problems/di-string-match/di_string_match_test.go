package di_string_match

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    string
	expected []int
}

func TestDiStringMatch(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "IDID",
			expected: []int{0, 4, 1, 3, 2},
		},
		{
			input:    "III",
			expected: []int{0, 1, 2, 3},
		},
		{
			input:    "DDI",
			expected: []int{3, 2, 0, 1},
		},
	}
	for _, tc := range tests {
		output := diStringMatch(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
