package string_compression

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []byte
	expected []byte
}

func TestCompress(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected: []byte{'a', '2', 'b', '2', 'c', '3'},
		},
		{
			input:    []byte{'a'},
			expected: []byte{'a'},
		},
		{
			input:    []byte{'a', 'a'},
			expected: []byte{'a', '2'},
		},
		{
			input:    []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expected: []byte{'a', 'b', '1', '2'},
		},
	}
	for _, tc := range tests {
		l := compress(tc.input)
		if !reflect.DeepEqual(tc.input[:l], tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, l, tc.expected)
		}
	}
}
