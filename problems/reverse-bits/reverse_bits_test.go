package reverse_bits

import "testing"

type caseType struct {
	input    uint32
	expected uint32
}

func TestReverseBits(t *testing.T) {
	tests := [...]caseType{
		{
			input:    43261596,
			expected: 964176192,
		},
		{
			input:    4294967293,
			expected: 3221225471,
		},
	}
	for _, tc := range tests {
		output := reverseBits(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
