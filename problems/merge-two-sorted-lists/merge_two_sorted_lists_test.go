package merge_two_sorted_lists

import (
	"reflect"
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	l1       []int
	l2       []int
	expected []int
}

func TestMergeTwoLists(t *testing.T) {
	tests := [...]caseType{
		{
			l1:       []int{1, 2, 4},
			l2:       []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			l1:       nil,
			l2:       nil,
			expected: nil,
		},
		{
			l1:       nil,
			l2:       []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			l1:       []int{1, 2, 3},
			l2:       nil,
			expected: []int{1, 2, 3},
		},
		{
			l1:       []int{1, 2, 3, 8, 16},
			l2:       []int{1, 2, 3, 10, 12, 24},
			expected: []int{1, 1, 2, 2, 3, 3, 8, 10, 12, 16, 24},
		},
	}
	for _, tc := range tests {
		l1 := SliceInt2ListNode(tc.l1)
		l2 := SliceInt2ListNode(tc.l2)
		output := ListNode2SliceInt(mergeTwoLists(l1, l2))
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.l1, tc.l2, output, tc.expected)
		}
	}
}
