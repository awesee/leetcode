package accounts_merge

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    [][]string
	expected [][]string
}

func TestAccountsMerge(t *testing.T) {
	tests := [...]caseType{
		{
			input: [][]string{
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"Mary", "mary@mail.com"},
			},
			expected: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"John", "johnnybravo@mail.com"}, {"Mary", "mary@mail.com"},
			},
		},
	}
	for _, tc := range tests {
		output := accountsMerge(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
