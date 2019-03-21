package complement_of_base_10_integer

import "testing"

type caseType struct {
	input    int
	expected int
}

func TestBitwiseComplement(t *testing.T) {
	tests := [...]caseType{
		{
			input:    5,
			expected: 2,
		},
		{
			input:    7,
			expected: 0,
		},
		{
			input:    10,
			expected: 5,
		},
		{
			input:    0,
			expected: 1,
		},
	}
	for _, tc := range tests {
		output := bitwiseComplement(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
