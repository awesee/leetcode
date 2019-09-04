package letter_combinations_of_a_phone_number

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    string
	expected []string
}

func TestLetterCombinations(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "",
			expected: nil,
		},
		{
			input:    "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}
	for _, tc := range tests {
		output := letterCombinations(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
