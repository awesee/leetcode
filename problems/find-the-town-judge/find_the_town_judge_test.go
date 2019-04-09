package find_the_town_judge

import "testing"

type caseType struct {
	N        int
	trust    [][]int
	expected int
}

func TestFindJudge(t *testing.T) {
	tests := [...]caseType{
		{
			N: 2,
			trust: [][]int{
				{1, 2},
			},
			expected: 2,
		},
		{
			N: 3,
			trust: [][]int{
				{1, 3},
				{2, 3},
			},
			expected: 3,
		},
		{
			N: 3,
			trust: [][]int{
				{1, 3},
				{2, 3},
				{3, 1},
			},
			expected: -1,
		},
		{
			N: 3,
			trust: [][]int{
				{1, 2},
				{2, 3},
			},
			expected: -1,
		},
		{
			N: 4,
			trust: [][]int{
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{4, 3},
			},
			expected: 3,
		},
		{
			N:        1,
			trust:    [][]int{},
			expected: 1,
		},
	}
	for _, tc := range tests {
		output := findJudge(tc.N, tc.trust)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.N, tc.trust, output, tc.expected)
		}
	}
}
