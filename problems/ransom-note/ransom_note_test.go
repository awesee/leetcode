package ransom_note

import "testing"

type caseType struct {
	ransom   string
	magazine string
	expected bool
}

func TestCanConstruct(t *testing.T) {
	tests := [...]caseType{
		{
			ransom:   "a",
			magazine: "b",
			expected: false,
		},
		{
			ransom:   "aa",
			magazine: "ab",
			expected: false,
		},
		{
			ransom:   "aa",
			magazine: "aab",
			expected: true,
		},
	}
	for _, tc := range tests {
		output := canConstruct(tc.ransom, tc.magazine)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.ransom, tc.magazine, output, tc.expected)
		}
	}
}
