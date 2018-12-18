package implement_strstr

import "testing"

type caseType struct {
	haystack string
	needle   string
	expected int
}

func TestStrStr(t *testing.T) {
	tests := [...]caseType{
		{
			haystack: "hello",
			needle:   "ll",
			expected: 2,
		}, {
			haystack: "this is test string",
			needle:   "a",
			expected: -1,
		},
	}

	for _, tc := range tests {
		output := strStr(tc.haystack, tc.needle)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.haystack, tc.needle, output, tc.expected)
		}
	}
}
