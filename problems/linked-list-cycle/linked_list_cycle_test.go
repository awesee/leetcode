package linked_list_cycle

import (
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	input    []int
	pos      int
	expected bool
}

func TestHasCycle(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{3, 2, 0, -4},
			pos:      1,
			expected: true,
		},
		{
			input:    []int{1, 2},
			pos:      0,
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			pos:      -1,
			expected: false,
		},
		{
			input:    []int{1},
			pos:      -1,
			expected: false,
		},
	}
	for _, tc := range tests {
		input := SliceInt2ListNode(tc.input)
		p, curr := input, input
		for i := 0; curr != nil && tc.pos >= 0; i, curr = i+1, curr.Next {
			if i == tc.pos {
				p = curr
			}
			if curr.Next == nil {
				curr.Next = p
				break
			}
		}
		output := hasCycle(input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
