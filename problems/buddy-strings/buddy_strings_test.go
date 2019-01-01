package buddy_strings

import "testing"

type caseType struct {
	a        string
	b        string
	expected bool
}

func TestBuddyStrings(t *testing.T) {
	tests := [...]caseType{
		{
			a:        "ab",
			b:        "ba",
			expected: true,
		},
		{
			a:        "aa",
			b:        "aa",
			expected: true,
		},
		{
			a:        "ab",
			b:        "ab",
			expected: false,
		},
		{
			a:        "aaaaaaabc",
			b:        "aaaaaaacb",
			expected: true,
		},
		{
			a:        "",
			b:        "aa",
			expected: false,
		},
		{
			a:        "hello",
			b:        "h0lle",
			expected: false,
		},
		{
			a:        "hello",
			b:        "hanna",
			expected: false,
		},
	}
	for _, tc := range tests {
		output := buddyStrings(tc.a, tc.b)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.a, tc.b, output, tc.expected)
		}
	}
}
