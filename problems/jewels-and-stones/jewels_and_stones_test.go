package jewels_and_stones

import "testing"

type caseType struct {
	j        string
	s        string
	expected int
}

func TestNumJewelsInStones(t *testing.T) {
	tests := [...]caseType{
		{
			j:        "aA",
			s:        "aAAbbbb",
			expected: 3,
		},
		{
			j:        "z",
			s:        "ZZ",
			expected: 0,
		},
		{
			j:        "Tb",
			s:        "abcdefhelloaabads",
			expected: 2,
		},
	}

	for _, tc := range tests {
		output := numJewelsInStones(tc.j, tc.s)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.j, tc.s, output, tc.expected)
		}
	}
}
