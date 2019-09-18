package problem_454

import "testing"

type caseType struct {
	a        []int
	b        []int
	c        []int
	d        []int
	expected int
}

func TestFourSumCount(t *testing.T) {
	tests := [...]caseType{
		{
			a:        []int{1, 2},
			b:        []int{-2, -1},
			c:        []int{-1, 2},
			d:        []int{0, 2},
			expected: 2,
		},
	}
	for _, tc := range tests {
		output := fourSumCount(tc.a, tc.b, tc.c, tc.d)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc, output, tc.expected)
		}
	}
}
