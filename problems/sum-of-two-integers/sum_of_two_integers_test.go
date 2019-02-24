package sum_of_two_integers

import "testing"

type caseType struct {
	a        int
	b        int
	expected int
}

func TestGetSum(t *testing.T) {
	tests := [...]caseType{
		{
			a:        1,
			b:        2,
			expected: 3,
		},
		{
			a:        -2,
			b:        3,
			expected: 1,
		},
	}
	for _, tc := range tests {
		output := getSum(tc.a, tc.b)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.a, tc.b, output, tc.expected)
		}
	}
}
