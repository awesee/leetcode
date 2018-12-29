package number_of_1_bits

import "testing"

type caseType struct {
	input    uint32
	expected int
}

func TestHammingWeight(t *testing.T) {
	tests := [...]caseType{
		{
			input:    3,
			expected: 2,
		},
		{
			input:    7,
			expected: 3,
		},
		{
			input:    8,
			expected: 1,
		},
		{
			input:    0x5555,
			expected: 8,
		},
	}
	for _, tc := range tests {
		output := hammingWeight(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}

func BenchmarkHammingWeight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hammingWeight(0x55555555)
	}
}
