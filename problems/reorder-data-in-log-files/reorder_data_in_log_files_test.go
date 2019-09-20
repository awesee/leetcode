package problem_937

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []string
	expected []string
}

func TestReorderLogFiles(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"},
			expected: []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
		},
		{
			input:    []string{"8 fj dzz k", "5r 446 6 3", "63 gu psub", "5 ba iedrz", "6s 87979 5", "3r 587 01", "jc 3480612", "bb wsrd kp", "b aq cojj", "r5 6316 71"},
			expected: []string{"b aq cojj", "5 ba iedrz", "8 fj dzz k", "63 gu psub", "bb wsrd kp", "5r 446 6 3", "6s 87979 5", "3r 587 01", "jc 3480612", "r5 6316 71"},
		},
		{
			input:    []string{"qi ir oo i", "cp vnzw i", "0 fkbikbts", "4 j trouka", "gn j q al", "5r w wgqc", "m8 x haje", "fg 28694 6", "i gf mwdoa", "ao 0850716"},
			expected: []string{"0 fkbikbts", "i gf mwdoa", "qi ir oo i", "gn j q al", "4 j trouka", "cp vnzw i", "5r w wgqc", "m8 x haje", "fg 28694 6", "ao 0850716"},
		},
	}
	for _, tc := range tests {
		output := make([]string, len(tc.input))
		copy(output, tc.input)
		output = reorderLogFiles(output)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
