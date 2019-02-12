package sum_of_even_numbers_after_queries

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []int
	queries  [][]int
	expected []int
}

func TestSumEvenAfterQueries(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 3, 4},
			queries:  [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}},
			expected: []int{8, 6, 2, 4},
		},
	}
	for _, tc := range tests {
		output := sumEvenAfterQueries(tc.input, tc.queries)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
