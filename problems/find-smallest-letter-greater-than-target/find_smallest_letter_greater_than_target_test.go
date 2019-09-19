package problem_744

import "testing"

type caseType struct {
	input    []byte
	target   byte
	expected byte
}

func TestNextGreatestLetter(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []byte{'c', 'f', 'j'},
			target:   'a',
			expected: 'c',
		},
		{
			input:    []byte{'c', 'f', 'j'},
			target:   'c',
			expected: 'f',
		},
		{
			input:    []byte{'c', 'f', 'j'},
			target:   'd',
			expected: 'f',
		},
		{
			input:    []byte{'c', 'f', 'j'},
			target:   'g',
			expected: 'j',
		},
		{
			input:    []byte{'c', 'f', 'j'},
			target:   'j',
			expected: 'c',
		},
		{
			input:    []byte{'c', 'f', 'j'},
			target:   'k',
			expected: 'c',
		},
	}
	for _, tc := range tests {
		output := nextGreatestLetter(tc.input, tc.target)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
