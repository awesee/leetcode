package self_dividing_numbers

import (
	"reflect"
	"testing"
)

type caseType struct {
	left     int
	right    int
	expected []int
}

func TestSelfDividingNumbers(t *testing.T) {
	tests := [...]caseType{
		{
			left:     1,
			right:    22,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},
	}
	for _, tc := range tests {
		output := selfDividingNumbers(tc.left, tc.right)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.left, tc.right, output, tc.expected)
		}
	}
}
