package find_common_characters

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []string
	expected []string
}

func TestCommonChars(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []string{"bella", "label", "roller"},
			expected: []string{"e", "l", "l"},
		},
		{
			input:    []string{"cool", "lock", "cook"},
			expected: []string{"c", "o"},
		},
	}
	for _, tc := range tests {
		output := commonChars(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
