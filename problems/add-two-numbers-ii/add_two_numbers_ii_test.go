package problem_445

import (
	"reflect"
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	l1       []int
	l2       []int
	expected []int
}

func TestAddTwoNumbers(t *testing.T) {
	tests := [...]caseType{
		{
			l1:       []int{7, 2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 8, 0, 7},
		},
		{
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{8, 0, 7},
		},
		{
			l1:       []int{5},
			l2:       []int{5},
			expected: []int{1, 0},
		},
		{
			l1:       []int{5},
			l2:       []int{5, 9, 9},
			expected: []int{6, 0, 4},
		},
	}
	for _, tc := range tests {
		l1 := kit.SliceInt2ListNode(tc.l1)
		l2 := kit.SliceInt2ListNode(tc.l2)
		output := kit.ListNode2SliceInt(addTwoNumbers(l1, l2))
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.l1, tc.l2, output, tc.expected)
		}
	}
}
