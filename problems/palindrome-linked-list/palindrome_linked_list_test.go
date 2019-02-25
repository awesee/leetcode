package palindrome_linked_list

import (
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	input    []int
	expected bool
}

func TestIsPalindrome(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2},
			expected: false,
		},
		{
			input:    []int{1, 2, 1},
			expected: true,
		},
		{
			input:    []int{1, 2, 2, 1},
			expected: true,
		},
	}
	for _, tc := range tests {
		output := isPalindrome(SliceInt2ListNode(tc.input))
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
