package delete_node_in_a_linked_list

import (
	"reflect"
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	input    []int
	node     int
	expected []int
}

func TestDeleteNode(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{4, 5, 1, 9},
			node:     5,
			expected: []int{4, 1, 9},
		},
		{
			input:    []int{4, 5, 1, 9},
			node:     1,
			expected: []int{4, 5, 9},
		},
		{
			input:    []int{4, 5, 1, 9},
			node:     4,
			expected: []int{5, 1, 9},
		},
		{
			input:    []int{4, 5, 1, 9},
			node:     9,
			expected: []int{4, 5, 1, 9},
		},
		{
			input:    []int{4, 5, 1, 9},
			node:     2,
			expected: []int{4, 5, 1, 9},
		},
	}
	for _, tc := range tests {
		head := SliceInt2ListNode(tc.input)
		node := head
		for node != nil && node.Val != tc.node {
			node = node.Next
		}
		deleteNode(node)
		output := ListNode2SliceInt(head)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
