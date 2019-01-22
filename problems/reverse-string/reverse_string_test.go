package reverse_string

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []byte
	expected []byte
}

func TestReverseString(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []byte{'h', 'e', 'l', 'l', 'o'},
			expected: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			input:    []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			expected: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for _, tc := range tests {
		output := make([]byte, len(tc.input))
		copy(output, tc.input)
		reverseString(output)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
