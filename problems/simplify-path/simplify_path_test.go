package simplify_path

import "testing"

type caseType struct {
	input    string
	expected string
}

func TestSimplifyPath(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "/home/",
			expected: "/home",
		},
		{
			input:    "/../",
			expected: "/",
		},
		{
			input:    "/home//foo/",
			expected: "/home/foo",
		},
		{
			input:    "/a/./b/../../c/",
			expected: "/c",
		},
		{
			input:    "/a/../../b/../c//.//",
			expected: "/c",
		},
		{
			input:    "/a//b////c/d//././/..",
			expected: "/a/b/c",
		},
	}
	for _, tc := range tests {
		output := simplifyPath(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
