package accounts_merge

import (
	"reflect"
	"sort"
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
				{"John", "johnnybravo@mail.com"},
				{"Mary", "mary@mail.com"},
			},
		},
		{
			input: [][]string{
				{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"},
				{"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"},
				{"Kevin", "Kevin4@m.co", "Kevin2@m.co", "Kevin2@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe2@m.co"},
				{"Gabe", "Gabe3@m.co", "Gabe4@m.co", "Gabe2@m.co"},
			},
			expected: [][]string{
				{"Alex", "Alex0@m.co", "Alex4@m.co", "Alex5@m.co"},
				{"Ethan", "Ethan0@m.co", "Ethan3@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe2@m.co", "Gabe3@m.co", "Gabe4@m.co"},
				{"Kevin", "Kevin2@m.co", "Kevin4@m.co"},
			},
		},
		{
			input: [][]string{
				{"David", "David0@m.co", "David4@m.co", "David3@m.co"},
				{"David", "David5@m.co", "David5@m.co", "David0@m.co"},
				{"David", "David1@m.co", "David4@m.co", "David0@m.co"},
				{"David", "David0@m.co", "David1@m.co", "David3@m.co"},
				{"David", "David4@m.co", "David1@m.co", "David3@m.co"},
			},
			expected: [][]string{
				{"David", "David0@m.co", "David1@m.co", "David3@m.co", "David4@m.co", "David5@m.co"},
			},
		},
		{
			input: [][]string{
				{"David", "David0@m.co", "David1@m.co"},
				{"David", "David3@m.co", "David4@m.co"},
				{"David", "David4@m.co", "David5@m.co"},
				{"David", "David2@m.co", "David3@m.co"},
				{"David", "David1@m.co", "David2@m.co"},
			},
			expected: [][]string{
				{"David", "David0@m.co", "David1@m.co", "David2@m.co", "David3@m.co", "David4@m.co", "David5@m.co"},
			},
		},
	}
	for _, tc := range tests {
		output := accountsMerge(tc.input)
		sort.Slice(output, func(i, j int) bool {
			if output[i][0] == output[i][0] {
				return output[i][1] < output[j][1]
			}
			return output[i][0] < output[i][0]
		})
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
