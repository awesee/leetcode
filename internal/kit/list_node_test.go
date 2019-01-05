package kit

import (
	"reflect"
	"testing"
)

func TestListNode(t *testing.T) {
	tests := [...][]int{
		{0, 1, 2},
		{1, 2, 3},
	}
	for _, input := range tests {
		l := SliceInt2ListNode(input)
		output := ListNode2SliceInt(l)
		if !reflect.DeepEqual(output, input) {
			t.Fatalf("input: %v, output: %v, expected: %v", input, output, input)
		}
	}
}
