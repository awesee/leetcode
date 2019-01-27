package delete_columns_to_make_sorted

import "testing"

type caseType struct {
	input    []string
	expected int
}

func TestMinDeletionSize(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []string{"cba", "daf", "ghi"},
			expected: 1,
		},
		{
			input:    []string{"a", "b"},
			expected: 0,
		},
		{
			input:    []string{"zyx", "wvu", "tsr"},
			expected: 3,
		},
	}
	for _, tc := range tests {
		output := minDeletionSize(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
