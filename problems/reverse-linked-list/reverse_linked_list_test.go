package reverse_linked_list

import (
	"reflect"
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	input    []int
	expected []int
}

func TestReverseList(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tc := range tests {
		output := ListNode2SliceInt(reverseList(SliceInt2ListNode(tc.input)))
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
