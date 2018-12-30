package integer_to_roman

import "testing"

type caseType struct {
	input    int
	expected string
}

func TestIntToRoman(t *testing.T) {
	tests := [...]caseType{
		{
			input:    3,
			expected: "III",
		},
		{
			input:    4,
			expected: "IV",
		},
		{
			input:    9,
			expected: "IX",
		},
		{
			input:    20,
			expected: "XX",
		},
		{
			input:    58,
			expected: "LVIII",
		},
		{
			input:    1994,
			expected: "MCMXCIV",
		},
		{
			input:    11111,
			expected: "MMMMMMMMMMMCXI",
		},
	}
	for _, tc := range tests {
		output := intToRoman(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}

func BenchmarkIntToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intToRoman(111111)
	}
}
