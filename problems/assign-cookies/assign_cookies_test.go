package problem_455

import "testing"

type caseType struct {
	g        []int
	s        []int
	expected int
}

func TestFindContentChildren(t *testing.T) {
	tests := [...]caseType{
		{
			g:        []int{1, 2, 3},
			s:        []int{1, 1},
			expected: 1,
		},
		{
			g:        []int{1, 2},
			s:        []int{1, 2, 3},
			expected: 2,
		},
	}
	for _, tc := range tests {
		output := findContentChildren(tc.g, tc.s)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.g, output, tc.expected)
		}
	}
}
