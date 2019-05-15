package min_stack

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []int
	expected []int
}

func TestMinStack(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{-2, 0, -3},
			expected: []int{-3, 0, -2},
		},
	}
	for _, tc := range tests {
		minStack, output := Constructor(), make([]int, 0)
		for _, x := range tc.input {
			minStack.Push(x)
		}
		output = append(output, minStack.GetMin())
		minStack.Pop()
		output = append(output, minStack.Top())
		output = append(output, minStack.GetMin())
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
