package problem_67

import "testing"

type caseType struct {
	a        string
	b        string
	expected string
}

func TestAddBinary(t *testing.T) {
	tests := [...]caseType{
		{
			a:        "11",
			b:        "1",
			expected: "100",
		},
		{
			a:        "1010",
			b:        "1011",
			expected: "10101",
		},
		{
			a:        "111",
			b:        "111",
			expected: "1110",
		},
	}
	for _, tc := range tests {
		output := addBinary(tc.a, tc.b)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.a, tc.b, output, tc.expected)
		}
	}
}
