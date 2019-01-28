package hamming_distance

import "testing"

type caseType struct {
	x        int
	y        int
	expected int
}

func TestHammingDistance(t *testing.T) {
	tests := [...]caseType{
		{
			x:        1,
			y:        4,
			expected: 2,
		},
	}
	for _, tc := range tests {
		output := hammingDistance(tc.x, tc.y)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.x, tc.y, output, tc.expected)
		}
	}
}
