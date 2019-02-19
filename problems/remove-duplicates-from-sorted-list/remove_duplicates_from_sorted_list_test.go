package remove_duplicates_from_sorted_list

import (
	"reflect"
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	input    []int
	expected []int
}

func TestDeleteDuplicates(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 1, 2},
			expected: []int{1, 2},
		},
		{
			input:    []int{1, 1, 2, 3, 3},
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{0, 0, 1, 1, 2, 2},
			expected: []int{0, 1, 2},
		},
	}
	for _, tc := range tests {
		output := deleteDuplicates(SliceInt2ListNode(tc.input))
		if !reflect.DeepEqual(ListNode2SliceInt(output), tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
