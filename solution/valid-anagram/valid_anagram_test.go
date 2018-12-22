package valid_anagram

import "testing"

type caseType struct {
	s        string
	t        string
	expected bool
}

func TestIsAnagram(t *testing.T) {
	tests := [...]caseType{
		{
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			s:        "hannah",
			t:        "hanna",
			expected: false,
		},
		{
			s:        "this is string",
			t:        "this is string long",
			expected: false,
		},
		{
			s:        "this is string",
			t:        "this as string",
			expected: false,
		},
		{
			s:        "你好，世界",
			t:        "世界，你好",
			expected: true,
		},
	}

	for _, tc := range tests {
		output := isAnagram(tc.s, tc.t)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.s, tc.t, output, tc.expected)
		}
	}
}
